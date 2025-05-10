package valutil

import (
	"strconv"
	"strings"
)

// Int 将任意值转换为整型，失败则返回0
func Int[T any](value T) int {
	v, err := MustInt(value)
	if err != nil {
		return 0
	}
	return v
}

// IntWith 将任意值转换为整型，失败则返回 defaulted
func IntWith[T any](value T, defaulted int) int {
	v, err := MustInt(value)
	if err != nil {
		return defaulted
	}
	return v
}

// MustInt 将任意值转换为整型
// 如果输入值是 数字，则直接转换；
// 如果输入值是 boolean，则 true 转为 1；false 转为 0；
// 如果输入值是 string，则按字符串转换规则
// 否则抛出 ERROR
func MustInt[T any](value T) (int, error) {
	switch v := any(value).(type) {
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case uint:
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint64:
		return int(v), nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	case string:
		return str2int(v)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, ErrUnsupportedType
	}
}

// str2int 将字符串转换为整型
func str2int(s string) (int, error) {
	s = strings.TrimSpace(s)

	// 尝试解析为整数
	if i, err := strconv.Atoi(s); err == nil {
		return i, nil
	}

	// 尝试解析为浮点数
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return int(f), nil
	}

	// 尝试布尔字符串
	switch strings.ToLower(s) {
	case "true", "yes", "y":
		return 1, nil
	case "false", "no", "n":
		return 0, nil
	}

	return 0, ErrInvalidIntString
}
