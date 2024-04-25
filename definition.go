package lingo

// SingleSelector represents a function that converts an object under T type to another object
type SingleSelector[T any] func(T) any

// CombinationSelector represents a function that combines two objects under T, K type to another object
type CombinationSelector[T any, K any] func(T, K) any

// GroupBySelector represents a function that combines two objects under T, []K type to another object
type GroupBySelector[T any, K any] func(T, []K) any

// Comparer represents a function that compares two variables of type T,
// there will be 3 cases: smaller, equal, larger depending on the purpose of use
type Comparer[T any] func(T, T) bool

// GetHashCode represents a function that gets hashcode from an object
type GetHashCode[T any] func(T) any
