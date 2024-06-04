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