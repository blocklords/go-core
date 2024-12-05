package fn

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"math"
	"math/rand"
	"time"
)

func Now() int64 {
	return time.Now().UTC().Unix()
}

func NowMs() int64 {
	return time.Now().UTC().UnixMilli()
}

func NowNs() int64 {
	return time.Now().UTC().UnixNano()
}

func Date() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}

func HashFromUUID() common.Hash {
	return common.HexToHash(hex.EncodeToString(Sha1([]byte(uuid.New().String()))))
}

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func RandomCode(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RoundDownToTwoDecimalPlaces(f float64, digits int) float64 {
	if digits < 1 {
		return 0
	}
	dunit := math.Pow(10, float64(digits))
	return math.Floor(f*dunit) / dunit
}
