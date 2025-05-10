package valutil_test

import (
	"testing"

	"github.com/gomooth/utils/valutil"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	t.Run("integer types", func(t *testing.T) {
		// 测试所有整数类型
		testCases := []struct {
			min, max int64
		}{
			{0, 100},        // int
			{-128, 127},     // int8
			{-32768, 32767}, // int16
			{0, 1<<31 - 1},  // int32
			{0, 1<<63 - 1},  // int64
		}

		for _, tc := range testCases {
			v := valutil.Rand(tc.min, tc.max)
			assert.True(t, v >= tc.min && v < tc.max,
				"Rand(%d, %d) = %d", tc.min, tc.max, v)
		}
	})

	t.Run("float types", func(t *testing.T) {
		// 测试浮点数类型
		v32 := valutil.Rand(float32(0.0), float32(1.0))
		assert.True(t, v32 >= 0.0 && v32 < 1.0)

		v64 := valutil.Rand(0.0, 1.0) // float64
		assert.True(t, v64 >= 0.0 && v64 < 1.0)
	})

	// 边界情况测试
	t.Run("edge cases", func(t *testing.T) {
		assert.Equal(t, 5, valutil.Rand(5, 6))

		v := valutil.Rand(0.5, 0.6)
		assert.True(t, v >= 0.5 && v < 0.6)

		assert.Equal(t, 10, valutil.Rand(10, 5))
	})
}

func TestSecureRandom(t *testing.T) {
	t.Run("integer types", func(t *testing.T) {
		// 测试所有整数类型的安全随机数
		types := []struct {
			min, max int64
		}{
			{0, 100},
			{-128, 127},
			{0, 1<<31 - 1},
		}

		for _, tc := range types {
			v, err := valutil.SecureRand(tc.min, tc.max)
			assert.NoError(t, err)
			assert.True(t, v >= tc.min && v < tc.max)
		}
	})

	t.Run("float types", func(t *testing.T) {
		// 测试浮点数类型的安全随机数
		v, err := valutil.SecureRand(0.0, 1.0)
		assert.NoError(t, err)
		assert.True(t, v >= 0.0 && v < 1.0)

		v32, err := valutil.SecureRand(float32(0.0), float32(1.0))
		assert.NoError(t, err)
		assert.True(t, v32 >= 0.0 && v32 < 1.0)
	})

	t.Run("error cases", func(t *testing.T) {
		// 无效范围测试
		_, err := valutil.SecureRand(10, 5)
		assert.Error(t, err)
	})
}

func TestRandString(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		s := valutil.RandAlphanum(10)
		assert.Len(t, s, 10)
		assert.Regexp(t, `^[a-zA-Z0-9]{10}$`, s)
	})

	t.Run("length variations", func(t *testing.T) {
		assert.Len(t, valutil.RandAlphanum(0), 0)
		assert.Len(t, valutil.RandAlphanum(1), 1)
		assert.Len(t, valutil.RandAlphanum(100), 100)
	})
}

func TestRandSecureString(t *testing.T) {
	t.Run("character set", func(t *testing.T) {
		s := valutil.RandSafeAlphanum(20)
		assert.Len(t, s, 20)
		// 验证不包含易混淆字符
		assert.NotContains(t, s, "l")
		assert.NotContains(t, s, "I")
		assert.NotContains(t, s, "1")
		assert.NotContains(t, s, "0")
		assert.NotContains(t, s, "O")
		assert.NotContains(t, s, "o")
		assert.NotContains(t, s, "Z")
		assert.NotContains(t, s, "z")
		assert.NotContains(t, s, "2")
	})

	t.Run("uniqueness", func(t *testing.T) {
		// 简单测试生成的字符串不相同
		s1 := valutil.RandSafeAlphanum(10)
		s2 := valutil.RandSafeAlphanum(10)
		assert.NotEqual(t, s1, s2)
	})
}

func TestRandomCoverage(t *testing.T) {
	// 测试所有数字类型的覆盖
	t.Run("all number types", func(t *testing.T) {
		assert.NotPanics(t, func() { valutil.Rand(int8(0), int8(100)) })
		assert.NotPanics(t, func() { valutil.Rand(uint(0), uint(100)) })
		assert.NotPanics(t, func() { valutil.Rand(float32(0), float32(1)) })
	})

	// 测试安全随机数的所有数字类型
	t.Run("all secure number types", func(t *testing.T) {
		_, err := valutil.SecureRand(int8(0), int8(100))
		assert.NoError(t, err)

		_, err = valutil.SecureRand(uint16(0), uint16(1000))
		assert.NoError(t, err)

		_, err = valutil.SecureRand(float32(0), float32(1))
		assert.NoError(t, err)
	})
}
