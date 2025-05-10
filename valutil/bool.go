package valutil

import (
	"strconv"
	"strings"
)

// Bool 将任意值转换为布尔值，失败则返回 false
func Bool[T any](value T) bool {
	v, err := MustBool(value)
	if err != nil {
		return false
	}
	return v
}

// BoolWith 将任意值转换为布尔值，失败则返回 defaulted
func BoolWith[T any](value T, defaulted bool) bool {
	v, err := MustBool(value)
	if err != nil {
		return defaulted
	}
	return v
}

// MustBool 将任意值转换为布尔值，如果无法转化则返回err
// 如果传入值是 boolean，直接强制转换返回；
// 如果传入值是 string，则按以下规则转换：
//
//	"true, yes, y" -> true
//	"false, no, n, " -> false
//	"" -> false
//	"0.0···001 ... 1 ... ∞" -> true
//	"-∞ ... -1 ... -0.1 ... 0" -> false
//	"other word" -> ERROR
//
// 如果传入值是 数字，大于零返回 true，否则返回 false
//
//	0.0···001 ... 1 ... ∞ -> true
//	-∞ ... -1 ... -0.1 ... 0 -> false
//
// 如果传入值是其他类型，则返回 ERROR
func MustBool[T any](value T) (bool, error) {
	switch v := any(value).(type) {
	case bool:
		return v, nil
	case string:
		return str2bool(v)
	case int:
		return num2bool(v)
	case int8:
		return num2bool(v)
	case int16:
		return num2bool(v)
	case int32:
		return num2bool(v)
	case int64:
		return num2bool(v)
	case uint:
		return num2bool(v)
	case uint8:
		return num2bool(v)
	case uint16:
		return num2bool(v)
	case uint32:
		return num2bool(v)
	case uint64:
		return num2bool(v)
	case float32:
		return num2bool(v)
	case float64:
		return num2bool(v)
	default:
		return false, ErrUnsupportedType
	}
}

// str2bool 将字符串转换为布尔值
func str2bool(s string) (bool, error) {
	s = strings.ToLower(strings.TrimSpace(s))

	switch s {
	case "true", "yes", "y":
		return true, nil
	case "false", "no", "n", "":
		return false, nil
	default:
		// 尝试解析为数字
		if num, err := strconv.ParseFloat(s, 64); err == nil {
			return num > 0, nil
		}
		return false, ErrInvalidBoolString
	}
}

// num2bool 将数字转换为布尔值
func num2bool[T number](num T) (bool, error) {
	return num > 0, nil
}
