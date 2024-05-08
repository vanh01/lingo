package lingo

import "unsafe"

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
// elementSelector can be nil.
func AsLookup[T any, K any, V any](
	e Enumerable[T],
	keySelector SingleSelectorFull[T, K],
	elementSelector SingleSelectorFull[T, V],
) Lookup[K, V] {
	// group by section
	res := map[any][]V{}
	for value := range e.getIter() {
		var element V
		key := keySelector(value)
		if elementSelector != nil {
			element = elementSelector(value)
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
							Key:   k.(K),
							Count: len(temp),
							Enumerable: Enumerable[V]{
								getIter: func() <-chan V {
									outGrouping := make(chan V)

									go func() {
										defer close(outGrouping)
										for i := range temp {
											outGrouping <- temp[i]
										}
									}()

									return outGrouping
								},
							},
						}
					}
				}()

				return out
			},
		},
	}
}
