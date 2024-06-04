### Converting data types
#### ToSlice
ToSlice converts the iterator into a slice

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source)
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 hnA} {3 A}]
```

#### ToMap
ToMap converts the iterator into a map with specific selector

```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

m := lingo.AsEnumerable(source).ToMap(func(s Student) any {
	return s.Id
}, func(s Student) any {
	return s.Name
})
fmt.Println(m)
// Result: map[1:Anh 2:hnA 3:A]
```