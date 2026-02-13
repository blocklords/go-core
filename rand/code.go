package rand_any

import (
	"crypto/aes"
	"encoding/binary"
	"fmt"
	"sync/atomic"
	"time"
)

type (
	RandCode struct {
		charset []rune
		base    uint64
		secret  []byte

		machineID uint64

		state uint64 // 原子状态：时间+序列
	}

	RandFn func(e *RandCode)
)

const (
	workerBits   = 10
	sequenceBits = 12

	maxSequence = (1 << sequenceBits) - 1

	workerShift    = sequenceBits
	timestampShift = sequenceBits + workerBits

	epoch = uint64(1704067200000)
)

//
// ========== Option ==========
//

func RandCodeMachineID(machineID uint64) RandFn {
	return func(e *RandCode) {
		e.machineID = machineID
	}
}

func RandCodeSecret(secret []byte) RandFn {
	return func(e *RandCode) {
		e.secret = secret
	}
}

//
// ========== Constructor ==========
//

func NewRandCode(fns ...RandFn) *RandCode {
	charset := []rune("ABCDEFGHJKMNPQRSTUVWXYZ23456789")

	engine := &RandCode{
		charset:   charset,
		base:      uint64(len(charset)),
		secret:    []byte("liu-willow-16key"), // 必须16字节
		machineID: 1,
	}

	for _, fn := range fns {
		fn(engine)
	}

	if len(engine.secret) != 16 {
		panic("secret must be 16 bytes")
	}

	return engine
}

//
// ========== 无锁 Snowflake ==========
//

func (rc *RandCode) nextID() uint64 {
	for {
		old := atomic.LoadUint64(&rc.state)

		lastTime := old >> sequenceBits
		sequence := old & maxSequence

		now := uint64(time.Now().UnixMilli() - int64(epoch))

		var newTime uint64
		var newSeq uint64

		if now == lastTime {
			newSeq = (sequence + 1) & maxSequence
			if newSeq == 0 {
				continue
			}
			newTime = now
		} else if now > lastTime {
			newTime = now
			newSeq = 0
		} else {
			continue
		}

		newState := (newTime << sequenceBits) | newSeq

		if atomic.CompareAndSwapUint64(&rc.state, old, newState) {
			id := (newTime << timestampShift) |
				(rc.machineID << workerShift) |
				newSeq

			return id
		}
	}
}

//
// ========== AES 加密 ==========
//

func (rc *RandCode) encrypt(id uint64) uint64 {
	block, _ := aes.NewCipher(rc.secret)

	buf := make([]byte, 16)
	binary.LittleEndian.PutUint64(buf, id)

	block.Encrypt(buf, buf)

	return binary.LittleEndian.Uint64(buf)
}

//
// ========== Base32 编码 ==========
//

func (rc *RandCode) encode(num uint64) string {
	buf := make([]rune, 12)

	for i := 11; i >= 0; i-- {
		buf[i] = rc.charset[num%rc.base]
		num /= rc.base
	}

	return fmt.Sprintf("%s-%s-%s",
		string(buf[0:4]),
		string(buf[4:8]),
		string(buf[8:12]),
	)
}

//
// ========== 对外接口 ==========
//

func (rc *RandCode) Make() string {
	id := rc.nextID()
	encrypted := rc.encrypt(id)
	return rc.encode(encrypted)
}
