package white_list

import (
	"math/rand/v2"
	"time"
)

var (
	dynastyIds = []uint64{
		891197,
		27643,
		244,
		28481,
		26527,
		115816,
		115610,
		891380,
		733108,
		230,
		31879,
		792752,
		734224,
		91757,
		792702,
		32883,
		36110,
		792870,
		26544,
		122532,
		115652,
		115890,
		762785,
		38194,
		1766,
		734049,
		128176,
	}
)

func GetRandMouseId() uint64 {
	destination := make([]uint64, len(dynastyIds))
	copy(destination, dynastyIds)
	rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))
	rand.Shuffle(len(destination), func(i, j int) {
		destination[i], destination[j] = destination[j], destination[i]
	})

	return destination[0]
}

func GetRandLength(length int) []uint64 {
	destination := make([]uint64, len(dynastyIds))
	copy(destination, dynastyIds)
	rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))
	rand.Shuffle(len(destination), func(i, j int) {
		destination[i], destination[j] = destination[j], destination[i]
	})
	return destination[:length]
}

func All() []uint64 {
	return dynastyIds
}
