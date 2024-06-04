### Set operations
#### Distinct
Distinct returns distinct elements from a parallel sequence by using the default equality comparer to compare values.

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

enumerable := lingo.AsParallelEnumerable(source).Distinct()
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 An} {3 A}] with unordered results
```

#### Except
Except returns the set difference, which means the elements of one parallel sequence that don't appear in a second parallel sequences.

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

enumerable := lingo.AsParallelEnumerable(source).Except(lingo.AsParallelEnumerable(second))
fmt.Println(enumerable.ToSlice())
// Result: [{3 A} {1 AKh} {2 Lah}] with unordered results
```

#### Intersect
Intersect returns the set intersection, which means elements that appear in each of two parallel sequences.

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

enumerable := lingo.AsParallelEnumerable(source).Intersect(lingo.AsParallelEnumerable(second))
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 hnA}] with unordered results
```

#### Union
Union returns the set union, which means unique elements that appear in either of two parallel sequences.

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

enumerable := lingo.AsParallelEnumerable(source).Union(lingo.AsParallelEnumerable(second))
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 hnA} {3 A} {3 Ab}] with unordered results
```