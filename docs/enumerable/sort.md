### Sorting data
#### OrderBy
OrderBy sorts values in ascending order.

Example:

Sort Student by Student Id
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 5, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).OrderBy(func(s Student) any {
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

enumerable := lingo.AsEnumerable(source).OrderBy(func(s Student) any {
	return s.Id
}, func(a1, a2 any) bool {
	return a1.(int) >= a2.(int)
})
fmt.Println(enumerable.ToSlice())
// Result: [{5 hnA} {3 A} {1 Anh}]
```

#### OrderByDescending
OrderByDescending sorts values in descending order.

Example:

```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 5, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).OrderByDescending(func(s Student) any {
	return s.Id
})
fmt.Println(enumerable.ToSlice())
// Result: [{5 hnA} {3 A} {1 Anh}]
```

#### Reverse
Reverse reverses the order of the elements in a collection.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 5, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).Reverse()
fmt.Println(enumerable.ToSlice())
// Result: [[{3 A} {5 hnA} {1 Anh}]]
```