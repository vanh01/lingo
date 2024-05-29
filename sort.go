package lingo

import (
	"sort"

	"github.com/vanh01/lingo/definition"
)

type sorter[T any, K any] struct {
	origin   []T
	keys     []K
	comparer definition.Comparer[K]
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

func NewSorter[T any, K any](origin []T, source []K, comparer ...definition.Comparer[K]) sorter[T, K] {
	res := sorter[T, K]{
		origin: origin,
		keys:   source,
	}
	var comp definition.Comparer[K]
	if definition.IsEmptyOrNil(comparer) {
		comp = func(t1, t2 K) bool {
			return definition.DefaultLessComparer(t1, t2)
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
func (e Enumerable[T]) OrderBy(selector definition.SingleSelector[T], comparer ...definition.Comparer[any]) Enumerable[T] {
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
func (e Enumerable[T]) OrderByDescending(selector definition.SingleSelector[T], comparer ...definition.Comparer[any]) Enumerable[T] {
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

// ParallelEnumerable

// OrderBy sorts in parallel values in ascending order. And then set ordered is true
//
// In this method, comparer is returns whether left is smaller than right or not.
// If comparer is empty or nil, we will use the default comparer. On the other hand, we just use the first comparer
func (p ParallelEnumerable[T]) OrderBy(selector definition.SingleSelector[T], comparer ...definition.Comparer[any]) ParallelEnumerable[T] {
	return ParallelEnumerable[T]{
		wasSetUnordered: true,
		ordered:         true,
		getIter: func() <-chan odata[T] {
			out := make(chan odata[T])

			go func() {
				defer close(out)
				origin := p.ToSlice()
				source := AsParallelEnumerable(origin).AsOrdered().Select(selector).ToSlice()
				sorter := NewSorter(origin, source, comparer...)
				sort.Sort(sorter)
				i := 0
				for _, value := range sorter.origin {
					out <- odata[T]{
						no:  i,
						val: value,
					}
					i++
				}
			}()

			return out
		},
	}
}

// OrderByDescending sorts in parallel values in descending order. And then set ordered is true
//
// In this method, comparer is returns whether left is smaller than right or not.
// If comparer is empty or nil, we will use the default comparer. On the other hand, we just use the first comparer
func (p ParallelEnumerable[T]) OrderByDescending(selector definition.SingleSelector[T], comparer ...definition.Comparer[any]) ParallelEnumerable[T] {
	return ParallelEnumerable[T]{
		wasSetUnordered: true,
		ordered:         true,
		getIter: func() <-chan odata[T] {
			out := make(chan odata[T])

			go func() {
				defer close(out)
				origin := p.ToSlice()
				if len(origin) == 0 {
					return
				}
				source := AsParallelEnumerable(origin).AsOrdered().Select(selector).ToSlice()
				sorter := NewSorter(origin, source, comparer...)
				oldComparer := sorter.comparer
				sorter.comparer = func(a1, a2 any) bool {
					return !oldComparer(a1, a2)
				}
				sort.Sort(sorter)
				i := 0
				for _, value := range sorter.origin {
					out <- odata[T]{
						no:  i,
						val: value,
					}
					i++
				}
			}()

			return out
		},
	}
}

// Reverse inverts the order of the elements in a parallel sequence.
// If this one is Unordered, does nothing.
func (p ParallelEnumerable[T]) Reverse() ParallelEnumerable[T] {
	if !p.ordered {
		return p
	}

	return ParallelEnumerable[T]{
		wasSetUnordered: p.wasSetUnordered,
		ordered:         p.ordered,
		getIter: func() <-chan odata[T] {
			out := make(chan odata[T])

			go func() {
				defer close(out)
				slice := p.ToSlice()
				index := 0
				for i := len(slice) - 1; i >= 0; i-- {
					out <- odata[T]{
						no:  index,
						val: slice[i],
					}
					index++
				}
			}()

			return out
		},
	}
}
