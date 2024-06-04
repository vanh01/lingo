### Join operations
Join joins two sequences based on key selector functions and extracts pairs of values.

Example:

```go
type StudentClass struct {
	Id        int
	Name      string
	ClassId   int
	ClassName string
	Total     int
}
type args struct {
	inner            []any
	innerKeySelector lingo.SingleSelector[any]
	outerKeySelector lingo.SingleSelector[Student]
	resultSelector   lingo.CombinationSelector[Student, any]
	comparer         lingo.Comparer[any]
}
source := []Student{
	{Id: 1, Name: "A", ClassId: 1},
	{Id: 2, Name: "B", ClassId: 2},
	{Id: 3, Name: "C", ClassId: 2},
	{Id: 4, Name: "D", ClassId: 1},
	{Id: 5, Name: "E", ClassId: 2},
	{Id: 6, Name: "F", ClassId: 1},
	{Id: 7, Name: "G", ClassId: 1},
	{Id: 8, Name: "H", ClassId: 2},
	{Id: 9, Name: "J", ClassId: 1},
	{Id: 10, Name: "K", ClassId: 2},
	{Id: 12, Name: "Q", ClassId: 2},
}
arg := args{
	inner: []any{
		Class{Id: 1, Name: "Class1", Total: 20},
		Class{Id: 2, Name: "Class2", Total: 10},
	},
	outerKeySelector: func(s Student) any { return s.ClassId },
	innerKeySelector: func(c any) any { return c.(Class).Id },
	resultSelector: func(s Student, c any) any {
		return StudentClass{
			Id:        s.Id,
			Name:      s.Name,
			ClassId:   s.ClassId,
			ClassName: c.(Class).Name,
			Total:     c.(Class).Total,
		}
	},
}

enumerable := lingo.AsEnumerable(source).
	Join(
		lingo.AsEnumerable(arg.inner),
		arg.outerKeySelector, arg.innerKeySelector,
		arg.resultSelector,
		arg.comparer,
	)
fmt.Println(enumerable.ToSlice())
// Result: [{1 A 1 Class1 20} {2 B 2 Class2 10} {3 C 2 Class2 10} {4 D 1 Class1 20} {5 E 2 Class2 10} {6 F 1 Class1 20} {7 G 1 Class1 20} {8 H 2 Class2 10} {9 J 1 Class1 20} {10 K 2 Class2 10} {12 Q 2 Class2 10}]
```