package bitutil

// unsignedBitCheckable 约束所有无符号整数类型
type unsignedBitCheckable interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// HasBitAt 检查 byte 中从右数第pos位是否为1(pos 0-based)
// pos 从0开始计数，0表示最后一位
func HasBitAt[T unsignedBitCheckable](n T, pos uint) bool {
	switch any(n).(type) {
	case uint8:
		if pos > 7 {
			return false
		}
	case uint16:
		if pos > 15 {
			return false
		}
	case uint32:
		if pos > 31 {
			return false
		}
	default:
		if pos > 63 {
			return false
		}
	}
	return uint64(n)&(1<<pos) != 0
}
