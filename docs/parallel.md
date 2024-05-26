## Parallel LINQ
Parallel LINQ (PLINQ) is a parallel implementation of the LINQ pattern. PLINQ has some additional operators for parallel operations. PLINQ combines the simplicity and readability of LINQ syntax with the power of parallel programming.

### Definition

```go
type ParallelEnumerable[T any] struct {
	getIter func() <-chan T
}
```

**Noted**: Note that ParallelEnumerable will be useful in some cases, you must really consider when using it.
In worse cases, not only will performance not be optimized, but costs will also increase.

Example:
```go
penum := lingo.Range(1, 100).AsParallel().Where(func(i int) bool {
	return i > 50
}) // not as good for PLINQ, should use Enumerable
```

```go
penum := lingo.Range(1, 100).AsParallel().Where(func(i int) bool {
	time.Sleep(time.Second)
	return i > 50
}) // good for PLINQ, performance will be improved
```


### Initialize PLINQ
#### AsParallel (Enumerable's method)
AsParallel creates a new ParallelEnumerable from an Enumerable

Example:
```go
penumerable := lingo.Range(1, 100).AsParallel()
```

#### AsParallelEnumerable
AsParallelEnumerable creates a new ParallelEnumerable

Example:
```go
penumerable := lingo.AsParallelEnumerable([]int{1, 2, 3})
```
#### GetIter
GetIter returns an unbuffered channel of T that iterates through a collection.

Example:
```go
penumerable := lingo.Range(1, 100).AsParallel()
iter := penumerable.GetIter()
```

#### Concat
Concat concatenates two sequences.

Example:
```go
first := lingo.Range(1, 10).AsParallel()
second := lingo.Range(11, 20).AsParallel()
first.Concat(second)
// [2 1 3 5 4 6 8 7 9 10 12 13 11 15 14 16 18 17 19 20]
```

### Retrieve data
#### FirstOrNil
FirstOrNil returns the first element of a parallel sequence (with condition if any), or a nil value if no element is found

Example:
```go
penumerable := lingo.Range(1, 10).AsParallel()
first := penumerable.FirstOrNil() // any element
```
#### FirstOrDefault
FirstOrDefault returns the first element of a parallel sequence (with condition if any), or a default value if no element is found

Example:
```go
penumerable := lingo.Empty[int]().AsParallel()
first := penumerable.FirstOrDefault(-999) // -999
```
#### LastOrNil
LastOrNil returns the last element of a parallel sequence (with condition if any), or a nil value if no element is found

Example:
```go
penumerable := lingo.Empty[int]().AsParallel()
last := enumerabpenumerablele.LastOrNil() // 0
```
#### LastOrDefault
LastOrDefault returns the last element of a parallel sequence (with condition if any), or a default value if no element is found

Example:
```go
penumerable := lingo.Empty[int]().AsParallel()
last := penumerable.LastOrDefault(999) // 999
```
#### ElementAtOrNil
ElementAtOrNil returns the element at a specified index in a parallel sequence or a default value if the index is out of range.

Example:
```go
enumerable := lingo.Range(1, 100).AsParallel()
element := enumerable.ElementAtOrNil(54) // any element
```
#### ElementAtOrDefault
ElementAtOrDefault returns the element at a specified index in a parallel sequence or a default value if the index is out of range.

Example:
```go
enumerable := lingo.Range(1, 100).AsParallel()
element := enumerable.ElementAtOrDefault(100, -1) // -1
```


### Filtering data
#### Where
Where filters in parallel a sequence of values based on a predicate.

Example:
```go
slice := lingo.Range(1, 100).AsParallel().Where(func(i int) bool {
	time.Sleep(time.Second)
	return i > 98
}).ToSlice() // [99, 100]
```
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

### Aggregation operations
#### MinBy
MinBy invokes in parallel a transform function on each element of a sequence and returns the minimum value.

Example:
```go
source := []Student{
    {Id: 1, Name: "Anh", ClassId: 1},
    {Id: 2, Name: "hnA", ClassId: 2},
    {Id: 3, Name: "Abcd", ClassId: 3},
    {Id: 44, Name: "Ank", ClassId: 1},
    {Id: 5, Name: "hnI", ClassId: 2},
    {Id: 6, Name: "A", ClassId: 3},
}
lingo.AsEnumerable(source).MinBy(func(s Student) any { return s.Id })
// {Id: 1, Name: "Anh", ClassId: 1}
```
#### MaxBy
MaxBy invokes in parallel a transform function on each element of a sequence and returns the maximum value.

Example:
```go
source := []Student{
    {Id: 1, Name: "Anh", ClassId: 1},
    {Id: 2, Name: "hnA", ClassId: 2},
    {Id: 3, Name: "Abcd", ClassId: 3},
    {Id: 44, Name: "Ank", ClassId: 1},
    {Id: 5, Name: "hnI", ClassId: 2},
    {Id: 6, Name: "A", ClassId: 3},
}
lingo.AsEnumerable(source).MaxBy(func(s Student) any { return s.Id })
// {Id: 44, Name: "Ank", ClassId: 1}
```
#### Sum
Sum computes in parallel the sum of the sequence of values that are obtained by invoking a transform function on each element of the input sequence.

Example:
```go
s := lingo.Range(1, 100).AsParallel().Where(func(i int) bool {
	return i > 98
}).Sum(func(i int) any {
	time.Sleep(time.Second)
	return i * i
})
fmt.Println(s) // 19801
```
#### Average
Average computes in parallel the average of a sequence of numeric values that are obtained by invoking a transform function on each element of the input sequence.

Example:
```go
a := lingo.Range(1, 100).AsParallel().Where(func(i int) bool {
	return i > 50
}).Average(func(i int) any {
	time.Sleep(time.Second)
	return i * i
})
fmt.Println(a) // 5908.5
```

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

### Partitioning data
#### Skip
Skip skips a specified number of elements in a parallel sequence and then returns the remaining elements.
If the source sequence is ordered, Skip skips first n elements. On the other hand, Skip skips any n elements.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsParallelEnumerable(source).Skip(2)
fmt.Println(enumerable.ToSlice())
// Result: [{3 A}] (in somecase)
```
#### SkipWhile
SkipWhile elements in a parallel sequence as long as a specified condition is true and then returns the remaining elements.
If the source sequence is ordered, SkipWhile skips according to the ordered element. On the other hand, performs SkipWhile on the current arbitrary order.

Example:

```go
source := []Student{
    {Id: 1, Name: "Anh"},
    {Id: 2, Name: "hnA"},
    {Id: 3, Name: "A"},
}

enumerable := lingo.AsParallelEnumerable(source).SkipWhile(func(s Student) bool {
    return s.Id < 2
})
fmt.Println(enumerable.ToSlice())
// Result: [{2 hnA} {1 Anh} {3 A}] (in somecase)
```
#### Take
Take returns a specified number of contiguous elements from the start of a parallel sequence.
If the source sequence is ordered, Take takes first n elements. On the other hand, Take takes any n elements.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsParallelEnumerable(source).Take(2)
fmt.Println(enumerable.ToSlice())
// Result: [{2 hnA} {1 Anh}] (in somecase)
```
#### TakeWhile
TakeWhile takes elements in a parallel sequence based on a predicate function until an element doesn't satisfy the condition.
If the source sequence is ordered, TakeWhile takes according to the ordered element. On the other hand, performs TakeWhile on the current arbitrary order.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsParallelEnumerable(source).TakeWhile(func(s Student) bool {
	return s.Id < 3
})
fmt.Println(enumerable.ToSlice())
// Result: [{2 hnA} {1 Anh}] (in somecase)
```

### Grouping data
GroupBy groups in parallel elements that share a common attribute.

Example: the following example will demonstrate the performance improvement of parallel, without using PLINQ it will take about 24s (11*2 + 2). but when using PLINQ it takes about 2s. That sounds great
```go
type args struct {
	keySelector     definition.SingleSelector[Student]
	elementSelector definition.SingleSelector[Student]
	resultSelector  definition.GroupBySelector[any, any]
	getHash         definition.GetHashCode[any]
}
source := []Student{
	{Id: 1, Name: "A", Score: 7, ClassId: 1},
	{Id: 2, Name: "B", Score: 3, ClassId: 2},
	{Id: 3, Name: "C", Score: 10, ClassId: 2},
	{Id: 4, Name: "D", Score: 2, ClassId: 1},
	{Id: 5, Name: "E", Score: 8, ClassId: 2},
	{Id: 6, Name: "F", Score: 6, ClassId: 1},
	{Id: 7, Name: "G", Score: 5, ClassId: 1},
	{Id: 8, Name: "H", Score: 8, ClassId: 2},
	{Id: 9, Name: "J", Score: 3, ClassId: 1},
	{Id: 10, Name: "K", Score: 4, ClassId: 2},
	{Id: 12, Name: "Q", Score: 8, ClassId: 2},
}
arg := args{
	keySelector: func(s Student) any {
		time.Sleep(time.Second)
		return s.ClassId
	},
	elementSelector: func(s Student) any {
		time.Sleep(time.Second)
		return s
	},
	resultSelector: func(a1 any, a2 []any) any {
		time.Sleep(time.Second)
		ss := lingo.AsEnumerableTFromSliceAny[Student](a2).OrderByDescending(func(s Student) any {
			return s.Score
		}).ToSlice()
		s := ss[0]
		return Student{
			Id:      s.Id,
			Name:    s.Name,
			Score:   s.Score,
			ClassId: a1.(int),
		}
	},
}

enumerable := lingo.AsEnumerable(source).
	AsParallel().
	GroupBy(
		arg.keySelector,
		arg.elementSelector,
		arg.resultSelector,
		arg.getHash,
	)
fmt.Println(enumerable.ToSlice())
// [{1 A 7 1} {3 C 10 2}]
```

### Initialize lookup
AsPLookup creates a generic Lookup[K, V] from an ParallelEnumerable[T].

```go
source := []Student{
	{Id: 1, Name: "Anh", ClassId: 1},
	{Id: 2, Name: "hnA", ClassId: 2},
	{Id: 3, Name: "Abcd", ClassId: 3},
	{Id: 4, Name: "Ank", ClassId: 1},
	{Id: 5, Name: "hnI", ClassId: 2},
	{Id: 6, Name: "A", ClassId: 3},
}
lookup := lingo.AsPLookup[Student, int, Student](lingo.AsParallelEnumerable(source), func(s Student) int {
	time.Sleep(time.Second)
	return s.ClassId
})

for grouping := range lookup.OrderByDescending(func(g lingo.Grouping[int, Student]) any { return g.Key }).GetIter() {
	fmt.Printf("%d:\n", grouping.Key)
	for v := range grouping.GetIter() {
		fmt.Println("\t", v)
	}
}
// Result:
// 3:
// 		{3 Abcd 3}
// 		{6 A 3}
// 2:
// 		{2 hnA 2}
// 		{5 hnI 2}
// 1:
// 		{1 Anh 1}
// 		{4 Ank 1}
```