### Retrieve data
#### FirstOrNil
FirstOrNil returns the first element of a sequence (with condition if any), or a nil value if no element is found

Example:
```go
enumerable := lingo.Range(1, 10)
first := enumerable.FirstOrNil() // 1
```
#### FirstOrDefault
FirstOrDefault returns the first element of a sequence (with condition if any), or a default value if no element is found

Example:
```go
enumerable := lingo.Empty[int]()
first := enumerable.FirstOrDefault(-999) // -999
```
#### LastOrNil
LastOrNil returns the last element of a sequence (with condition if any), or a nil value if no element is found

Example:
```go
enumerable := lingo.Empty[int]()
last := enumerable.LastOrNil() // 0
```
#### LastOrDefault
LastOrDefault returns the last element of a sequence (with condition if any), or a default value if no element is found

Example:
```go
enumerable := lingo.Empty[int]()
last := enumerable.LastOrDefault(999) // 999
```
#### ElementAtOrNil
ElementAtOrNil returns the element at a specified index in a sequence or a default value if the index is out of range.

Example:
```go
enumerable := lingo.Range(1, 100)
element := enumerable.ElementAtOrNil(54) // 55
```
#### ElementAtOrDefault
ElementAtOrDefault returns the element at a specified index in a sequence or a default value if the index is out of range.

Example:
```go
enumerable := lingo.Range(1, 100)
element := enumerable.ElementAtOrDefault(100, -1) // -1
```