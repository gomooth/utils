package valutil

import (
	"reflect"
	"unsafe"
)

// IsNil 检查值是否为nil或指向nil
// 支持指针、slice、map、channel、func和interface类型
func IsNil[T any](v T) bool {
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Ptr, reflect.Slice, reflect.Map,
		reflect.Chan, reflect.Func, reflect.Interface:
		return val.IsNil()
	case reflect.UnsafePointer:
		return unsafe.Pointer(&v) == nil
	default:
		return false
	}
}
