### Projection operations

#### Select
Select projects values that are based on a transform function.

Example: Take all Ids from students
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "An"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).Select(func(t Student) any {
	return t.Id
})
fmt.Println(enumerable.ToSlice())
// Result: [1 2 3]
```

#### SelectMany
SelectMany projects sequences of values that are based on a transform function and then flattens them into one sequence.

Example: Take all Ids from slice-in-slice students
```go
source := [][]Student{
	{
		{Id: 1, Name: "Anh"},
		{Id: 11, Name: "An"},
	},
	{
		{Id: 2, Name: "Anh"},
		{Id: 21, Name: "An"},
	},
	{
		{Id: 3, Name: "Anh"},
		{Id: 31, Name: "An"},
	},
}

enumerable1 := lingo.AsEnumerable(source).SelectMany(func(ss []Student) []any {
	return lingo.AsEnumerable(ss).Select(func(s Student) any {
		return s.Id
	}).ToSlice()
})
fmt.Println(enumerable1.ToSlice())
// Result: [1 11 2 21 3 31]
```

#### Zip
Zip produces a sequence of tuples with elements from 2 specified sequences.
If resultSelector is nil, the default result is a slice combined with each element

Examples:

Take all Ids from slice-in-slice students
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "An"},
	{Id: 3, Name: "A"},
}
classIds := []any{11, 22, 33}

enumerable := lingo.AsEnumerable(source).Zip(lingo.AsEnumerable(classIds), func(s Student, a any) any {
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
// Result: [{1 Anh 11} {2 An 22} {3 A 33}]
```

In case resultSelector is nil
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "An"},
	{Id: 3, Name: "A"},
}
classIds := []int{11, 22, 33}

enumerable := lingo.AsEnumerable(source).Zip(lingo.AsEnumerableAnyFromSliceT(classIds))
fmt.Println(enumerable.ToSlice())
// Result: [[{1 Anh} 11] [{2 An} 22] [{3 A} 33]]
```
