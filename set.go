package lingo

// Distinct removes duplicate values from a collection.
func (e Enumerable[T]) Distinct() Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				m := map[any]struct{}{}
				for value := range e.getIter() {
					if _, ex := m[value]; !ex {
						m[value] = struct{}{}
						out <- value
					}
				}
			}()

			return out
		},
	}
}

// DistinctBy removes duplicate values from a collection with keySelector and comparer,
//
// in this method, comparer is returns whether left is equal to right or not.
//
// comparer can be nil without special comparer.
func (e Enumerable[T]) DistinctBy(keySelector SingleSelector[T], comparer Comparer[any]) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				if comparer == nil {
					m := map[any]struct{}{}
					for value := range e.getIter() {
						key := keySelector(value)
						if _, ex := m[key]; !ex {
							m[key] = struct{}{}
							out <- value
						}
					}
				} else {
					temp := []any{}
					for value := range e.getIter() {
						key := keySelector(value)
						exist := false
						for i := range temp {
							if comparer(key, temp[i]) {
								exist = true
								break
							}
						}
						if !exist {
							temp = append(temp, value)
							out <- value
						}
					}
				}
			}()

			return out
		},
	}
}

// Except returns the set difference, which means the elements of one collection
// that don't appear in a second collection.
func (e Enumerable[T]) Except(second Enumerable[T]) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				secondSlice := second.ToSlice()
				secondMap := map[any]int{}
				for _, e := range secondSlice {
					if v, ex := secondMap[e]; ex {
						secondMap[e] = v + 1
					} else {
						secondMap[e] = 1
					}
				}
				for value := range e.getIter() {
					v, ex := secondMap[value]
					if !ex || v == 0 {
						out <- value
					} else {
						secondMap[value] = v - 1
					}
				}
			}()

			return out
		},
	}
}

// Intersect returns the set intersection, which means elements
// that appear in each of two collections.
func (e Enumerable[T]) Intersect(second Enumerable[T]) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				secondSlice := second.ToSlice()
				secondMap := map[any]int{}
				for _, e := range secondSlice {
					if v, ex := secondMap[e]; ex {
						secondMap[e] = v + 1
					} else {
						secondMap[e] = 1
					}
				}
				for value := range e.getIter() {
					if v, ex := secondMap[value]; ex && v > 0 {
						secondMap[value] = v - 1
						out <- value
					}
				}
			}()

			return out
		},
	}
}

// Union returns the set union, which means unique elements that
// appear in either of two collections.
func (e Enumerable[T]) Union(second Enumerable[T]) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				secondSlice := second.ToSlice()
				secondMap := map[any]int{}
				for _, e := range secondSlice {
					if v, ex := secondMap[e]; ex {
						secondMap[e] = v + 1
					} else {
						secondMap[e] = 1
					}
				}
				for value := range e.getIter() {
					if v, ex := secondMap[value]; ex {
						secondMap[value] = v - 1
					}
					out <- value
				}
				for k, v := range secondMap {
					if v > 0 {
						out <- k.(T)
					}
				}
			}()

			return out
		},
	}
}
