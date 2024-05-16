package lingo

import (
	"reflect"
	"sync"

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

// MinBy returns the minimum value in a sequence of values according to a specified key selector function.
//
// In this method, comparer is returns whether left is smaller than right or not.
// The left one will be returned
//
// If comparer is empty or nil, we will use the default comparer.
// On the other hand, we just use the first comparer
func (e Enumerable[T]) MinBy(selector definition.SingleSelector[T], comparer ...definition.Comparer[any]) T {
	var t T
	var minKey any
	first := true
	for value := range e.getIter() {
		key := selector(value)
		if first {
			t = value
			minKey = key
			first = false
		}
		if definition.IsEmptyOrNil(comparer) {
			if definition.DefaultLessComparer(key, minKey) {
				t = value
				minKey = key
			}
		} else if comparer[0](key, minKey) {
			t = value
			minKey = key
		}
	}
	return t
}

// Max returns the maximum value in a sequence of values.
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

// MaxBy returns the maximum value in a sequence of values according to a specified key selector function.
//
// In this method, comparer is returns whether left is smaller than right or not.
// The left one will be returned
//
// If comparer is empty or nil, we will use the default comparer.
// On the other hand, we just use the first comparer
func (e Enumerable[T]) MaxBy(selector definition.SingleSelector[T], comparer ...definition.Comparer[any]) T {
	var t T
	var maxKey any
	first := true
	for value := range e.getIter() {
		key := selector(value)
		if first {
			t = value
			maxKey = key
			first = false
		}
		if definition.IsEmptyOrNil(comparer) {
			if definition.DefaultMoreComparer(key, maxKey) {
				t = value
				maxKey = key
			}
		} else if comparer[0](key, maxKey) {
			t = value
			maxKey = key
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

// ParallelEnumerable

// MinBy invokes in parallel a transform function on each element of a sequence and returns the minimum value.
//
// In this method, comparer is returns whether left is smaller than right or not.
// The left one will be returned
//
// If comparer is empty or nil, we will use the default comparer.
// On the other hand, we just use the first comparer
func (p ParallelEnumerable[T]) MinBy(selector definition.SingleSelector[T], comparer ...definition.Comparer[any]) T {
	var t T
	keyVal := make(chan definition.KeyValData[any, T])
	first := true

	go func() {
		defer close(keyVal)

		var wg sync.WaitGroup
		for value := range p.getIter() {
			wg.Add(1)
			temp := value
			go func() {
				defer wg.Done()
				keyVal <- definition.KeyValData[any, T]{
					Key: selector(temp.val),
					Val: temp.val,
				}
			}()
		}
		wg.Wait()
	}()

	var minKey any
	for key := range keyVal {
		if first {
			t = key.Val
			minKey = key.Key
			first = false
		}
		if definition.IsEmptyOrNil(comparer) {
			if definition.DefaultLessComparer(key.Key, minKey) {
				t = key.Val
				minKey = key.Key
			}
		} else if comparer[0](key.Key, minKey) {
			t = key.Val
			minKey = key.Key
		}
	}
	return t
}

// MaxBy invokes in parallel a transform function on each element of a sequence and returns the maximum value.
//
// In this method, comparer is returns whether left is smaller than right or not.
// The left one will be returned
//
// If comparer is empty or nil, we will use the default comparer.
// On the other hand, we just use the first comparer
func (p ParallelEnumerable[T]) MaxBy(selector definition.SingleSelector[T], comparer ...definition.Comparer[any]) T {
	var t T
	keyVal := make(chan definition.KeyValData[any, T])
	first := true

	go func() {
		defer close(keyVal)

		var wg sync.WaitGroup
		for value := range p.getIter() {
			wg.Add(1)
			temp := value
			go func() {
				defer wg.Done()
				keyVal <- definition.KeyValData[any, T]{
					Key: selector(temp.val),
					Val: temp.val,
				}
			}()
		}
		wg.Wait()
	}()

	var maxKey any
	for key := range keyVal {
		if first {
			t = key.Val
			maxKey = key.Key
			first = false
		}
		if definition.IsEmptyOrNil(comparer) {
			if definition.DefaultMoreComparer(key.Key, maxKey) {
				t = key.Val
				maxKey = key.Key
			}
		} else if comparer[0](key.Key, maxKey) {
			t = key.Val
			maxKey = key.Key
		}
	}
	return t
}

// Sum computes in parallel the sum of the sequence of values that are obtained
// by invoking a transform function on each element of the input sequence.
//
// If selector is not empty or nil, we will use the first comparer
func (p ParallelEnumerable[T]) Sum(selector ...definition.SingleSelector[T]) any {
	out := make(chan any)

	go func() {
		defer close(out)
		var wg sync.WaitGroup
		for value := range p.getIter() {
			wg.Add(1)
			temp := value
			go func() {
				defer wg.Done()
				if definition.IsEmptyOrNil(selector) {
					out <- temp
				} else {
					out <- selector[0](temp.val)
				}
			}()
		}
		wg.Wait()
	}()

	var sumInt64 int64 = 0
	var sumUint64 uint64 = 0
	var sumFloat64 float64 = 0
	var temp any
	for value := range out {
		temp = value
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

// Average computes in parallel the average of a sequence of numeric values
// that are obtained by invoking a transform function on each element of the input sequence.
//
// If selector is not empty or nil, we will use the first comparer
func (p ParallelEnumerable[T]) Average(selector ...definition.SingleSelector[T]) float64 {
	out := make(chan any)

	go func() {
		defer close(out)
		var wg sync.WaitGroup
		for value := range p.getIter() {
			wg.Add(1)
			temp := value
			go func() {
				defer wg.Done()
				if definition.IsEmptyOrNil(selector) {
					out <- temp
				} else {
					out <- selector[0](temp.val)
				}
			}()
		}
		wg.Wait()
	}()

	var sumInt64 int64 = 0
	var sumUint64 uint64 = 0
	var sumFloat64 float64 = 0
	var temp any
	i := 0
	for value := range out {
		temp = value
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
