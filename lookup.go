package lingo

import (
	"sync"
	"unsafe"

	"github.com/vanh01/lingo/definition"
)

// Lookup represents a collection of keys each mapped to one or more values.
type Lookup[K any, V any] struct {
	Enumerable[Grouping[K, V]]
	item  map[any][]V
	Count int
}

// ContainsKey etermines whether a specified key exists in the Lookup[T, K].
func (l Lookup[K, V]) ContainsKey(t K) bool {
	_, ex := l.item[t]
	return ex
}

// GetValue returns a sequence of values indexed by a specified key.
func (l Lookup[K, V]) GetValue(key K) []V {
	if v, ex := l.item[key]; ex {
		return SliceAnyToT[V](v)
	}
	return nil
}

// Grouping represents a collection of objects that have a common key.
type Grouping[K any, V any] struct {
	Enumerable[V]
	Key   K
	Count int
}

// AsLookup Creates a generic Lookup[K, V] from an Enumerable[T].
//
// elementSelector can be nil. If elementSelector is not empty or nil, we will use the first elementSelector
func AsLookup[T any, K any, V any](
	e Enumerable[T],
	keySelector definition.SingleSelectorFull[T, K],
	elementSelector ...definition.SingleSelectorFull[T, V],
) Lookup[K, V] {
	// group by section
	res := map[any][]V{}
	for value := range e.getIter() {
		var element V
		key := keySelector(value)
		if !definition.IsEmptyOrNil(elementSelector) {
			element = elementSelector[0](value)
		} else {
			temp, ok := any(value).(V)
			if !ok {
				element = *(*V)(unsafe.Pointer(&value))
			} else {
				element = temp
			}
		}
		res[key] = append(res[key], element)
	}

	return Lookup[K, V]{
		Count: len(res),
		item:  res,
		Enumerable: Enumerable[Grouping[K, V]]{
			getIter: func() <-chan Grouping[K, V] {
				out := make(chan Grouping[K, V])

				go func() {
					defer close(out)
					// Initialize LookUp
					for k := range res {
						// must declare here to load data into memmory
						temp := res[k]
						out <- Grouping[K, V]{
							Key:        k.(K),
							Count:      len(temp),
							Enumerable: AsEnumerable(temp),
						}
					}
				}()

				return out
			},
		},
	}
}

// ParallelEnumerable

// AsPLookup creates a generic Lookup[K, V] from an ParallelEnumerable[T].
//
// elementSelector can be nil. If elementSelector is not empty or nil, we will use the first elementSelector
func AsPLookup[T any, K any, V any](
	p ParallelEnumerable[T],
	keySelector definition.SingleSelectorFull[T, K],
	elementSelector ...definition.SingleSelectorFull[T, V],
) Lookup[K, V] {
	// group by section
	mapdata := make(chan definition.KeyValData[any, V])

	go func() {
		defer close(mapdata)
		var wg sync.WaitGroup
		for value := range p.getIter() {
			wg.Add(1)
			tempValue := value
			go func() {
				defer wg.Done()

				ele := make(chan V)
				go func() {
					defer close(ele)
					var element V
					if !definition.IsEmptyOrNil(elementSelector) {
						element = elementSelector[0](tempValue.val)
					} else {
						temp, ok := any(tempValue.val).(V)
						if !ok {
							element = *(*V)(unsafe.Pointer(&tempValue.val))
						} else {
							element = temp
						}
					}
					ele <- element
				}()

				key := keySelector(tempValue.val)

				mapdata <- definition.KeyValData[any, V]{
					Key: key,
					Val: <-ele,
				}
			}()
		}
		wg.Wait()
	}()

	res := map[any][]V{}
	for d := range mapdata {
		res[d.Key] = append(res[d.Key], d.Val)
	}

	return Lookup[K, V]{
		Count: len(res),
		item:  res,
		Enumerable: Enumerable[Grouping[K, V]]{
			getIter: func() <-chan Grouping[K, V] {
				out := make(chan Grouping[K, V])

				go func() {
					defer close(out)
					// Initialize LookUp
					for k := range res {
						// must declare here to load data into memmory
						temp := res[k]
						out <- Grouping[K, V]{
							Key:        k.(K),
							Count:      len(temp),
							Enumerable: AsEnumerable(temp),
						}
					}
				}()

				return out
			},
		},
	}
}
