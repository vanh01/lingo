### Retrieve data
#### FirstOrNil
FirstOrNil returns the first element of a parallel sequence (with condition if any), or a nil value if no element is found

Example:
```go
penumerable := lingo.Range(1, 10).AsParallel()
first := penumerable.FirstOrNil() // any element
```
#### FirstOrDefault
FirstOrDefault returns the first element of a parallel sequence (with condition if any), or a default value if no element is found

Example:
```go
penumerable := lingo.Empty[int]().AsParallel()
first := penumerable.FirstOrDefault(-999) // -999
```
#### LastOrNil
LastOrNil returns the last element of a parallel sequence (with condition if any), or a nil value if no element is found

Example:
```go
penumerable := lingo.Empty[int]().AsParallel()
last := enumerabpenumerablele.LastOrNil() // 0
```
#### LastOrDefault
LastOrDefault returns the last element of a parallel sequence (with condition if any), or a default value if no element is found

Example:
```go
penumerable := lingo.Empty[int]().AsParallel()
last := penumerable.LastOrDefault(999) // 999
```
#### ElementAtOrNil
ElementAtOrNil returns the element at a specified index in a parallel sequence or a default value if the index is out of range.

Example:
```go
enumerable := lingo.Range(1, 100).AsParallel()
element := enumerable.ElementAtOrNil(54) // any element
```
#### ElementAtOrDefault
ElementAtOrDefault returns the element at a specified index in a parallel sequence or a default value if the index is out of range.

Example:
```go
enumerable := lingo.Range(1, 100).AsParallel()
element := enumerable.ElementAtOrDefault(100, -1) // -1
```