### Modification enumerable
#### Append
Append appends a value to the end of the sequence.

Example:
```go
enum := lingo.Range(1, 10)
enum.Append(99).ToSlice() // 1-10, 99
```

#### AppendRange
AppendRange appends the elements of the specified collection to the end of the sequence.

Example:
```go
enum := lingo.Range(1, 10)
enum.AppendRange(lingo.Range(100, 110)).ToSlice() // 1-10, 100-110
```

#### Prepend
Prepend adds a value to the beginning of the sequence.

Example:
```go
enum := lingo.Range(1, 10)
enum.Prepend(99).ToSlice() // 99, 1-10
```

#### PrependRange
PrependRange appends the elements of the specified collection to the beginning of the sequence.

Example:
```go
enum := lingo.Range(1, 10)
enum.PrependRange(lingo.Range(100, 110)).ToSlice() // 100-110, 1-10
```

#### Clear
Clear removes all elements of the sequence.

Example:
```go
enum := lingo.Range(1, 10)
enum.Clear().ToSlice() // []
```

#### Insert
Insert inserts an element into the sequence at the specified index.

Example:
```go
enum := lingo.Range(1, 10)
enum.Insert(3, 999).ToSlice() // 1-999, 4-10
```

#### Remove
Remove removes the first occurrence of the given element, if found.

Example:
```go
enum := lingo.Range(1, 10)
enum.Remove(9).ToSlice() // 1-8, 10
```

#### RemoveAt
RemoveAt removes the element at the specified index of the sequence.

Example:
```go
enum := lingo.Range(1, 10)
enum.RemoveAt(6).ToSlice() // 1-6, 8-10
```

#### RemoveRange
RemoveRange removes a range of elements from the sequence.

Example:
```go
enum := lingo.Range(1, 10)
enum.RemoveRange(3, 3).ToSlice() // 1-3, 7-10
```