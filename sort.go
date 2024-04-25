package lingo

import (
	"reflect"
	"sort"
)

type sorter[T any, K any] struct {
	origin   []T
	keys     []K
	comparer Comparer[K]
}

func (s sorter[T, K]) Len() int {
	return len(s.keys)
}

func (s sorter[T, K]) Swap(i, j int) {
	s.origin[i], s.origin[j] = s.origin[j], s.origin[i]
	s.keys[i], s.keys[j] = s.keys[j], s.keys[i]
}

func (s sorter[T, K]) Less(i, j int) bool {
	return s.comparer(s.keys[i], s.keys[j])
}

func defaultSort[T any](t1, t2 T) bool {
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

func NewSorter[T any, K any](origin []T, source []K, comparer Comparer[K]) sorter[T, K] {
	res := sorter[T, K]{
		origin: origin,
		keys:   source,
	}
	var comp Comparer[K] = comparer
	if comparer == nil {
		comp = func(t1, t2 K) bool {
			return defaultSort(t1, t2)
		}
	}

	res.comparer = comp
	return res
}

// OrderBy sorts values in ascending order.
//
// in this method, comparer is returns whether left is smaller than right or not
//
// if comparer is nill, we will use the default comparer
func (e Enumerable[T]) OrderBy(selector SingleSelector[T], comparer Comparer[any]) Enumerable[T] {
	out := make(chan T)

	go func() {
		defer close(out)
		origin := e.ToSlice()
		source := AsEnumerable(origin).Select(selector).ToSlice()
		sorter := NewSorter(origin, source, comparer)
		sort.Sort(sorter)
		for _, value := range sorter.origin {
			out <- value
		}
	}()

	return Enumerable[T]{
		iterator: out,
	}
}

// OrderByDescending sorts values in descending order.
//
// in this method, comparer is returns whether left is smaller than right or not
func (e Enumerable[T]) OrderByDescending(selector SingleSelector[T], comparer Comparer[any]) Enumerable[T] {
	out := make(chan T)

	go func() {
		defer close(out)
		origin := e.ToSlice()
		if len(origin) == 0 {
			return
		}
		source := AsEnumerable(origin).Select(selector).ToSlice()
		sorter := NewSorter(origin, source, comparer)
		oldComparer := sorter.comparer
		sorter.comparer = func(a1, a2 any) bool {
			return !oldComparer(a1, a2)
		}
		sort.Sort(sorter)
		for _, value := range sorter.origin {
			out <- value
		}
	}()

	return Enumerable[T]{
		iterator: out,
	}
}

// Reverse reverses the order of the elements in a collection.
func (e Enumerable[T]) Reverse() Enumerable[T] {
	out := make(chan T)

	go func() {
		defer close(out)
		slice := e.ToSlice()
		for i := len(slice) - 1; i >= 0; i-- {
			out <- slice[i]
		}
	}()

	return Enumerable[T]{
		iterator: out,
	}
}
