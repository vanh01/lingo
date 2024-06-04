### Quantifier operations
#### All
All determines whether all the elements in a sequence satisfy a condition.

Example:

```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).All(func(s Student) bool {
	return len(s.Name) > 0
})
fmt.Println(enumerable)
// Result: true
```
#### Any
Any determines whether any elements in a sequence satisfy a condition.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).Any(func(s Student) bool {
	return len(s.Name) > 5
})
fmt.Println(enumerable)
// Result: false
```

#### Contains
Contains determines whether a sequence contains a specified element.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).Contains(Student{Id: 5, Name: "A"}, func(s1, s2 Student) bool {
	return s1.Name == s1.Name
})
fmt.Println(enumerable)
// Result: true
```