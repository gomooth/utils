package valutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNil(t *testing.T) {
	// 定义测试用的自定义类型
	type MyStruct struct{}
	type MyInterface interface{}

	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		// nil 指针测试
		{"nil *int", (*int)(nil), true},
		{"nil *struct", (*MyStruct)(nil), true},
		{"non-nil *int", new(int), false},
		{"non-nil *struct", &MyStruct{}, false},

		// slice测试
		{"nil slice", ([]int)(nil), true},
		{"empty slice", []int{}, false},
		{"non-empty slice", []int{1}, false},

		// map测试
		{"nil map", (map[string]int)(nil), true},
		{"empty map", map[string]int{}, false},
		{"non-empty map", map[string]int{"a": 1}, false},

		// channel测试
		{"nil channel", (chan int)(nil), true},
		{"initialized channel", make(chan int), false},

		// func测试
		{"nil func", (func())(nil), true},
		{"non-nil func", func() {}, false},

		// interface测试
		{"nil interface", (interface{})(nil), true},
		{"nil custom interface", (MyInterface)(nil), true},
		{"non-nil interface", "hello", false},

		// 非nil类型测试
		{"integer", 42, false},
		{"string", "hello", false},
		{"struct", MyStruct{}, false},
		{"array", [3]int{1, 2, 3}, false},
		{"bool", true, false},
		{"float", 3.14, false},

		// 指针的指针
		{"pointer to nil pointer", func() any {
			var nilPtr *int
			return &nilPtr
		}(), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNil(tt.input)
			assert.Equal(t, tt.expected, result,
				"IsNil(%v) should be %v", tt.input, tt.expected)
		})
	}
}
