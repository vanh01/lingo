### Partitioning data
#### Skip
Skip skips a specified number of elements in a parallel sequence and then returns the remaining elements.
If the source sequence is ordered, Skip skips first n elements. On the other hand, Skip skips any n elements.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsParallelEnumerable(source).Skip(2)
fmt.Println(enumerable.ToSlice())
// Result: [{3 A}] (in somecase)
```
#### SkipWhile
SkipWhile elements in a parallel sequence as long as a specified condition is true and then returns the remaining elements.
If the source sequence is ordered, SkipWhile skips according to the ordered element. On the other hand, performs SkipWhile on the current arbitrary order.

Example:

```go
source := []Student{
    {Id: 1, Name: "Anh"},
    {Id: 2, Name: "hnA"},
    {Id: 3, Name: "A"},
}

enumerable := lingo.AsParallelEnumerable(source).SkipWhile(func(s Student) bool {
    return s.Id < 2
})
fmt.Println(enumerable.ToSlice())
// Result: [{2 hnA} {1 Anh} {3 A}] (in somecase)
```
#### Take
Take returns a specified number of contiguous elements from the start of a parallel sequence.
If the source sequence is ordered, Take takes first n elements. On the other hand, Take takes any n elements.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsParallelEnumerable(source).Take(2)
fmt.Println(enumerable.ToSlice())
// Result: [{2 hnA} {1 Anh}] (in somecase)
```
#### TakeWhile
TakeWhile takes elements in a parallel sequence based on a predicate function until an element doesn't satisfy the condition.
If the source sequence is ordered, TakeWhile takes according to the ordered element. On the other hand, performs TakeWhile on the current arbitrary order.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsParallelEnumerable(source).TakeWhile(func(s Student) bool {
	return s.Id < 3
})
fmt.Println(enumerable.ToSlice())
// Result: [{2 hnA} {1 Anh}] (in somecase)
```