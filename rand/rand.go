package rand_any

import (
	"math/rand/v2"
	"time"
)

type (
	IInt interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}
	IUint interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
	}
	IFloat interface {
		~float32 | ~float64
	}

	IntX[I IInt]     struct{}
	UintX[I IUint]   struct{}
	FloatX[I IFloat] struct{}

	Model interface {
		IInt | IUint | IFloat
	}

	IRand[M Model] interface {
		Rand(min, max M) M
	}
	Rand[M Model, I IRand[M]] struct {
		engine I
	}
	Fn[M Model, I IRand[M]] func(e *Rand[M, I])
)

func (r *IntX[I]) Rand(min, max I) I {
	if min == max {
		return max
	}
	if min > max {
		min, max = max, min
	}
	return I(rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0)).Int64N(int64(max-min)) + int64(min))
}

func (r *UintX[I]) Rand(min, max I) I {
	if min == max {
		return max
	}
	if min > max {
		min, max = max, min
	}
	return I(rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0)).Uint64N(uint64(max-min)) + uint64(min))
}

func (r *FloatX[I]) Rand(min, max I) I {
	if min == max {
		return max
	}
	if min > max {
		min, max = max, min
	}
	return I(rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0)).Float64()*(float64(max)-float64(min)) + float64(min))
}

func RandWithEngine[M Model, I IRand[M]](en I) Fn[M, I] {
	return func(e *Rand[M, I]) {
		e.engine = en
	}
}

func NewRand[M Model, I IRand[M]](fns ...Fn[M, I]) *Rand[M, I] {
	e := &Rand[M, I]{}
	for _, fn := range fns {
		fn(e)
	}
	return e
}

func (r *Rand[M, I]) Rand(min, max M) M {
	return r.engine.Rand(min, max)
}
