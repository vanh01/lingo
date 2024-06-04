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