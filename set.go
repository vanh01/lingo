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
							temp = append(temp, key)
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
				secondMap := second.ToMap(func(t T) any { return t }, func(t T) any {
					return struct{}{}
				})
				for value := range e.getIter() {
					_, ex := secondMap[value]
					if !ex {
						out <- value
						secondMap[value] = struct{}{}
					}
				}
			}()

			return out
		},
	}
}

// ExceptBy returns the set difference, which means the elements of one collection
// that don't appear in a second collection according to a specified key selector function.
//
// In this method, comparer is returns whether left is equal to right or not, comparer can be nil without special comparer.
func (e Enumerable[T]) ExceptBy(second Enumerable[any], keySelector SingleSelector[T], comparer Comparer[any]) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				if comparer == nil {
					secondMap := second.ToMap(func(a any) any { return a }, func(a any) any {
						return struct{}{}
					})
					for value := range e.getIter() {
						key := keySelector(value)
						_, ex := secondMap[key]
						if !ex {
							out <- value
							secondMap[key] = struct{}{}
						}
					}
				} else {
					setKey := []any{}
					secondKey := second.ToSlice()
					for value := range e.getIter() {
						fKey := keySelector(value)
						exist := false
						for i := range setKey {
							if comparer(fKey, setKey[i]) {
								exist = true
								break
							}
						}
						if exist {
							continue
						}
						for i := range secondKey {
							if comparer(fKey, secondKey[i]) {
								exist = true
								break
							}
						}
						if !exist {
							out <- value
							setKey = append(setKey, fKey)
						}
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
				secondMap := map[any]bool{}
				for value := range second.getIter() {
					secondMap[value] = true
				}
				for value := range e.getIter() {
					if v, ex := secondMap[value]; ex && v {
						secondMap[value] = false
						out <- value
					}
				}
			}()

			return out
		},
	}
}

// IntersectBy returns the set intersection, which means elements
// that appear in each of two collections according to a specified key selector function.
//
// In this method, comparer is returns whether left is equal to right or not, comparer can be nil without special comparer.
func (e Enumerable[T]) IntersectBy(second Enumerable[any], keySelector SingleSelector[T], comparer Comparer[any]) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				if comparer == nil {
					secondMap := second.ToMap(func(a any) any { return a }, func(t any) any {
						return true
					})
					for value := range e.getIter() {
						key := keySelector(value)
						if v, ex := secondMap[key]; ex && v.(bool) {
							secondMap[key] = false
							out <- value
						}
					}
				} else {
					setKey := []any{}
					// because there is a loop for the second in the first, avoid duplicate logic for keySelector
					secondKey := second.ToSlice()
					for value := range e.getIter() {
						fKey := keySelector(value)
						exist := false
						for i := range setKey {
							if comparer(fKey, setKey[i]) {
								exist = true
								break
							}
						}
						if exist {
							continue
						}
						for i := range secondKey {
							if comparer(fKey, secondKey[i]) {
								exist = true
								break
							}
						}
						if exist {
							out <- value
							setKey = append(setKey, fKey)
						}
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
				m := map[any]struct{}{}
				for value := range e.getIter() {
					if _, ex := m[value]; !ex {
						out <- value
						m[value] = struct{}{}
					}
				}
				for value := range second.getIter() {
					if _, ex := m[value]; !ex {
						out <- value
						m[value] = struct{}{}
					}
				}
			}()

			return out
		},
	}
}

// UnionBy returns the set union, which means unique elements that
// appear in either of two collections according to a specified key selector function.
//
// In this method, comparer is returns whether left is equal to right or not, comparer can be nil without special comparer.
func (e Enumerable[T]) UnionBy(second Enumerable[T], keySelector SingleSelector[T], comparer Comparer[any]) Enumerable[T] {
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
							out <- value
							m[key] = struct{}{}
						}
					}
					for value := range second.getIter() {
						key := keySelector(value)
						if _, ex := m[key]; !ex {
							out <- value
							m[key] = struct{}{}
						}
					}
				} else {
					setKey := []any{}
					for value := range e.getIter() {
						key := keySelector(value)
						exist := false
						for i := range setKey {
							if comparer(key, setKey[i]) {
								exist = true
								break
							}
						}
						if exist {
							continue
						}
						out <- value
						setKey = append(setKey, key)
					}
					for value := range second.getIter() {
						key := keySelector(value)
						exist := false
						for i := range setKey {
							if comparer(key, setKey[i]) {
								exist = true
								break
							}
						}
						if exist {
							continue
						}
						out <- value
						setKey = append(setKey, key)
					}
				}
			}()

			return out
		},
	}
}
