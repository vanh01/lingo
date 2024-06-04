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