package sliceutil_test

import (
	"testing"

	"github.com/gomooth/utils/sliceutil"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"all unique", []int{1, 2, 3}, []int{1, 2, 3}},
		{"all duplicates", []int{1, 1, 1}, []int{1}},
		{"mixed duplicates", []int{1, 2, 1, 3, 2}, []int{1, 2, 3}},
		{"with zero value", []int{0, 1, 0, 2}, []int{0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceutil.Unique(tt.input...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUnique_String(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"empty strings", []string{"", ""}, []string{""}},
		{"mixed strings", []string{"a", "b", "a", "c"}, []string{"a", "b", "c"}},
		{"unicode strings", []string{"世界", "hello", "世界"}, []string{"世界", "hello"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceutil.Unique(tt.input...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUnique_CustomType(t *testing.T) {
	type point struct{ x, y int }
	tests := []struct {
		name     string
		input    []point
		expected []point
	}{
		{
			"struct duplicates",
			[]point{{1, 2}, {3, 4}, {1, 2}},
			[]point{{1, 2}, {3, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceutil.Unique(tt.input...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUniqueUnordered(t *testing.T) {
	t.Run("order not guaranteed", func(t *testing.T) {
		input := []int{3, 1, 2, 3, 1}
		result := sliceutil.UniqueUnordered(input...)

		// 只检查长度和内容，不检查顺序
		assert.Len(t, result, 3)
		assert.ElementsMatch(t, []int{1, 2, 3}, result)
	})
}

func TestUnique_PointerType(t *testing.T) {
	a, b := 1, 2
	tests := []struct {
		name     string
		input    []*int
		expected []*int
	}{
		{
			"pointer duplicates",
			[]*int{&a, &b, &a},
			[]*int{&a, &b},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceutil.Unique(tt.input...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUnique_Performance(t *testing.T) {
	// 生成包含10000个元素的切片，其中50%重复
	input := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		input[i] = i % 5000
	}

	t.Run("large slice", func(t *testing.T) {
		result := sliceutil.Unique(input...)
		assert.Len(t, result, 5000)
	})
}
