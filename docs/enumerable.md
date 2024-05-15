### Initialize enumerable
#### AsEnumerable
AsEnumerable creates a new Enumerable from a silce of specific type.

Example:
```go
enumerable := lingo.AsEnumerable([]int{1, 2, 3})
```
#### AsEnumerableFromChannel
AsEnumerableFromChannel creates a new Enumerable from a receive-only channel

Example:
```go
enumerable := lingo.AsEnumerableFromChannel(lingo.AsEnumerable([]int{1, 2, 3}).GetIter())
```
#### AsEnumerableTFromAny
AsEnumerableTFromAny creates a new Enumerable of specific type from Enumerable of any type, this method will be useful after using projection operations.

Example:
```go
enumerable := lingo.AsEnumerable([]any{1, 2, 3})
enumerableInt := lingo.AsEnumerableTFromAny[int](enumerable)
```
#### AsEnumerableAnyFromT
AsEnumerableAnyFromT creates a new Enumerable of any type from Enumerable of specific type.

Example:
```go
enumerableInt := lingo.AsEnumerable([]int{1, 2, 3})
enumerableAny := lingo.AsEnumerableAnyFromT(enumerableInt)
```
#### AsEnumerableTFromSliceAny
AsEnumerableTFromSliceAny creates a new Enumerable of specific type from slice of any type

Example:
```go
enumerableInt := lingo.AsEnumerableTFromSliceAny[int]([]any{1, 2, 3})
```
#### AsEnumerableAnyFromSliceT
AsEnumerableAnyFromSliceT creates a new Enumerable of any type from slice of specific type

Example:
```go
enumerableAny := lingo.AsEnumerableAnyFromSliceT([]int{1, 2, 3})
```
#### Empty
Empty returns an empty Enumerable[T] that has the specified type argument.

Example:
```go
emptyInt := lingo.Empty[int]()
```
#### Range
Range generates a sequence of integral numbers within a specified range.

Example: create an enumerable of int from 1 to 10
```go
rangeInt := lingo.Range(1, 10)
```
#### Repeat
Repeat generates a sequence that contains one repeated value.

Example: create a enumerable of int, it contains 10 elements 1
```go
repeatInt := lingo.Repeat(1, 10)
```
#### Concat
Concat concatenates two sequences.

Example:
```go
first := lingo.Range(1, 10)
second := lingo.Range(11, 20)
first.Concat(second) // 1-20
```