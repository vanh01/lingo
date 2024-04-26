package lingo

import "reflect"

// Min returns the minimum value in a sequence of values.
//
// in this method, comparer is returns whether left is smaller than right or not
//
// if comparer is nill, we will use the default comparer
func (e Enumerable[T]) Min(comparer Comparer[T]) T {
	var t T
	first := true
	for value := range e.iterator {
		if first {
			t = value
			first = false
		}
		if comparer == nil {
			if defaultLessComparer(value, t) {
				t = value
			}
		} else if comparer(value, t) {
			t = value
		}
	}
	return t
}

// Max returns the minimum value in a sequence of values.
//
// in this method, comparer is returns whether left is greater than right or not
//
// if comparer is nill, we will use the default comparer
func (e Enumerable[T]) Max(comparer Comparer[T]) T {
	var t T
	first := true
	for value := range e.iterator {
		if first {
			t = value
			first = false
		}
		if comparer == nil {
			if defaultMoreComparer(value, t) {
				t = value
			}
		} else if comparer(value, t) {
			t = value
		}
	}
	return t
}

// Sum computes the sum of a sequence of numeric values.
func (e Enumerable[T]) Sum(selector SingleSelector[T]) any {
	var sumInt64 int64 = 0
	var sumUint64 uint64 = 0
	var sumFloat64 float64 = 0
	var temp any
	for value := range e.iterator {
		temp = value
		if selector != nil {
			temp = selector(value)
		}
		if !isNumber(temp) {
			break
		}
		switch {
		case isInt(temp):
			sumInt64 += reflect.ValueOf(temp).Int()
		case isUint(temp):
			sumUint64 += reflect.ValueOf(temp).Uint()
		case isFloat(temp):
			sumFloat64 += reflect.ValueOf(temp).Float()
		}
	}
	if isInt(temp) {
		return sumInt64
	}
	if isUint(temp) {
		return sumUint64
	}
	return sumFloat64
}

// Average computes the average of a sequence of numeric values.
func (e Enumerable[T]) Average(selector SingleSelector[T]) float64 {
	var sumInt64 int64 = 0
	var sumUint64 uint64 = 0
	var sumFloat64 float64 = 0
	var temp any
	i := 0
	for value := range e.iterator {
		temp = value
		if selector != nil {
			temp = selector(value)
		}
		if !isNumber(temp) {
			break
		}
		switch {
		case isInt(temp):
			sumInt64 += reflect.ValueOf(temp).Int()
		case isUint(temp):
			sumUint64 += reflect.ValueOf(temp).Uint()
		case isFloat(temp):
			sumFloat64 += reflect.ValueOf(temp).Float()
		}
		i++
	}
	if isInt(temp) {
		return float64(sumInt64) / float64(i)
	}
	if isUint(temp) {
		return float64(sumUint64) / float64(i)
	}
	return sumFloat64 / float64(i)
}

// Count returns the number of elements in a sequence.
func (e Enumerable[T]) Count() int64 {
	var i int64 = 0
	for value := range e.iterator {
		_ = value
		i++
	}
	return i
}

// Aggregate applies an accumulator function over a sequence.
// The specified seed value is used as the initial accumulator value,
// and the specified function is used to select the result value.
//
// resultSelector can be nil
//
// # Noted that the type of seed, the left type of Accumulator function and the input type of selector must be the same
func (e Enumerable[T]) Aggregate(
	seed any,
	accumulator Accumulator[any, T],
	resultSelector SingleSelector[any],
) any {
	var res any = seed
	for value := range e.iterator {
		res = accumulator(res, value)
	}
	if resultSelector != nil {
		return resultSelector(res)
	}
	return res
}
