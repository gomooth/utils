package sliceutil_test

import (
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/gomooth/utils/sliceutil"
	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5}
		copied := make([]int, len(original))
		copy(copied, original)

		sliceutil.Shuffle(copied)

		assert.ElementsMatch(t, original, copied, "元素应该相同")
		assert.NotEqual(t, original, copied, "顺序应该不同")
	})

	t.Run("string slice", func(t *testing.T) {
		original := []string{"a", "b", "c", "d", "e"}
		copied := make([]string, len(original))
		copy(copied, original)

		sliceutil.Shuffle(copied)

		assert.ElementsMatch(t, original, copied)
		assert.NotEqual(t, original, copied)
	})

	t.Run("empty slice", func(t *testing.T) {
		empty := []int{}
		sliceutil.Shuffle(empty)
		assert.Empty(t, empty)
	})

	t.Run("single element", func(t *testing.T) {
		single := []string{"alone"}
		sliceutil.Shuffle(single)
		assert.Equal(t, []string{"alone"}, single)
	})

	t.Run("struct slice", func(t *testing.T) {
		type point struct{ x, y int }
		original := []point{{1, 2}, {3, 4}, {5, 6}}
		copied := make([]point, len(original))
		copy(copied, original)

		sliceutil.Shuffle(copied)

		assert.ElementsMatch(t, original, copied)
		assert.NotEqual(t, original, copied)
	})

	t.Run("duplicate elements", func(t *testing.T) {
		original := []int{1, 1, 2, 2, 3, 3}
		copied := make([]int, len(original))
		copy(copied, original)

		sliceutil.Shuffle(copied)

		assert.ElementsMatch(t, original, copied)
	})
}

func TestShuffle_DeterministicWithFixedSeed(t *testing.T) {
	// 创建使用固定种子的本地随机源
	r := rand.New(rand.NewPCG(42, 0)) // 固定种子

	original := []int{1, 2, 3, 4, 5}
	copied1 := make([]int, len(original))
	copied2 := make([]int, len(original))
	copy(copied1, original)
	copy(copied2, original)

	// 第一次打乱
	for i := len(copied1) - 1; i > 0; i-- {
		j := r.IntN(i + 1)
		copied1[i], copied1[j] = copied1[j], copied1[i]
	}

	// 重置随机源
	r = rand.New(rand.NewPCG(42, 0)) // 相同的种子

	// 第二次打乱
	for i := len(copied2) - 1; i > 0; i-- {
		j := r.IntN(i + 1)
		copied2[i], copied2[j] = copied2[j], copied2[i]
	}

	assert.Equal(t, copied1, copied2, "相同种子应产生相同结果")
}

func TestShuffle_AllPermutationsPossible(t *testing.T) {
	// 测试所有元素都有机会出现在每个位置
	original := []int{1, 2, 3}
	positionCount := make([]map[int]int, len(original))
	for i := range positionCount {
		positionCount[i] = make(map[int]int)
	}

	const trials = 10000
	for i := 0; i < trials; i++ {
		copied := make([]int, len(original))
		copy(copied, original)
		sliceutil.Shuffle(copied)

		for pos, val := range copied {
			positionCount[pos][val]++
		}
	}

	// 检查每个值在每个位置出现的频率
	for pos, counts := range positionCount {
		for _, val := range original {
			count := counts[val]
			expected := trials / len(original)
			assert.InDelta(t, expected, count, float64(expected)*0.1,
				"值 %d 在位置 %d 出现次数异常", val, pos)
		}
	}
}

func BenchmarkShuffle(b *testing.B) {
	sizes := []int{10, 100, 1000}
	for _, size := range sizes {
		b.Run("size-"+strconv.Itoa(size), func(b *testing.B) {
			data := make([]int, size)
			for i := range data {
				data[i] = i
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sliceutil.Shuffle(data)
			}
		})
	}
}
