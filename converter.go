package lingo

import (
	"reflect"
	"unsafe"

	"github.com/vanh01/lingo/definition"
)

// ToSlice converts the iterator into a slice
func (e Enumerable[T]) ToSlice() []T {
	res := []T{}
	for value := range e.getIter() {
		res = append(res, value)
	}
	return res
}

// ToMap converts the iterator into a map with specific selector
func (e Enumerable[T]) ToMap(keySelector definition.SingleSelector[T], elementSelector definition.SingleSelector[T]) map[any]any {
	res := map[any]any{}
	for value := range e.getIter() {
		res[keySelector(value)] = elementSelector(value)
	}
	return res
}

// SliceAnyToT converts from slice of any to slice of specific type
func SliceAnyToT[T any](source interface{}) []T {
	var t []T
	switch reflect.TypeOf(source).Kind() {
	case reflect.Slice:
		sourceValue := reflect.ValueOf(source)
		var temp T
		t = make([]T, sourceValue.Len())
		for i := 0; i < sourceValue.Len(); i++ {
			value := sourceValue.Index(i).Interface()
			switch {
			case definition.IsNumber(temp):
				t[i] = definition.DefaultConvertToNumber[T](value)
			case reflect.TypeOf(temp) == reflect.TypeOf(value):
				t[i] = value.(T)
			default:
				t[i] = *(*T)(unsafe.Pointer(&value))
			}
		}
	}
	return t
}

// SliceTToAny converts from slice of specific type to slice of any
func SliceTToAny[T any](source []T) []any {
	var ins []any = make([]any, len(source))
	for i := 0; i < len(source); i++ {
		ins[i] = source[i]
	}
	return ins
}

// ParallelEnumerable

// ToSlice converts the iterator into a slice
func (p ParallelEnumerable[T]) ToSlice() []T {
	res := []T{}
	temp := p
	if p.ordered {
		temp = temp.order()
	}
	for data := range temp.getIter() {
		res = append(res, data.val)
	}
	return res
}

// ToMap converts the iterator into a map with specific selector
func (p ParallelEnumerable[T]) ToMap(keySelector definition.SingleSelector[T], elementSelector definition.SingleSelector[T]) map[any]any {
	res := map[any]any{}
	temp := p
	if p.ordered {
		temp = temp.order()
	}
	for value := range temp.getIter() {
		res[keySelector(value.val)] = elementSelector(value.val)
	}
	return res
}
