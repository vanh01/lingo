### Partitioning data
#### Skip
Skip skips elements up to a specified position in a sequence.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).Skip(2)
fmt.Println(enumerable.ToSlice())
// Result: [{3 A}]
```
#### SkipWhile
SkipWhile skips elements based on a predicate function until an element doesn't satisfy the condition.

Example:

```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).SkipWhile(func(s Student) bool {
	return s.Id < 3
})
fmt.Println(enumerable.ToSlice())
// Result: [{3 A}]
```
#### Take
Take takes elements up to a specified position in a sequence.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).Take(2)
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 hnA}]
```
#### TakeWhile
TakeWhile takes elements based on a predicate function until an element doesn't satisfy the condition.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).TakeWhile(func(s Student) bool {
	return s.Id < 3
})
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 hnA}]
```

#### Chunk
Chunk splits the elements of a sequence into chunks of a specified maximum size.

Example:
```go
lingo.Chunk(lingo.Range(1, 11), 5).ToSlice()
// 1-5
// 6-10
// 11
```