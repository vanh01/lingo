package definition

import (
	"reflect"
)

// SingleSelector represents a function that converts an object under T type to another object
type SingleSelector[T any] func(T) any

// SingleSelectorFull represents a function that converts an object under T type to another object under K type
type SingleSelectorFull[T any, K any] func(T) K

// CombinationSelector represents a function that combines two objects under T, K type to another object
type CombinationSelector[T any, K any] func(T, K) any

// GroupBySelector represents a function that combines two objects under T, []K type to another object
type GroupBySelector[T any, K any] func(T, []K) any

// Comparer represents a function that compares two variables of type T,
// there will be 3 cases: smaller, equal, larger depending on the purpose of use
type Comparer[T any] func(T, T) bool

// GetHashCode represents a function that gets hashcode from an object
type GetHashCode[T any] func(T) any

// Accumulator represents an accumulator function to be invoked on each element.
type Accumulator[T any, K any] func(T, K) T

func IsInt(i any) bool {
	switch reflect.ValueOf(i).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

func IsUint(i any) bool {
	switch reflect.ValueOf(i).Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	}
	return false
}

func IsFloat(i any) bool {
	switch reflect.ValueOf(i).Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

func IsNumber(i any) bool {
	return IsInt(i) || IsUint(i) || IsFloat(i)
}

// convert number type to int
func floatToActualInt[T any](value any) T {
	var t T
	valueFloat64 := reflect.ValueOf(value).Float()
	switch reflect.ValueOf(t).Kind() {
	case reflect.Int:
		t = reflect.ValueOf(int(valueFloat64)).Interface().(T)
	case reflect.Int8:
		t = reflect.ValueOf(int8(valueFloat64)).Interface().(T)
	case reflect.Int16:
		t = reflect.ValueOf(int16(valueFloat64)).Interface().(T)
	case reflect.Int32:
		t = reflect.ValueOf(int32(valueFloat64)).Interface().(T)
	case reflect.Int64:
		t = reflect.ValueOf(int64(valueFloat64)).Interface().(T)
	}
	return t
}
func uintToActualInt[T any](value any) T {
	var t T
	valueUint64 := reflect.ValueOf(value).Uint()
	switch reflect.ValueOf(t).Kind() {
	case reflect.Int:
		t = reflect.ValueOf(int(valueUint64)).Interface().(T)
	case reflect.Int8:
		t = reflect.ValueOf(int8(valueUint64)).Interface().(T)
	case reflect.Int16:
		t = reflect.ValueOf(int16(valueUint64)).Interface().(T)
	case reflect.Int32:
		t = reflect.ValueOf(int32(valueUint64)).Interface().(T)
	case reflect.Int64:
		t = reflect.ValueOf(int64(valueUint64)).Interface().(T)
	}
	return t
}
func intToActualInt[T any](value any) T {
	var t T
	valueInt64 := reflect.ValueOf(value).Int()
	switch reflect.ValueOf(t).Kind() {
	case reflect.Int:
		t = reflect.ValueOf(int(valueInt64)).Interface().(T)
	case reflect.Int8:
		t = reflect.ValueOf(int8(valueInt64)).Interface().(T)
	case reflect.Int16:
		t = reflect.ValueOf(int16(valueInt64)).Interface().(T)
	case reflect.Int32:
		t = reflect.ValueOf(int32(valueInt64)).Interface().(T)
	case reflect.Int64:
		t = reflect.ValueOf(valueInt64).Interface().(T)
	}
	return t
}

// convert number type to int
func floatToActualUint[T any](value any) T {
	var t T
	valueFloat64 := reflect.ValueOf(value).Float()
	switch reflect.ValueOf(t).Kind() {
	case reflect.Uint:
		t = reflect.ValueOf(uint(valueFloat64)).Interface().(T)
	case reflect.Uint8:
		t = reflect.ValueOf(uint8(valueFloat64)).Interface().(T)
	case reflect.Uint16:
		t = reflect.ValueOf(uint16(valueFloat64)).Interface().(T)
	case reflect.Uint32:
		t = reflect.ValueOf(uint32(valueFloat64)).Interface().(T)
	case reflect.Uint64:
		t = reflect.ValueOf(uint64(valueFloat64)).Interface().(T)
	}
	return t
}
func uintToActualUint[T any](value any) T {
	var t T
	valueUint64 := reflect.ValueOf(value).Uint()
	switch reflect.ValueOf(t).Kind() {
	case reflect.Uint:
		t = reflect.ValueOf(uint(valueUint64)).Interface().(T)
	case reflect.Uint8:
		t = reflect.ValueOf(uint8(valueUint64)).Interface().(T)
	case reflect.Uint16:
		t = reflect.ValueOf(uint16(valueUint64)).Interface().(T)
	case reflect.Uint32:
		t = reflect.ValueOf(uint32(valueUint64)).Interface().(T)
	case reflect.Uint64:
		t = reflect.ValueOf(valueUint64).Interface().(T)
	}
	return t
}
func intToActualUint[T any](value any) T {
	var t T
	valueInt64 := reflect.ValueOf(value).Int()
	switch reflect.ValueOf(t).Kind() {
	case reflect.Uint:
		t = reflect.ValueOf(uint(valueInt64)).Interface().(T)
	case reflect.Uint8:
		t = reflect.ValueOf(uint8(valueInt64)).Interface().(T)
	case reflect.Uint16:
		t = reflect.ValueOf(uint16(valueInt64)).Interface().(T)
	case reflect.Uint32:
		t = reflect.ValueOf(uint32(valueInt64)).Interface().(T)
	case reflect.Uint64:
		t = reflect.ValueOf(uint64(valueInt64)).Interface().(T)
	}
	return t
}

// convert number type to float
func intToActualFloat[T any](value any) T {
	var t T
	valueInt64 := reflect.ValueOf(value).Int()
	switch reflect.ValueOf(t).Kind() {
	case reflect.Float32:
		t = reflect.ValueOf(float32(valueInt64)).Interface().(T)
	case reflect.Float64:
		t = reflect.ValueOf(float64(valueInt64)).Interface().(T)
	}
	return t
}
func uintToActualFloat[T any](value any) T {
	var t T
	valueInt64 := reflect.ValueOf(value).Uint()
	switch reflect.ValueOf(t).Kind() {
	case reflect.Float32:
		t = reflect.ValueOf(float32(valueInt64)).Interface().(T)
	case reflect.Float64:
		t = reflect.ValueOf(float64(valueInt64)).Interface().(T)
	}
	return t
}
func floatToActualFloat[T any](value any) T {
	var t T
	valueFloat64 := reflect.ValueOf(value).Float()
	switch reflect.ValueOf(t).Kind() {
	case reflect.Float32:
		t = reflect.ValueOf(float32(valueFloat64)).Interface().(T)
	case reflect.Float64:
		t = reflect.ValueOf(valueFloat64).Interface().(T)
	}
	return t
}

// defaultConvertToInt
func defaultConvertToInt[T any](value any) T {
	var t T

	switch {
	case IsFloat(value):
		t = floatToActualInt[T](value)
	case IsUint(value):
		t = uintToActualInt[T](value)
	case IsInt(value):
		t = intToActualInt[T](value)
	}
	return t
}

// defaultConvertToUint
func defaultConvertToUint[T any](value any) T {
	var t T

	switch {
	case IsFloat(value):
		t = floatToActualUint[T](value)
	case IsUint(value):
		t = uintToActualUint[T](value)
	case IsInt(value):
		t = intToActualUint[T](value)
	}
	return t
}

// defaultConvertToFloat
func defaultConvertToFloat[T any](value any) T {
	var t T

	switch {
	case IsInt(value):
		t = intToActualFloat[T](value)
	case IsUint(value):
		t = uintToActualFloat[T](value)
	case IsFloat(value):
		t = floatToActualFloat[T](value)
	}
	return t
}

// DefaultConvertToNumber
func DefaultConvertToNumber[T any](value any) T {
	var t T
	switch {
	case IsInt(t):
		t = defaultConvertToInt[T](value)
	case IsUint(t):
		t = defaultConvertToUint[T](value)
	case IsFloat(t):
		t = defaultConvertToFloat[T](value)
	}
	return t
}

// DefaultLessComparer is the default compare when the first one is smaller than the second one
func DefaultLessComparer[T any](t1, t2 T) bool {
	switch reflect.ValueOf(t1).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(t1).Int() < reflect.ValueOf(t2).Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.ValueOf(t1).Uint() < reflect.ValueOf(t2).Uint()
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(t1).Float() < reflect.ValueOf(t2).Float()
	case reflect.String:
		return reflect.ValueOf(t1).String() < reflect.ValueOf(t2).String()
	case reflect.Bool:
		// t1 < t2 when: t1 is false, and t2 is true
		v1, v2 := reflect.ValueOf(t1).Bool(), reflect.ValueOf(t2).Bool()
		if !v1 && v2 {
			return true
		}
		return false
	}
	return false
}

// DefaultMoreComparer is the default compare when the first one is greater than the second one
func DefaultMoreComparer[T any](t1, t2 T) bool {
	switch reflect.ValueOf(t1).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(t1).Int() > reflect.ValueOf(t2).Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.ValueOf(t1).Uint() > reflect.ValueOf(t2).Uint()
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(t1).Float() > reflect.ValueOf(t2).Float()
	case reflect.String:
		return reflect.ValueOf(t1).String() > reflect.ValueOf(t2).String()
	case reflect.Bool:
		// t1 > t2 when: t1 is true, and t2 is false
		v1, v2 := reflect.ValueOf(t1).Bool(), reflect.ValueOf(t2).Bool()
		if v1 && !v2 {
			return true
		}
		return false
	}
	return false
}

func IsEmptyOrNil[T any](t []T) bool {
	if reflect.ValueOf(t).IsNil() {
		return true
	}
	if len(t) == 0 {
		return true
	}
	if reflect.ValueOf(t[0]).IsNil() {
		return true
	}
	return false
}
