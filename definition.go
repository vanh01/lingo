package lingo

import "reflect"

// SingleSelector represents a function that converts an object under T type to another object
type SingleSelector[T any] func(T) any

// CombinationSelector represents a function that combines two objects under T, K type to another object
type CombinationSelector[T any, K any] func(T, K) any

// GroupBySelector represents a function that combines two objects under T, []K type to another object
type GroupBySelector[T any, K any] func(T, []K) any

// Comparer represents a function that compares two variables of type T,
// there will be 3 cases: smaller, equal, larger depending on the purpose of use
type Comparer[T any] func(T, T) bool

// GetHashCode represents a function that gets hashcode from an object
type GetHashCode[T any] func(T) any

func isInt(i any) bool {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

func isUint(i any) bool {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	}
	return false
}

func isFloat(i any) bool {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

func isNumber[T any](t T) bool {
	return isInt(t) || isUint(t) || isFloat(t)
}

// defaultLessComparer is the default compare when the first one is smaller than the second one
func defaultLessComparer[T any](t1, t2 T) bool {
	switch reflect.TypeOf(t1).Kind() {
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

// defaultMoreComparer is the default compare when the first one is greater than the second one
func defaultMoreComparer[T any](t1, t2 T) bool {
	switch reflect.TypeOf(t1).Kind() {
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
