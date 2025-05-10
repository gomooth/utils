package valutil

// Ptr 返回值的指针
func Ptr[T any](v T) *T {
	return &v
}

// Val 返回指针的值，如果指针为nil则返回类型零值
func Val[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

// SlicePtr 将值切片转换为指针切片
func SlicePtr[T any](s []T) []*T {
	if s == nil {
		return nil
	}
	result := make([]*T, len(s))
	for i := range s {
		result[i] = &s[i]
	}
	return result
}

// SliceVal 将指针切片转换为值切片
func SliceVal[T any](s []*T) []T {
	if s == nil {
		return nil
	}
	result := make([]T, len(s))
	for i, p := range s {
		if p != nil {
			result[i] = *p
		}
		// 如果p为nil，result[i]将保持类型零值
	}
	return result
}
