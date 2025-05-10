package sliceutil

// Unique 返回去重后的切片，保持原始顺序
func Unique[T comparable](vals ...T) []T {
	set := make(map[T]struct{}, len(vals))
	for _, val := range vals {
		set[val] = struct{}{}
	}

	res := make([]T, 0, len(set))
	for _, val := range vals {
		if _, ok := set[val]; ok {
			res = append(res, val)
			delete(set, val) // 避免重复添加
		}
	}

	return res
}

// UniqueUnordered 返回去重后的切片，不保证顺序
func UniqueUnordered[T comparable](vals ...T) []T {
	set := make(map[T]struct{}, len(vals))
	for _, val := range vals {
		set[val] = struct{}{}
	}

	res := make([]T, 0, len(set))
	for val := range set {
		res = append(res, val)
	}

	return res
}
