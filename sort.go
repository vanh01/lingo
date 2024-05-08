package lingo

import (
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

func NewSorter[T any, K any](origin []T, source []K, comparer ...Comparer[K]) sorter[T, K] {
	res := sorter[T, K]{
		origin: origin,
		keys:   source,
	}
	var comp Comparer[K]
	if isEmptyOrNil(comparer) {
		comp = func(t1, t2 K) bool {
			return defaultLessComparer(t1, t2)
		}
	} else {
		comp = comparer[0]
	}

	res.comparer = comp
	return res
}

// OrderBy sorts values in ascending order.
//
// In this method, comparer is returns whether left is smaller than right or not.
// If comparer is empty or nil, we will use the default comparer. On the other hand, we just use the first comparer
func (e Enumerable[T]) OrderBy(selector SingleSelector[T], comparer ...Comparer[any]) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				origin := e.ToSlice()
				source := AsEnumerable(origin).Select(selector).ToSlice()
				sorter := NewSorter(origin, source, comparer...)
				sort.Sort(sorter)
				for _, value := range sorter.origin {
					out <- value
				}
			}()

			return out
		},
	}
}

// OrderByDescending sorts values in descending order.
//
// In this method, comparer is returns whether left is smaller than right or not.
// If comparer is empty or nil, we will use the default comparer. On the other hand, we just use the first comparer
func (e Enumerable[T]) OrderByDescending(selector SingleSelector[T], comparer ...Comparer[any]) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				origin := e.ToSlice()
				if len(origin) == 0 {
					return
				}
				source := AsEnumerable(origin).Select(selector).ToSlice()
				sorter := NewSorter(origin, source, comparer...)
				oldComparer := sorter.comparer
				sorter.comparer = func(a1, a2 any) bool {
					return !oldComparer(a1, a2)
				}
				sort.Sort(sorter)
				for _, value := range sorter.origin {
					out <- value
				}
			}()

			return out
		},
	}
}

// Reverse reverses the order of the elements in a collection.
func (e Enumerable[T]) Reverse() Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				slice := e.ToSlice()
				for i := len(slice) - 1; i >= 0; i-- {
					out <- slice[i]
				}
			}()

			return out
		},
	}
}
