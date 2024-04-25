package lingo

import "reflect"

// ToSlice converts the iterator into a slice
func (e Enumerable[T]) ToSlice() []T {
	res := []T{}
	for value := range e.iterator {
		res = append(res, value)
	}
	return res
}

// ToMap converts the iterator into a map with specific selector
func (e Enumerable[T]) ToMap(keySelector SingleSelector[T], elementSelector SingleSelector[T]) map[any]any {
	res := map[any]any{}
	for value := range e.iterator {
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
		t = make([]T, sourceValue.Len())
		for i := 0; i < sourceValue.Len(); i++ {
			t[i] = sourceValue.Index(i).Interface().(T)
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
