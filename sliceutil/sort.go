package sliceutil

import "math/rand/v2"

// Shuffle 随机打乱切片元素 (Fisher-Yates 算法)
func Shuffle[T comparable](slice []T) {
	// 不需要手动初始化随机种子，rand/v2 自动处理
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
