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