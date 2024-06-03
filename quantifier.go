package lingo

import (
	"reflect"
	"sync"

	"github.com/vanh01/lingo/definition"
)

// All determines whether all the elements in a sequence satisfy a condition.
func (e Enumerable[T]) All(predicate Predicate[T]) bool {
	result := true
	for value := range e.getIter() {
		if !result {
			continue
		}
		if !predicate(value) {
			result = false
		}
	}
	return result
}

// Any determines whether any elements in a sequence satisfy a condition.
func (e Enumerable[T]) Any(predicate Predicate[T]) bool {
	result := false
	for value := range e.getIter() {
		if result {
			continue
		}
		if predicate(value) {
			result = true
		}
	}
	return result
}

// Contains determines whether a sequence contains a specified element.
//
// In this method, comparer is returns whether left is equal to right or not.
//
// If comparer is empty or nil, we will use the default comparer.
// On the other hand, we just use the first comparer
func (e Enumerable[T]) Contains(value T, comparer ...definition.Comparer[T]) bool {
	result := false
	for v := range e.getIter() {
		if result {
			continue
		}
		if definition.IsEmptyOrNil(comparer) {
			if reflect.ValueOf(v).Interface() == reflect.ValueOf(value).Interface() {
				result = true
			}
		} else {
			if comparer[0](v, value) {
				result = true
			}
		}
	}
	return result
}

// ParallelEnumerable

// All determines in parallel whether all elements of a sequence satisfy a condition.
func (p ParallelEnumerable[T]) All(predicate Predicate[T]) bool {
	res := make(chan bool)

	go func() {
		defer close(res)

		var wg sync.WaitGroup
		for value := range p.getIter() {
			wg.Add(1)
			temp := value
			go func() {
				defer wg.Done()
				res <- predicate(temp.val)
			}()
		}
		wg.Wait()
	}()

	result := true
	for value := range res {
		if !result {
			continue
		}
		if !value {
			result = false
		}
	}
	return result
}

// Any determines in parallel whether any element of a sequence satisfies a condition.
func (p ParallelEnumerable[T]) Any(predicate Predicate[T]) bool {
	res := make(chan bool)

	go func() {
		defer close(res)

		var wg sync.WaitGroup
		for value := range p.getIter() {
			wg.Add(1)
			temp := value
			go func() {
				defer wg.Done()
				res <- predicate(temp.val)
			}()
		}
		wg.Wait()
	}()

	result := false
	for value := range res {
		if result {
			continue
		}
		if value {
			result = true
		}
	}
	return result
}

// Contains determines in parallel whether a sequence contains a specified element.
//
// In this method, comparer is returns whether left is equal to right or not.
//
// If comparer is empty or nil, we will use the default comparer.
// On the other hand, we just use the first comparer
func (p ParallelEnumerable[T]) Contains(value T, comparer ...definition.Comparer[T]) bool {
	res := make(chan bool)

	go func() {
		defer close(res)

		var wg sync.WaitGroup
		for v := range p.getIter() {
			wg.Add(1)
			temp := v
			go func() {
				defer wg.Done()
				if definition.IsEmptyOrNil(comparer) {
					if reflect.ValueOf(temp.val).Interface() == reflect.ValueOf(value).Interface() {
						res <- true
					}
				} else {
					if comparer[0](temp.val, value) {
						res <- true
					}
				}
			}()
		}
		wg.Wait()
	}()

	result := false
	for value := range res {
		if result {
			continue
		}
		if value {
			result = true
		}
	}
	return result
}
