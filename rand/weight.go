package rand_any

// // 随机权重

type (
	IWeight interface {
		Weight() int64
	}
	Weights[T IWeight] struct {
		items []T
	}
)

// Rand 加权随机
func (w Weights[T]) Rand() T {
	totalWeight := int64(0)
	for _, item := range w.items {
		totalWeight += item.Weight()
	}

	intX := IntX[int64]{}
	r := intX.Rand(0, totalWeight)

	// 根据随机数选择项目
	for _, item := range w.items {
		r -= item.Weight()
		if r < 0 {
			return item
		}
	}

	// 默认情况下返回最后一个项目（理论上不应该发生）
	return w.items[len(w.items)-1]
}
