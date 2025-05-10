package valutil_test

import (
	"sync"
	"testing"

	"github.com/gomooth/utils/valutil"

	"github.com/stretchr/testify/assert"
)

func TestWrap_Ptr_Val(t *testing.T) {
	// 基本类型测试
	t.Run("string类型", func(t *testing.T) {
		s := "hello"
		sp := valutil.Ptr(s)
		assert.Equal(t, s, *sp)
		assert.Equal(t, s, valutil.Val(sp))
	})

	t.Run("int类型", func(t *testing.T) {
		n := 42
		np := valutil.Ptr(n)
		assert.Equal(t, n, *np)
		assert.Equal(t, n, valutil.Val(np))
	})

	t.Run("float64类型", func(t *testing.T) {
		f := 3.14
		fp := valutil.Ptr(f)
		assert.Equal(t, f, *fp)
		assert.Equal(t, f, valutil.Val(fp))
	})

	t.Run("bool类型", func(t *testing.T) {
		bp := valutil.Ptr(true)
		assert.Equal(t, true, *bp)
		assert.Equal(t, true, valutil.Val(bp))
	})

	// nil指针测试
	t.Run("nil指针返回零值", func(t *testing.T) {
		var nilStr *string
		assert.Equal(t, "", valutil.Val(nilStr))

		var nilInt *int
		assert.Equal(t, 0, valutil.Val(nilInt))

		var nilBool *bool
		assert.Equal(t, false, valutil.Val(nilBool))
	})

	// 自定义结构体测试
	t.Run("自定义结构体", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		p := Person{"Alice", 30}
		pp := valutil.Ptr(p)
		assert.Equal(t, p, *pp)
		assert.Equal(t, p, valutil.Val(pp))
	})
}

func TestSlicePtrAndSliceVal(t *testing.T) {
	// 基本类型切片测试
	t.Run("int切片转换", func(t *testing.T) {
		nums := []int{1, 2, 3}
		ptrs := valutil.SlicePtr(nums)

		// 验证指针切片内容
		for i := range nums {
			assert.Equal(t, nums[i], *ptrs[i])
		}

		// 验证反向转换
		back := valutil.SliceVal(ptrs)
		assert.Equal(t, nums, back)
	})

	t.Run("string切片转换", func(t *testing.T) {
		strs := []string{"a", "b", "c"}
		ptrs := valutil.SlicePtr(strs)

		for i := range strs {
			assert.Equal(t, strs[i], *ptrs[i])
		}

		back := valutil.SliceVal(ptrs)
		assert.Equal(t, strs, back)
	})

	// nil切片测试
	t.Run("nil切片处理", func(t *testing.T) {
		var nilSlice []*int
		ptrs := valutil.SlicePtr(nilSlice)
		assert.Nil(t, ptrs)

		back := valutil.SliceVal(nilSlice)
		assert.Nil(t, back)
	})

	// 包含nil元素的切片测试
	t.Run("包含nil指针的切片", func(t *testing.T) {
		ptrs := []*int{nil, valutil.Ptr(1), nil, valutil.Ptr(2)}
		vals := valutil.SliceVal(ptrs)

		expected := []int{0, 1, 0, 2}
		assert.Equal(t, expected, vals)
	})

	// 空切片测试
	t.Run("空切片处理", func(t *testing.T) {
		empty := []float64{}
		ptrs := valutil.SlicePtr(empty)
		assert.NotNil(t, ptrs)
		assert.Empty(t, ptrs)

		back := valutil.SliceVal(ptrs)
		assert.Empty(t, back)
	})

	// 自定义结构体切片测试
	t.Run("自定义结构体切片", func(t *testing.T) {
		type Point struct{ X, Y int }
		points := []Point{{1, 2}, {3, 4}}
		ptrs := valutil.SlicePtr(points)

		for i := range points {
			assert.Equal(t, points[i], *ptrs[i])
		}

		back := valutil.SliceVal(ptrs)
		assert.Equal(t, points, back)
	})
}

func TestEdgeCases(t *testing.T) {
	// 零值测试
	t.Run("零值转换", func(t *testing.T) {
		var zeroInt int
		assert.Equal(t, zeroInt, valutil.Val(&zeroInt))

		var zeroStr string
		assert.Equal(t, zeroStr, valutil.Val(&zeroStr))
	})

	// 大切片测试
	t.Run("大切片性能", func(t *testing.T) {
		large := make([]int, 10000)
		for i := range large {
			large[i] = i
		}

		ptrs := valutil.SlicePtr(large)
		assert.Len(t, ptrs, len(large))
		assert.Equal(t, 9999, *ptrs[9999])

		back := valutil.SliceVal(ptrs)
		assert.Equal(t, large, back)
	})

	// 并发安全测试
	t.Run("并发安全", func(t *testing.T) {
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s := "test"
				sp := valutil.Ptr(s)
				assert.Equal(t, s, valutil.Val(sp))
			}()
		}
		wg.Wait()
	})
}
