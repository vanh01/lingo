### Filtering data

I want to take all students who have the student name contains "an"

```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "An"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).
	Where(func(t Student) bool {
		return strings.Contains(strings.ToLower(t.Name), "an")
	})
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 An}]
```