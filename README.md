# lingo

[![go report card](https://goreportcard.com/badge/github.com/vanh01/lingo "go report card")](https://goreportcard.com/report/github.com/vanh01/lingo)
[![test status](https://github.com/vanh01/lingo/actions/workflows/test.yml/badge.svg "test status")](https://github.com/vanh01/lingo/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/vanh01/lingo.svg)](https://pkg.go.dev/github.com/vanh01/lingo)

lingo is a library written in Go. It is LinQ in .NET for Go. It will help the array processing code more neat.

**The key features of lingo are:**
- Enumerable
	- [Initialize enumerable](./docs/enumerable.md#initialize-enumerable)
	- [Initialize lookup](./docs/lookup.md#initialize-lookup)
	- [Modification enumerable](./docs/modification.md#modification-enumerable)
	- [Retrieve data](./docs/retrieve.md#retrieve-data)
	- [Filtering data](./docs/filter.md#filtering-data)
	- [Projection operations](./docs/projection.md#projection-operations)
	- [Aggregation operations](./docs/aggregation.md#aggregation-operations)
	- [Set operations](./docs/set.md#set-operations)
	- [Sorting data](./docs/sort.md#sorting-data)
	- [Quantifier operations](./docs/quantifier.md#quantifier-operations)
	- [Partitioning data](./docs/partition.md#partitioning-data)
	- [Convert data types](./docs/converter.md#converting-data-types)
	- [Join operations](./docs/join.md#join-operations)
	- [Group data](./docs/group.md#grouping-data)
- [Parallel](./docs/parallel.md)
	- [Definition](./docs/parallel.md#definition)
	- [Initialize PLINQ](./docs/parallel.md#initialize-plinq)
	- [Retrieve data](./docs/parallel.md#retrieve-data)
	- [Filtering data](./docs/parallel.md#filtering-data)
	- [Quantifier operations](./docs/parallel.md#quantifier-operations)
	- [Aggregation operations](./docs/parallel.md#aggregation-operations)
	- [Set operations](./docs/parallel.md#set-operations)
	- [Partitioning data](./docs/parallel.md#partitioning-data)
	- [Group data](./docs/parallel.md#grouping-data)

## Getting started

### Prerequisites

- **[Go](https://go.dev/)**: version is newer than 1.18

### Installation

```sh
go get -u github.com/vanh01/lingo
```

### Quick Start

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

## Features

### Classification table
The following table classifies each supported method, there are two types: Immediate execution and Chain execution.

- Immediate execution terminates an enumerable and returns the result
- Chain execution can be called many times.

| Standard query operator | Return type | Immediate execution | Chain execution |
|-|-|-|-|
|[Min](./docs/aggregation.md#min)|T|x||
|[MinBy](./docs/aggregation.md#minby)|T|x||
|[Max](./docs/aggregation.md#max)|T|x||
|[MaxBy](./docs/aggregation.md#maxby)|T|x||
|[Sum](./docs/aggregation.md#sum)|Number|x||
|[Average](./docs/aggregation.md#average)|float64|x||
|[Count](./docs/aggregation.md#count)|int64|x||
|[Aggregate](./docs/aggregation.md#aggregate)|any|x||
|[ToSlice](./docs/converter.md#toslice)|[]T|x||
|[ToMap](./docs/converter.md#tomap)|map[any]any|x||
|[Concat](./docs/enumerable.md#concat)|Enumerable[T]||x|
|[Where](./docs/filter.md#filtering-data)|Enumerable[T]||x|
|[GroupBy](./docs/group.md#grouping-data)|Enumerable[any]||x|
|[Join](./docs/join.md#join-operations)|Enumerable[any]||x|
|[Skip](./docs/partition.md#skip)|Enumerable[T]||x|
|[SkipWhile](./docs/partition.md#skipwhile)|Enumerable[T]||x|
|[Take](./docs/partition.md#take)|Enumerable[T]||x|
|[TakeWhile](./docs/partition.md#takewhile)|Enumerable[T]||x|
|[Select](./docs/projection.md#select)|Enumerable[any]||x|
|[SelectMany](./docs/projection.md#selectmany)|Enumerable[any]||x|
|[Zip](./docs/projection.md#zip)|Enumerable[any]||x|
|[All](./docs/quantifier.md#all)|bool|x||
|[Any](./docs/quantifier.md#any)|bool|x||
|[Contains](./docs/quantifier.md#contains)|bool|x||
|[FirstOrNil](./docs/retrieve.md#firstornil)|T|x||
|[FirstOrDefault](./docs/retrieve.md#firstordefault)|T|x||
|[LastOrNil](./docs/retrieve.md#lastornil)|T|x||
|[LastOrDefault](./docs/retrieve.md#lastordefault)|T|x||
|[ElementAtOrNil](./docs/retrieve.md#elementatornil)|T|x||
|[ElementAtOrDefault](./docs/retrieve.md#elementatordefault)|T|x||
|[Distinct](./docs/set.md#distinct)|Enumerable[T]||x|
|[DistinctBy](./docs/set.md#distinctby)|Enumerable[T]||x|
|[Except](./docs/set.md#except)|Enumerable[T]||x|
|[ExceptBy](./docs/set.md#exceptby)|Enumerable[T]||x|
|[Intersect](./docs/set.md#intersect)|Enumerable[T]||x|
|[IntersectBy](./docs/set.md#intersectby)|Enumerable[T]||x|
|[Union](./docs/set.md#union)|Enumerable[T]||x|
|[UnionBy](./docs/set.md#unionby)|Enumerable[T]||x|
|[OrderBy](./docs/sort.md#orderby)|Enumerable[T]||x|
|[OrderByDescending](./docs/sort.md#orderbydescending)|Enumerable[T]||x|
|[Reverse](./docs/sort.md#reverse)|Enumerable[T]||x|
|[Append](./docs/modification.md#append)|Enumerable[T]||x|
|[AppendRange](./docs/modification.md#appendrange)|Enumerable[T]||x|
|[Prepend](./docs/modification.md#prepend)|Enumerable[T]||x|
|[PrependRange](./docs/modification.md#prependrange)|Enumerable[T]||x|
|[Clear](./docs/modification.md#clear)|Enumerable[T]||x|
|[Insert](./docs/modification.md#insert)|Enumerable[T]||x|
|[Remove](./docs/modification.md#remove)|Enumerable[T]||x|
|[RemoveAt](./docs/modification.md#removeat)|Enumerable[T]||x|
|[RemoveRange](./docs/modification.md#removerange)|Enumerable[T]||x|

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

Please click [here](./docs/enumerable.md)

## Reference
- [LinQ](https://learn.microsoft.com/en-us/dotnet/csharp/linq/)

## Contributors
Thanks to all the people who already contributed!
