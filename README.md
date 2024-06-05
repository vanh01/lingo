# lingo

[![go report card](https://goreportcard.com/badge/github.com/vanh01/lingo "go report card")](https://goreportcard.com/report/github.com/vanh01/lingo)
[![test status](https://github.com/vanh01/lingo/actions/workflows/test.yml/badge.svg "test status")](https://github.com/vanh01/lingo/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/vanh01/lingo.svg)](https://pkg.go.dev/github.com/vanh01/lingo)

lingo is a library written in Go. It is LinQ in .NET for Go. It will help the array processing code more neat.

**The key features of lingo are:**
- [Enumerable](./docs/enumerable/enumerable.md)
	- [Initialize enumerable](./docs/enumerable/enumerable.md#initialize-enumerable)
	- [Initialize lookup](./docs/enumerable/lookup.md#initialize-lookup)
	- [Modification enumerable](./docs/enumerable/modification.md#modification-enumerable)
	- [Retrieve data](./docs/enumerable/retrieve.md#retrieve-data)
	- [Filtering data](./docs/enumerable/filter.md#filtering-data)
	- [Projection operations](./docs/enumerable/projection.md#projection-operations)
	- [Aggregation operations](./docs/enumerable/aggregation.md#aggregation-operations)
	- [Set operations](./docs/enumerable/set.md#set-operations)
	- [Sorting data](./docs/enumerable/sort.md#sorting-data)
	- [Quantifier operations](./docs/enumerable/quantifier.md#quantifier-operations)
	- [Partitioning data](./docs/enumerable/partition.md#partitioning-data)
	- [Convert data types](./docs/enumerable/converter.md#converting-data-types)
	- [Join operations](./docs/enumerable/join.md#join-operations)
	- [Group data](./docs/enumerable/group.md#grouping-data)
- [Parallel](./docs/parallel/parallel.md)
	- [Definition](./docs/parallel/parallel.md#definition)
	- [Initialize PLINQ](./docs/parallel/parallel.md#initialize-plinq)
	- [Retrieve data](./docs/parallel/retrieve.md#retrieve-data)
	- [Sorting data](./docs/parallel/sort.md#sorting-data)
	- [Filtering data](./docs/parallel/filter.md#filtering-data)
	- [Quantifier operations](./docs/parallel/quantifier.md#quantifier-operations)
	- [Aggregation operations](./docs/parallel/aggregation.md#aggregation-operations)
	- [Set operations](./docs/parallel/set.md#set-operations)
	- [Partitioning data](./docs/parallel/partition.md#partitioning-data)
	- [Join operations](./docs/parallel/join.md#join-operations)
	- [Group data](./docs/parallel/group.md#grouping-data)

## Getting started

### Prerequisites

- **[Go](https://go.dev/)**: version is newer than 1.18

### Installation

```sh
go get -u github.com/vanh01/lingo
```

### Quick Start
<details><summary>Example</summary>
For example, I want to get top 3 students in a class whose name contains "1"

```go
package main

import (
	"fmt"

	lingo "github.com/vanh01/lingo"
)

type Student struct {
	Id      int
	Name    string
	Score	int
	ClassId int
}

type Class struct {
	Id    int
	Name  string
	Total int
}

func main() {
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

	classes := []Class{
		{Id: 1, Name: "Class 1", Total: 20},
		{Id: 2, Name: "Class 2", Total: 10},
	}
	classIds := lingo.AsEnumerable(classes).
		Where(func(c Class) bool {
			return strings.Contains(c.Name, "1")
		}).
		Select(func(c Class) any {
			return c.Id
		}).
		ToSlice()

	enumerable := lingo.AsEnumerable(source).
		Where(func(t Student) bool {
			return lingo.AsEnumerable(classIds).Contains(t.ClassId)
		}).
		OrderByDescending(func(s Student) any {
			return s.Score
		}).
		Take(3)
	fmt.Println(enumerable.ToSlice())
}
// Result: [{9 J 10 1} {4 D 7 1} {1 A 6 1}]
```
</details>

## Features

### Classification table
The following table classifies each supported method, there are two types: Immediate execution and Chain execution.

- Immediate execution terminates an enumerable and returns the result
- Chain execution can be called many times.

| Standard query operator | Return type | Immediate execution | Chain execution | Parallel support |
|-|-|-|-|-|
|[Min](./docs/enumerable/aggregation.md#min)|T|x|||
|[MinBy](./docs/enumerable/aggregation.md#minby)|T|x||x|
|[Max](./docs/enumerable/aggregation.md#max)|T|x|||
|[MaxBy](./docs/enumerable/aggregation.md#maxby)|T|x||x|
|[Sum](./docs/enumerable/aggregation.md#sum)|Number|x||x|
|[Average](./docs/enumerable/aggregation.md#average)|float64|x||x|
|[Count](./docs/enumerable/aggregation.md#count)|int64|x|||
|[Aggregate](./docs/enumerable/aggregation.md#aggregate)|any|x|||
|[ToSlice](./docs/enumerable/converter.md#toslice)|[]T|x||x|
|[ToMap](./docs/enumerable/converter.md#tomap)|map[any]any|x||x|
|[Concat](./docs/enumerable/enumerable.md#concat)|Enumerable[T]||x|x|
|[Where](./docs/enumerable/filter.md#filtering-data)|Enumerable[T]||x|x|
|[GroupBy](./docs/enumerable/group.md#grouping-data)|Enumerable[any]||x|x|
|[Join](./docs/enumerable/join.md#join-operations)|Enumerable[any]||x|x|
|[Skip](./docs/enumerable/partition.md#skip)|Enumerable[T]||x|x|
|[SkipWhile](./docs/enumerable/partition.md#skipwhile)|Enumerable[T]||x|x|
|[Take](./docs/enumerable/partition.md#take)|Enumerable[T]||x|x|
|[TakeWhile](./docs/enumerable/partition.md#takewhile)|Enumerable[T]||x|x|
|[Select](./docs/enumerable/projection.md#select)|Enumerable[any]||x|x|
|[SelectMany](./docs/enumerable/projection.md#selectmany)|Enumerable[any]||x|x|
|[Zip](./docs/enumerable/projection.md#zip)|Enumerable[any]||x|x|
|[All](./docs/enumerable/quantifier.md#all)|bool|x||x|
|[Any](./docs/enumerable/quantifier.md#any)|bool|x||x|
|[Contains](./docs/enumerable/quantifier.md#contains)|bool|x||x|
|[FirstOrNil](./docs/enumerable/retrieve.md#firstornil)|T|x||x|
|[FirstOrDefault](./docs/enumerable/retrieve.md#firstordefault)|T|x||x|
|[LastOrNil](./docs/enumerable/retrieve.md#lastornil)|T|x||x|
|[LastOrDefault](./docs/enumerable/retrieve.md#lastordefault)|T|x||x|
|[ElementAtOrNil](./docs/enumerable/retrieve.md#elementatornil)|T|x||x|
|[ElementAtOrDefault](./docs/enumerable/retrieve.md#elementatordefault)|T|x||x|
|[Distinct](./docs/enumerable/set.md#distinct)|Enumerable[T]||x|x|
|[DistinctBy](./docs/enumerable/set.md#distinctby)|Enumerable[T]||x||
|[Except](./docs/enumerable/set.md#except)|Enumerable[T]||x|x|
|[ExceptBy](./docs/enumerable/set.md#exceptby)|Enumerable[T]||x||
|[Intersect](./docs/enumerable/set.md#intersect)|Enumerable[T]||x|x|
|[IntersectBy](./docs/enumerable/set.md#intersectby)|Enumerable[T]||x||
|[Union](./docs/enumerable/set.md#union)|Enumerable[T]||x|x|
|[UnionBy](./docs/enumerable/set.md#unionby)|Enumerable[T]||x|x|
|[OrderBy](./docs/enumerable/sort.md#orderby)|Enumerable[T]||x|x|
|[OrderByDescending](./docs/enumerable/sort.md#orderbydescending)|Enumerable[T]||x|x|
|[Reverse](./docs/enumerable/sort.md#reverse)|Enumerable[T]||x|x|
|[Append](./docs/enumerable/modification.md#append)|Enumerable[T]||x||
|[AppendRange](./docs/enumerable/modification.md#appendrange)|Enumerable[T]||x||
|[Prepend](./docs/enumerable/modification.md#prepend)|Enumerable[T]||x||
|[PrependRange](./docs/enumerable/modification.md#prependrange)|Enumerable[T]||x||
|[Clear](./docs/enumerable/modification.md#clear)|Enumerable[T]||x||
|[Insert](./docs/enumerable/modification.md#insert)|Enumerable[T]||x||
|[Remove](./docs/enumerable/modification.md#remove)|Enumerable[T]||x||
|[RemoveAt](./docs/enumerable/modification.md#removeat)|Enumerable[T]||x||
|[RemoveRange](./docs/enumerable/modification.md#removerange)|Enumerable[T]||x||

I have defined the Student structure for the following examples

```go
type Student struct {
	Id      int
	Name    string
	ClassId int
}

type Class struct {
	Id    int
	Name  string
	Total int
}
```
Following are all the examples related to the supported methods in this library.

Please click [here](./docs/enumerable/enumerable.md)

## Reference
- [LinQ](https://learn.microsoft.com/en-us/dotnet/csharp/linq/)

## Contributors
Thanks to all the people who already contributed!
