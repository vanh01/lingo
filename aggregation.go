package lingo

import (
	"reflect"

	"github.com/vanh01/lingo/definition"
)

// Min returns the minimum value in a sequence of values.
//
// In this method, comparer is returns whether left is smaller than right or not.
// The left one will be returned
//
// If comparer is empty or nil, we will use the default comparer.
// On the other hand, we just use the first comparer
func (e Enumerable[T]) Min(comparer ...definition.Comparer[T]) T {
	var t T
	first := true
	for value := range e.getIter() {
		if first {
			t = value
			first = false
		}
		if definition.IsEmptyOrNil(comparer) {
			if definition.DefaultLessComparer(value, t) {
				t = value
			}
		} else if comparer[0](value, t) {
			t = value
		}
	}
	return t
}

// Max returns the minimum value in a sequence of values.
//
// In this method, comparer is returns whether left is greater than right or not.
// The left one will be returned
//
// If comparer is empty or nil, we will use the default comparer.
// On the other hand, we just use the first comparer
func (e Enumerable[T]) Max(comparer ...definition.Comparer[T]) T {
	var t T
	first := true
	for value := range e.getIter() {
		if first {
			t = value
			first = false
		}
		if definition.IsEmptyOrNil(comparer) {
			if definition.DefaultMoreComparer(value, t) {
				t = value
			}
		} else if comparer[0](value, t) {
			t = value
		}
	}
	return t
}

// Sum computes the sum of a sequence of numeric values.
//
// If selector is not empty or nil, we will use the first comparer
func (e Enumerable[T]) Sum(selector ...definition.SingleSelector[T]) any {
	var sumInt64 int64 = 0
	var sumUint64 uint64 = 0
	var sumFloat64 float64 = 0
	var temp any
	for value := range e.getIter() {
		temp = value
		if !definition.IsEmptyOrNil(selector) {
			temp = selector[0](value)
		}
		if !definition.IsNumber(temp) {
			break
		}
		switch {
		case definition.IsInt(temp):
			sumInt64 += reflect.ValueOf(temp).Int()
		case definition.IsUint(temp):
			sumUint64 += reflect.ValueOf(temp).Uint()
		case definition.IsFloat(temp):
			sumFloat64 += reflect.ValueOf(temp).Float()
		}
	}
	if definition.IsInt(temp) {
		return sumInt64
	}
	if definition.IsUint(temp) {
		return sumUint64
	}
	return sumFloat64
}

// Average computes the average of a sequence of numeric values.
//
// If selector is not empty or nil, we will use the first comparer
func (e Enumerable[T]) Average(selector ...definition.SingleSelector[T]) float64 {
	var sumInt64 int64 = 0
	var sumUint64 uint64 = 0
	var sumFloat64 float64 = 0
	var temp any
	i := 0
	for value := range e.getIter() {
		temp = value
		if !definition.IsEmptyOrNil(selector) {
			temp = selector[0](value)
		}
		if !definition.IsNumber(temp) {
			break
		}
		switch {
		case definition.IsInt(temp):
			sumInt64 += reflect.ValueOf(temp).Int()
		case definition.IsUint(temp):
			sumUint64 += reflect.ValueOf(temp).Uint()
		case definition.IsFloat(temp):
			sumFloat64 += reflect.ValueOf(temp).Float()
		}
		i++
	}
	if definition.IsInt(temp) {
		return float64(sumInt64) / float64(i)
	}
	if definition.IsUint(temp) {
		return float64(sumUint64) / float64(i)
	}
	return sumFloat64 / float64(i)
}

// Count returns the number of elements in a sequence.
func (e Enumerable[T]) Count() int64 {
	var i int64 = 0
	for value := range e.getIter() {
		_ = value
		i++
	}
	return i
}

// Aggregate applies an accumulator function over a sequence.
// The specified seed value is used as the initial accumulator value,
// and the specified function is used to select the result value.
//
// # If resultSelector is not empty or nil, we will use the first comparer
//
// # Noted that the type of seed, the left type of Accumulator function and the input type of selector must be the same
func (e Enumerable[T]) Aggregate(
	seed any,
	accumulator definition.Accumulator[any, T],
	resultSelector ...definition.SingleSelector[any],
) any {
	var res any = seed
	for value := range e.getIter() {
		res = accumulator(res, value)
	}
	if !definition.IsEmptyOrNil(resultSelector) {
		return resultSelector[0](res)
	}
	return res
}
