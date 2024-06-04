### Quantifier operations
#### All
All determines in parallel whether all elements of a sequence satisfy a condition.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).AsParallel().All(func(s Student) bool {
	return len(s.Name) > 0
})
fmt.Println(enumerable)
// Result: true
```
#### Any
Any determines in parallel whether any element of a sequence satisfies a condition.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).AsParallel().Any(func(s Student) bool {
	return len(s.Name) > 5
})
fmt.Println(enumerable)
// Result: false
```
#### Contains
Contains determines in parallel whether a sequence contains a specified element.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).AsParallel().Contains(Student{Id: 5, Name: "A"}, func(s1, s2 Student) bool {
	return s1.Name == s1.Name
})
fmt.Println(enumerable)
// Result: true
```