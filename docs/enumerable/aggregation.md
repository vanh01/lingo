### Aggregation operations
#### Min
Min returns the minimum value in a sequence of values.
In this method, comparer is returns whether left is smaller than right or not. The left one will be returned
If comparer is empty or nil, we will use the default comparer. On the other hand, we just use the first comparer

Example:
```go
enumerable := lingo.Range(1, 100)
m := enumerable.Min() // 1
```
#### MinBy
MinBy returns the minimum value in a sequence of values according to a specified key selector function.
In this method, comparer is returns whether left is smaller than right or not. The left one will be returned
If comparer is empty or nil, we will use the default comparer. On the other hand, we just use the first comparer

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
#### Max
Max returns the minimum value in a sequence of values.
In this method, comparer is returns whether left is greater than right or not. The left one will be returned
If comparer is empty or nil, we will use the default comparer. On the other hand, we just use the first comparer

Example:
```go
enumerable := lingo.Range(1, 100)
m := enumerable.Max() // 100
```
#### MaxBy
MaxBy returns the maximum value in a sequence of values according to a specified key selector function.
In this method, comparer is returns whether left is greater than right or not. The left one will be returned
If comparer is empty or nil, we will use the default comparer. On the other hand, we just use the first comparer

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
Sum computes the sum of a sequence of numeric values.

Example:
```go
enumerable := lingo.Range(1, 100)
sum := enumerable.Sum() // 5050
```
#### Average
Average computes the average of a sequence of numeric values.

Example:
```go
enumerable := lingo.Range(1, 100)
sum := enumerable.Average() // 50.5
```
#### Count
Count returns the number of elements in a sequence.

Example:
```go
enumerable := lingo.Range(1, 100)
count := enumerable.Count() // 100
```

#### Aggregate
Aggregate applies an accumulator function over a sequence. The specified seed value is used as the initial accumulator value, and the specified function is used to select the result value.

Example:
```go
sentence := "the quick brown fox jumps over the lazy dog"
word := strings.Split(sentence, " ")
reversed := lingo.AsEnumerable(word).Aggregate(
	"",
	func(a any, s string) any {
		return s + " " + a.(string)
	},
)
fmt.Println(reversed)
// Result: "dog lazy the over jumps fox brown quick the"
```