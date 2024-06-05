package lingo

import "sync"

type odata[T any] struct {
	no  int
	val T
}

// ParallelEnumerable provides a set of methods for querying objects that implement ParallelQuery[T].
// This is the parallel equivalent of Enumerable.
//
// # Note that ParallelEnumerable will be useful in some cases, you must really consider when using it.
//
// # In worse cases, not only will performance not be optimized, but costs will also increase.
type ParallelEnumerable[T any] struct {
	wasSetUnordered bool
	ordered         bool
	getIter         func() <-chan odata[T]
}

// GetIter returns an unbuffered channel of T that iterates through a collection.
func (p ParallelEnumerable[T]) GetIter() <-chan T {
	out := make(chan T)

	temp := p
	if p.ordered {
		temp = temp.order()
	}
	go func() {
		defer close(out)
		for value := range temp.getIter() {
			out <- value.val
		}
	}()

	return out
}

// AsParallel creates a new ParallelEnumerable from an Enumerable, it will need more resources
// but performance will be improved
//
// # Note that ParallelEnumerable will be useful in some cases, you must really consider when using it.
//
// # In worse cases, not only will performance not be optimized, but costs will also increase.
func (e Enumerable[T]) AsParallel() ParallelEnumerable[T] {
	return ParallelEnumerable[T]{
		wasSetUnordered: false,
		ordered:         false,
		getIter: func() <-chan odata[T] {
			ch := make(chan odata[T])

			go func() {
				defer close(ch)
				i := 0
				for value := range e.getIter() {
					ch <- odata[T]{
						no:  i,
						val: value,
					}
					i++
				}
			}()

			return ch
		},
	}
}

// AsParallelEnumerable creates a new ParallelEnumerable, it will need more resources
// but performance will be improved
//
// # Note that ParallelEnumerable will be useful in some cases, you must really consider when using it.
//
// # In worse cases, not only will performance not be optimized, but costs will also increase.
func AsParallelEnumerable[T any](t []T) ParallelEnumerable[T] {
	return ParallelEnumerable[T]{
		wasSetUnordered: false,
		ordered:         false,
		getIter: func() <-chan odata[T] {
			ch := make(chan odata[T])

			go func() {
				defer close(ch)
				for _, value := range t {
					ch <- odata[T]{
						no:  1,
						val: value,
					}
				}
			}()

			return ch
		},
	}
}

// Concat concatenates two parallel sequences.
func (p ParallelEnumerable[T]) Concat(second ParallelEnumerable[T]) ParallelEnumerable[T] {
	return ParallelEnumerable[T]{
		wasSetUnordered: p.wasSetUnordered,
		ordered:         p.ordered,
		getIter: func() <-chan odata[T] {
			out := make(chan odata[T])

			go func() {
				defer close(out)

				maxNo := -1
				var mu sync.Mutex
				var wg sync.WaitGroup
				for value := range p.getIter() {
					wg.Add(1)
					temp := value
					go func() {
						defer mu.Unlock()
						defer wg.Done()
						mu.Lock()
						if maxNo < temp.no {
							maxNo = temp.no
						}
						out <- temp
					}()
				}
				wg.Wait()

				startNo := maxNo + 1

				var wg1 sync.WaitGroup
				for value := range second.getIter() {
					wg1.Add(1)
					temp := value
					temp.no = temp.no + startNo
					go func() {
						defer wg1.Done()
						out <- temp
					}()
				}
				wg1.Wait()
			}()

			return out
		},
	}
}

// This is helper method. getIterAny return an unbuffered channel of any that iterates through a collection.
func (p ParallelEnumerable[T]) getIterAny() <-chan any {
	ch := make(chan any)

	go func() {
		defer close(ch)
		for value := range p.getIter() {
			ch <- value
		}
	}()

	return ch
}

// This is helper method. order orders by original order in ParallelEnumerable.
func (p ParallelEnumerable[T]) order() ParallelEnumerable[T] {
	return ParallelEnumerable[T]{
		ordered:         true,
		wasSetUnordered: true,
		getIter: func() <-chan odata[T] {
			out := make(chan odata[T])

			go func() {
				defer close(out)
				orderedEnum := AsEnumerableFromChannel(p.getIterAny()).OrderBy(func(o any) any { return o.(odata[T]).no })
				for value := range orderedEnum.getIter() {
					out <- value.(odata[T])
				}
			}()

			return out
		},
	}
}

// AsOrdered enables treatment of a data source as if it were ordered, overriding the default of unordered.
func (p ParallelEnumerable[T]) AsOrdered() ParallelEnumerable[T] {
	reOrdered := func() <-chan odata[T] {
		out := make(chan odata[T])

		go func() {
			defer close(out)
			i := 0
			for value := range p.getIter() {
				value.no = i
				out <- value
				i++
			}
		}()

		return out
	}
	iter := p.getIter
	// in case the current Parallel Enumerable is not ordered, we must re-ordered the origin iterator data
	if !p.ordered {
		iter = reOrdered
	}
	return ParallelEnumerable[T]{
		wasSetUnordered: p.wasSetUnordered,
		ordered:         true,
		getIter:         iter,
	}
}

// AsOrdered aollows an intermediate query to be treated as if no ordering is implied among the elements.
func (p ParallelEnumerable[T]) AsUnordered() ParallelEnumerable[T] {
	return ParallelEnumerable[T]{
		wasSetUnordered: true,
		ordered:         false,
		getIter:         p.getIter,
	}
}
