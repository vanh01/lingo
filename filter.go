package lingo

// Where selects values that are based on a predicate function.
func (e Enumerable[T]) Where(predicate Predicate[T]) Enumerable[T] {
	output := make(chan T)

	go func() {
		defer close(output)
		for value := range e.iterator {
			if predicate(value) {
				output <- value
			}
		}
	}()

	return Enumerable[T]{
		iterator: output,
	}
}
