### Projection operations
#### Select
Select projects in parallel each element of a sequence into a new form.

Example:
```go
penum := lingo.Range(1, 100).AsParallel().Where(func(i int) bool {
	time.Sleep(time.Second)
	return i > 98
}).Select(func(i int) any {
	time.Sleep(time.Second)
	return i * i
}).ToSlice() // [10000 9801]
```
#### SelectMany
SelectMany projects in parallel each element of a sequence to an []T and flattens the resulting sequences into one sequence.

Example:
```go
penum := lingo.Range(1, 10).AsParallel().SelectMany(func(i int) any {
	time.Sleep(time.Second)
	return lingo.Repeat(i, 5).ToSlice()
})
```
#### Zip
Zip merges in parallel two sequences by using the specified predicate function.
If resultSelector is nil, the default result is a slice combined with each element. On the other hand, we just use the first resultSelector

Examples:

Take all Ids from slice-in-slice students
```go
source := []Student{
    {Id: 1, Name: "Anh"},
    {Id: 2, Name: "An"},
    {Id: 3, Name: "A"},
}
classIds := []any{11, 22, 33}

enumerable := lingo.AsParallelEnumerable(source).Zip(lingo.AsParallelEnumerable(classIds), func(s Student, a any) any {
    return struct {
        Id      int
        Name    string
        ClassId int
    }{
        Id:      s.Id,
        Name:    s.Name,
        ClassId: a.(int),
    }
})
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh 11} {2 An 22} {3 A 33}] (any order)
```

In case resultSelector is nil
```go
source := []Student{
    {Id: 1, Name: "Anh"},
    {Id: 2, Name: "An"},
    {Id: 3, Name: "A"},
}
classIds := []int{11, 22, 33}

enumerable := lingo.AsParallelEnumerable(source).Zip(lingo.AsEnumerableAnyFromSliceT(classIds).AsParallel())
fmt.Println(enumerable.ToSlice())
// Result: [[{1 Anh} 11] [{2 An} 22] [{3 A} 33]] (any order)
```