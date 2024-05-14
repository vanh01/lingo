### Set operations
#### Distinct
Distinct removes duplicate values from a collection.

Example:

Get all unique students
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "An"},
	{Id: 3, Name: "A"},
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "An"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).Distinct()
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 An} {3 A}]
```

#### DistinctBy
DistinctBy removes duplicate values from a collection with keySelector and comparer, comparer can be nil without special comparer.

Examples:

Get all unique students by length of name
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
	{Id: 1, Name: "AKh"},
	{Id: 2, Name: "Lah"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).DistinctBy(func(s Student) any {
	return len(s.Name)
})
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {3 A}]
```

#### Except
Except returns the set difference, which means the elements of one collection that don't appear in a second collection.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
	{Id: 1, Name: "AKh"},
	{Id: 2, Name: "Lah"},
	{Id: 3, Name: "A"},
}

second := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
}

enumerable := lingo.AsEnumerable(source).Except(lingo.AsEnumerable(second))
fmt.Println(enumerable.ToSlice())
// Result: [{3 A} {1 AKh} {2 Lah}]
```

#### ExceptBy
ExceptBy returns the set difference, which means the elements of one collection that don't appear in a second collection according to a specified key selector function.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

second := []any{
	"Anh",
	"hnA",
}

enumerable := lingo.AsEnumerable(source).ExceptBy(lingo.AsEnumerable(second), func(s Student) any {
	return s.Name
})
fmt.Println(enumerable.ToSlice())
// Result: [{3 A}]
```

#### Intersect
Intersect returns the set intersection, which means elements that appear in each of two collections.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
	{Id: 1, Name: "AKh"},
	{Id: 2, Name: "Lah"},
	{Id: 3, Name: "A"},
}

second := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "Ab"},
	{Id: 1, Name: "Ah"},
	{Id: 2, Name: "Lbh"},
	{Id: 3, Name: "B"},
}

enumerable := lingo.AsEnumerable(source).Intersect(lingo.AsEnumerable(second))
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 hnA}]
```

#### IntersectBy
Intersect returns the set intersection, which means elements that appear in each of two collections according to a specified key selector function.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

second := []any{
	"Anh",
	"hnA",
}

enumerable := lingo.AsEnumerable(source).IntersectBy(lingo.AsEnumerable(second), func(s Student) any {
	return s.Name
})
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 hnA}]
```

#### Union
Union returns the set union, which means unique elements that appear in either of two collections.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

second := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "Ab"},
}

enumerable := lingo.AsEnumerable(source).Union(lingo.AsEnumerable(second))
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 hnA} {3 A} {3 Ab}]
```

#### UnionBy
Union returns the set union, which means unique elements that appear in either of two collections according to a specified key selector function.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

second := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
}

enumerable := lingo.AsEnumerable(source).UnionBy(lingo.AsEnumerable(second), func(s Student) any {
	return s.Name
})
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 hnA} {3 A}]
```