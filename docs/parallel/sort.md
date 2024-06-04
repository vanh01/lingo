### Sorting data
#### OrderBy
OrderBy sorts in parallel values in ascending order. And then set ordered is true

Example:

Sort Student by Student Id
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 5, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsParallelEnumerable(source).OrderBy(func(s Student) any {
	return s.Id
})
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {3 A} {5 hnA}]
```

Sort Student by Student Id with special comparer
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 5, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsParallelEnumerable(source).OrderBy(func(s Student) any {
	return s.Id
}, func(a1, a2 any) bool {
	return a1.(int) >= a2.(int)
})
fmt.Println(enumerable.ToSlice())
// Result: [{5 hnA} {3 A} {1 Anh}]
```

#### OrderByDescending
OrderByDescending sorts in parallel values in descending order. And then set ordered is true

Example:

```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 5, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsParallelEnumerable(source).OrderByDescending(func(s Student) any {
	return s.Id
})
fmt.Println(enumerable.ToSlice())
// Result: [{5 hnA} {3 A} {1 Anh}]
```

#### Reverse
Reverse inverts the order of the elements in a parallel sequence. If this one is Unordered, does nothing.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 5, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsParallelEnumerable(source).AsOrdered().Reverse()
fmt.Println(enumerable.ToSlice())
// Result: [[{3 A} {5 hnA} {1 Anh}]]
```