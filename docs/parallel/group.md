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