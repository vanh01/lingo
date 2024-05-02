package lingo

// Skip skips elements up to a specified position in a sequence.
func (e Enumerable[T]) Skip(number int) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)
			go func() {
				defer close(out)
				i := 1
				for value := range e.getIter() {
					if i > number {
						out <- value
					}
					i++
				}
			}()
			return out
		},
	}
}

// SkipWhile skips elements based on a predicate function until an element doesn't satisfy the condition.
func (e Enumerable[T]) SkipWhile(predicate Predicate[T]) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				stopped := false
				for value := range e.getIter() {
					if !predicate(value) {
						stopped = true
					}
					if stopped {
						out <- value
					}
				}
			}()

			return out
		},
	}
}

// Take takes elements up to a specified position in a sequence.
func (e Enumerable[T]) Take(number int) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				i := 0
				for value := range e.getIter() {
					if i < number {
						out <- value
					}
					i++
				}
			}()

			return out
		},
	}
}

// TakeWhile takes elements based on a predicate function until an element doesn't satisfy the condition.
func (e Enumerable[T]) TakeWhile(predicate Predicate[T]) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				stopped := false
				for value := range e.getIter() {
					if !predicate(value) {
						stopped = true
					}
					if !stopped {
						out <- value
					}
				}
			}()

			return out
		},
	}
}
