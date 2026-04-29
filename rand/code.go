package rand_any

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"encoding/binary"
	"math"
	"sync/atomic"
	"time"
)

const (
	workerBits   = 10
	sequenceBits = 12

	maxSequence = (1 << sequenceBits) - 1

	workerShift    = sequenceBits
	timestampShift = sequenceBits + workerBits

	epoch = uint64(1704067200000)
)

type (
	IFormat interface {
		Format(string) string
	}
	FormatString struct {
		group int    // 分割位数
		sep   string // 分割字符
	}
	FormatNone struct{}

	FormatFn func(f *FormatString)
	RandCode struct {
		charset   []byte
		base      uint64
		machineID uint64
		secret    [16]byte
		state     uint64
		block     cipher.Block

		length   int
		formatFn IFormat
	}
	RandFn func(e *RandCode)
)

// ================= Option =================

func FormatStringGroup(group int) FormatFn {
	return func(f *FormatString) {
		f.group = group
	}
}
func FormatStringSep(sep string) FormatFn {
	return func(f *FormatString) {
		f.sep = sep
	}
}

func NewFormatString(fns ...FormatFn) *FormatString {
	format := &FormatString{group: 4, sep: "-"}
	for _, fn := range fns {
		fn(format)
	}
	return format
}

func (f *FormatString) Format(s string) string {
	if f.group <= 0 {
		return s
	}
	n := len(s)
	out := make([]byte, 0, n+n/f.group)
	for i := 0; i < n; i++ {
		if i > 0 && i%f.group == 0 {
			out = append(out, f.sep...)
		}
		out = append(out, s[i])
	}

	return string(out)
}

func NewFormatNone() *FormatNone {
	return &FormatNone{}
}

func (f *FormatNone) Format(s string) string {
	return s
}

func RandCodeMachineID(machineID uint64) RandFn {
	return func(e *RandCode) {
		e.machineID = machineID
	}
}

func RandCodeSecret(secret [16]byte) RandFn {
	return func(e *RandCode) {
		e.secret = secret
	}
}

func RandCodeLength(length int) RandFn {
	return func(e *RandCode) {
		e.length = length
	}
}

func RandCodeFormat(fn IFormat) RandFn {
	return func(e *RandCode) {
		e.formatFn = fn
	}
}

// ================= Constructor =================

func NewRandCode(fns ...RandFn) *RandCode {
	charset := []byte("ABCDEFGHJKMNPQRSTUVWXYZ23456789")

	rc := &RandCode{
		charset:   charset,
		base:      uint64(len(charset)),
		machineID: 1,
		secret:    [16]byte{'l', 'i', 'u', '-', 'w', 'i', 'l', 'l', 'o', 'w', '-', '1', '6', 'k', 'e', 'y'},
		length:    12,
		formatFn:  NewFormatNone(),
	}

	for _, fn := range fns {
		fn(rc)
	}

	block, err := aes.NewCipher(rc.secret[:])
	if err != nil {
		panic(err)
	}
	rc.block = block

	if rc.length < 6 {
		panic("length must >= 6")
	}

	return rc
}

// ================= Snowflake =================

func (rc *RandCode) nextID() uint64 {
	for {
		old := atomic.LoadUint64(&rc.state)

		lastTime := old >> sequenceBits
		sequence := old & maxSequence

		now := uint64(time.Now().UnixMilli()) - epoch

		if now < lastTime {
			now = lastTime
		}

		var newSeq uint64
		if now == lastTime {
			newSeq = (sequence + 1) & maxSequence
			if newSeq == 0 {
				continue
			}
		} else {
			newSeq = 0
		}

		newState := (now << sequenceBits) | newSeq

		if atomic.CompareAndSwapUint64(&rc.state, old, newState) {
			return (now << timestampShift) |
				(rc.machineID << workerShift) |
				newSeq
		}
	}
}

// ================= crypto random =================

func randUint32() uint32 {
	var b [4]byte
	_, _ = crand.Read(b[:])
	return binary.LittleEndian.Uint32(b[:])
}

// ================= AES 混淆 =================

func (rc *RandCode) encrypt(id uint64) uint64 {
	salt := uint64(randUint32())

	id ^= salt << 32

	var buf [16]byte
	binary.LittleEndian.PutUint64(buf[:], id)

	rc.block.Encrypt(buf[:], buf[:])

	return binary.LittleEndian.Uint64(buf[:])
}

// ================= 编码 =================

func (rc *RandCode) encodeFixed(num uint64, length int) string {
	buf := make([]byte, length)

	for i := length - 1; i >= 0; i-- {
		buf[i] = rc.charset[num%rc.base]
		num /= rc.base
	}

	return string(buf)
}

// 计算需要多少位可以容纳 bits
func calcCoreLen(bits int, base uint64) int {
	return int(math.Ceil(float64(bits) / math.Log2(float64(base))))
}

// ================= 外部接口 =================

func (rc *RandCode) Make() string {

	id := rc.nextID()

	// ===== 根据长度决定使用多少bit =====
	bits := 60
	if rc.length > 12 {
		bits = 64 // ✅ 压缩到 60bit，安全适配12位
	}

	if bits < 64 {
		id &= (1 << bits) - 1
	}

	encrypted := rc.encrypt(id)

	// 核心编码长度
	coreLen := calcCoreLen(bits, rc.base)

	encode := rc.encodeFixed(encrypted, coreLen)

	var code string

	if rc.length <= coreLen {
		code = encode[:rc.length]
	} else {
		prefixLen := rc.length - coreLen
		buf := make([]byte, rc.length)

		// 随机前缀
		for i := 0; i < prefixLen; i++ {
			buf[i] = rc.charset[randUint32()%uint32(len(rc.charset))]
		}

		copy(buf[prefixLen:], encode)

		code = string(buf)
	}

	if rc.formatFn != nil {
		return rc.formatFn.Format(code)
	}

	return code
}
