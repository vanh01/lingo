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