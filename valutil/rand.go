package valutil

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand/v2"
)

// Rand 生成 [min, max) 范围内的随机数
func Rand[T number](min, max T) T {
	if max < min {
		return min
	}
	switch any(min).(type) {
	case float32, float64:
		return T(rand.Float64()*(float64(max)-float64(min)) + float64(min))
	default:
		return T(rand.Int64N(int64(max)-int64(min)) + int64(min))
	}
}

// SecureRand 生成 [min, max) 范围内的加密安全随机数
func SecureRand[T number](min, max T) (T, error) {
	if min > max {
		return 0, fmt.Errorf("min > max error")
	}
	switch any(min).(type) {
	case float32, float64:
		t, err := secureRandFloat(float64(min), float64(max))
		return T(t), err
	default:
		t, err := secureRandInt(int64(min), int64(max))
		return T(t), err
	}
}

// secureRandInt 生成 [min, max) 范围内整数类型的安全随机数
func secureRandInt(min, max int64) (int64, error) {
	bigN := new(big.Int).SetInt64(max - min)
	n, err := crand.Int(crand.Reader, bigN)
	if err != nil {
		return 0, err
	}
	return n.Int64() + min, nil
}

// secureRandFloat 生成 [min, max) 范围内浮点数类型的安全随机数
func secureRandFloat(min, max float64) (float64, error) {
	// 生成 [0, 1<<53) 的大随机整数 (保证浮点数精度)
	bigN := new(big.Int).Lsh(big.NewInt(1), 53)
	n, err := crand.Int(crand.Reader, bigN)
	if err != nil {
		return 0, err
	}

	// 转换为 [0,1) 范围的浮点数
	f := float64(n.Int64()) / (1 << 53)

	// 调整到 [min, max) 范围
	return min + f*(max-min), nil
}

// RandAlphanum 生成固定长度的随机字符串
func RandAlphanum(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.IntN(len(letterBytes))]
	}
	return string(b)
}

// RandSafeAlphanum 生成固定长度的随机易辨识的字符串
// 适合用于验证码、密码等需要人工识别的场景
func RandSafeAlphanum(n int) string {
	const letterBytes = "abcdefghijkmnpqrstuvwxyABCDEFGHJKLMNPQRSTUVWXY3456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.IntN(len(letterBytes))]
	}
	return string(b)
}
