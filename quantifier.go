package lingo

import (
	"reflect"
)

// All determines whether all the elements in a sequence satisfy a condition.
func (e Enumerable[T]) All(predicate Predicate[T]) bool {
	for value := range e.getIter() {
		if !predicate(value) {
			return false
		}
	}
	return true
}

// Any determines whether any elements in a sequence satisfy a condition.
func (e Enumerable[T]) Any(predicate Predicate[T]) bool {
	for value := range e.getIter() {
		if predicate(value) {
			return true
		}
	}
	return false
}

// Contains determines whether a sequence contains a specified element.
//
// In this method, comparer is returns whether left is equal to right or not.
//
// If comparer is empty or nil, we will use the default comparer.
// On the other hand, we just use the first comparer
func (e Enumerable[T]) Contains(value T, comparer ...Comparer[T]) bool {
	for v := range e.getIter() {
		if isEmptyOrNil(comparer) {
			if reflect.ValueOf(v).Interface() == reflect.ValueOf(value).Interface() {
				return true
			}
		} else {
			if comparer[0](v, value) {
				return true
			}
		}
	}
	return false
}
