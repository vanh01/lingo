package lingo

// FirstOrNil returns the first element of a sequence (with condition if any),
// or a nil value if no element is found
//
// predicate can be nil. If predicate is not empty or nil, we will use the first predicate
func (e Enumerable[T]) FirstOrNil(predicate ...Predicate[T]) T {
	var t T
	first := true
	for value := range e.getIter() {
		if !isEmptyOrNil(predicate) {
			if !predicate[0](value) {
				continue
			}
		}
		if first {
			first = false
			t = value
		}
	}
	return t
}

// FirstOrDefault returns the first element of a sequence (with condition if any),
// or a default value if no element is found
//
// predicate can be nil. If predicate is not empty or nil, we will use the first predicate
func (e Enumerable[T]) FirstOrDefault(defaultValue T, predicate ...Predicate[T]) T {
	var t T = defaultValue
	first := true
	for value := range e.getIter() {
		if !isEmptyOrNil(predicate) {
			if !predicate[0](value) {
				continue
			}
		}
		if first {
			first = false
			t = value
		}
	}
	return t
}

// LastOrNil returns the last element of a sequence (with condition if any),
// or a nil value if no element is found
//
// predicate can be nil. If predicate is not empty or nil, we will use the first predicate
func (e Enumerable[T]) LastOrNil(predicate ...Predicate[T]) T {
	var t T
	for value := range e.getIter() {
		if !isEmptyOrNil(predicate) {
			if !predicate[0](value) {
				continue
			}
		}
		t = value
	}
	return t
}

// LastOrDefault returns the last element of a sequence (with condition if any),
// or a default value if no element is found
//
// predicate can be nil. If predicate is not empty or nil, we will use the first predicate
func (e Enumerable[T]) LastOrDefault(defaultValue T, predicate ...Predicate[T]) T {
	var t T = defaultValue
	for value := range e.getIter() {
		if !isEmptyOrNil(predicate) {
			if !predicate[0](value) {
				continue
			}
		}
		t = value
	}
	return t
}

// ElementAtOrNil returns the element at a specified index in a sequence or a default value if the index is out of range.
func (e Enumerable[T]) ElementAtOrNil(index int64) T {
	var t T
	var i int64 = 0
	for value := range e.getIter() {
		if i <= index {
			t = value
			i++
		}
	}
	if i != index+1 {
		var tt T
		return tt
	}
	return t
}

// ElementAtOrDefault returns the element at a specified index in a sequence or a default value if the index is out of range.
func (e Enumerable[T]) ElementAtOrDefault(index int64, defaultValue T) T {
	var t T = defaultValue
	var i int64 = 0
	for value := range e.getIter() {
		if i <= index {
			t = value
			i++
		}
	}
	if i != index+1 {
		return defaultValue
	}
	return t
}
