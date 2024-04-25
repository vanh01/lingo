package lingo

type Predicate[T any] func(T) bool

// And returns a new predicate, it is current && right
func (p Predicate[T]) And(right Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return p(t) && right(t)
	}
}

// Or returns a new predicate, it is current || right
func (p Predicate[T]) Or(right Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return p(t) || right(t)
	}
}
