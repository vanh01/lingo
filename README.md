# lingo

[![go report card](https://goreportcard.com/badge/github.com/vanh01/lingo "go report card")](https://goreportcard.com/report/github.com/vanh01/lingo)
[![test status](https://github.com/vanh01/lingo/actions/workflows/test.yml/badge.svg "test status")](https://github.com/vanh01/lingo/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/vanh01/lingo.svg)](https://pkg.go.dev/github.com/vanh01/lingo)

lingo is a library written in Go. It is LinQ in .NET for Go. It will help the array processing code more neat.

**The key features of lingo are:**
- [Initialize enumerable](#initialize-enumerable)
- [Modification enumerable](#modification-enumerable)
- [Retrieve data](#retrieve-data)
- [Filtering data](#filtering-data)
- [Projection operations](#projection-operations)
- [Aggregation operations](#aggregation-operations)
- [Set operations](#set-operations)
- [Sorting data](#sorting-data)
- [Quantifier operations](#quantifier-operations)
- [Partitioning data](#partitioning-data)
- [Convert data types](#converting-data-types)
- [Join operations](#join-operations)
- [Group data](#grouping-data)

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
			return lingo.AsEnumerable(classIds).Contains(t.ClassId, nil)
		}).
		OrderByDescending(func(s Student) any {
			return s.Score
		}, nil).
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
|[Min](#min)|T|x||
|[Max](#max)|T|x||
|[Sum](#sum)|Number|x||
|[Average](#average)|float64|x||
|[Count](#count)|int64|x||
|[Aggregate](#aggregate)|any|x||
|[ToSlice](#toslice)|[]T|x||
|[ToMap](#tomap)|map[any]any|x||
|[Concat](#concat)|Enumerable[T]||x|
|[Where](#filtering-data)|Enumerable[T]||x|
|[GroupBy](#grouping-data)|Enumerable[any]||x|
|[Join](#join-operations)|Enumerable[any]||x|
|[Skip](#skip)|Enumerable[T]||x|
|[SkipWhile](#skipwhile)|Enumerable[T]||x|
|[Take](#take)|Enumerable[T]||x|
|[TakeWhile](#takewhile)|Enumerable[T]||x|
|[Select](#select)|Enumerable[any]||x|
|[SelectMany](#selectmany)|Enumerable[any]||x|
|[Zip](#zip)|Enumerable[any]||x|
|[All](#all)|bool|x||
|[Any](#any)|bool|x||
|[Contains](#contains)|bool|x||
|[FirstOrNil](#firstornil)|T|x||
|[FirstOrDefault](#firstordefault)|T|x||
|[LastOrNil](#lastornil)|T|x||
|[LastOrDefault](#lastordefault)|T|x||
|[ElementAtOrNil](#elementatornil)|T|x||
|[ElementAtOrDefault](#elementatordefault)|T|x||
|[Distinct](#distinct)|Enumerable[T]||x|
|[DistinctBy](#distinctby)|Enumerable[T]||x|
|[Except](#except)|Enumerable[T]||x|
|[Intersect](#intersect)|Enumerable[T]||x|
|[Union](#union)|Enumerable[T]||x|
|[OrderBy](#orderby)|Enumerable[T]||x|
|[OrderByDescending](#orderbydescending)|Enumerable[T]||x|
|[Reverse](#reverse)|Enumerable[T]||x|
|[Append](#append)|Enumerable[T]||x|
|[AppendRange](#appendrange)|Enumerable[T]||x|
|[Prepend](#prepend)|Enumerable[T]||x|
|[PrependRange](#prependrange)|Enumerable[T]||x|
|[Clear](#clear)|Enumerable[T]||x|
|[Insert](#insert)|Enumerable[T]||x|
|[Remove](#remove)|Enumerable[T]||x|
|[RemoveAt](#removeat)|Enumerable[T]||x|
|[RemoveRange](#removerange)|Enumerable[T]||x|




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

### Initialize enumerable
#### AsEnumerable
AsEnumerable creates a new Enumerable from a silce of specific type.

Example:
```go
enumerable := lingo.AsEnumerable([]int{1, 2, 3})
```
#### AsEnumerableTFromAny
AsEnumerableTFromAny creates a new Enumerable of specific type from Enumerable of any type, this method will be useful after using projection operations.

Example:
```go
enumerable := lingo.AsEnumerable([]any{1, 2, 3})
enumerableInt := lingo.AsEnumerableTFromAny[int](enumerable)
```
#### AsEnumerableTFromSliceAny
AsEnumerableTFromSliceAny creates a new Enumerable of specific type from slice of any type

Example:
```go
enumerableInt := lingo.AsEnumerableTFromSliceAny[int]([]any{1, 2, 3})
```
#### Empty
Empty returns an empty Enumerable[T] that has the specified type argument.

Example:
```go
emptyInt := lingo.Empty[int]()
```
#### Range
Range generates a sequence of integral numbers within a specified range.

Example: create an enumerable of int from 1 to 10
```go
rangeInt := lingo.Range(1, 10)
```
#### Repeat
Repeat generates a sequence that contains one repeated value.

Example: create a enumerable of int, it contains 10 elements 1
```go
repeatInt := lingo.Repeat(1, 10)
```
#### Concat
Concat concatenates two sequences.

Example:
```go
first := lingo.Range(1, 10)
second := lingo.Range(11, 20)
first.Concat(second) // 1-20
```

### Modification enumerable
#### Append
Append appends a value to the end of the sequence.

Example:
```go
enum := lingo.Range(1, 10)
enum.Append(99).ToSlice() // 1-10, 99
```

#### AppendRange
AppendRange appends the elements of the specified collection to the end of the sequence.

Example:
```go
enum := lingo.Range(1, 10)
enum.AppendRange(lingo.Range(100, 110)).ToSlice() // 1-10, 100-110
```

#### Prepend
Prepend adds a value to the beginning of the sequence.

Example:
```go
enum := lingo.Range(1, 10)
enum.Prepend(99).ToSlice() // 99, 1-10
```

#### PrependRange
PrependRange appends the elements of the specified collection to the beginning of the sequence.

Example:
```go
enum := lingo.Range(1, 10)
enum.PrependRange(lingo.Range(100, 110)).ToSlice() // 100-110, 1-10
```

#### Clear
Clear removes all elements of the sequence.

Example:
```go
enum := lingo.Range(1, 10)
enum.Clear().ToSlice() // []
```

#### Insert
Insert inserts an element into the sequence at the specified index.

Example:
```go
enum := lingo.Range(1, 10)
enum.Insert(3, 999).ToSlice() // 1-999, 4-10
```

#### Remove
Remove removes the first occurrence of the given element, if found.

Example:
```go
enum := lingo.Range(1, 10)
enum.Remove(9, nil).ToSlice() // 1-8, 10
```

#### RemoveAt
RemoveAt removes the element at the specified index of the sequence.

Example:
```go
enum := lingo.Range(1, 10)
enum.RemoveAt(6).ToSlice() // 1-6, 8-10
```

#### RemoveRange
RemoveRange removes a range of elements from the sequence.

Example:
```go
enum := lingo.Range(1, 10)
enum.RemoveRange(3, 3).ToSlice() // 1-3, 7-10
```

### Retrieve data
#### FirstOrNil
FirstOrNil returns the first element of a sequence (with condition if any), or a nil value if no element is found

Example:
```go
enumerable := lingo.Range(1, 10)
first := enumerable.FirstOrNil(nil) // 1
```
#### FirstOrDefault
FirstOrDefault returns the first element of a sequence (with condition if any), or a default value if no element is found

Example:
```go
enumerable := lingo.Empty[int]()
first := enumerable.FirstOrDefault(-999, nil) // -999
```
#### LastOrNil
LastOrNil returns the last element of a sequence (with condition if any), or a nil value if no element is found

Example:
```go
enumerable := lingo.Empty[int]()
last := enumerable.LastOrNil(nil) // 0
```
#### LastOrDefault
LastOrDefault returns the last element of a sequence (with condition if any), or a default value if no element is found

Example:
```go
enumerable := lingo.Empty[int]()
last := enumerable.LastOrDefault(999, nil) // 999
```
#### ElementAtOrNil
ElementAtOrNil returns the element at a specified index in a sequence or a default value if the index is out of range.

Example:
```go
enumerable := lingo.Range(1, 100)
element := enumerable.ElementAtOrNil(54) // 55
```
#### ElementAtOrDefault
ElementAtOrDefault returns the element at a specified index in a sequence or a default value if the index is out of range.

Example:
```go
enumerable := lingo.Range(1, 100)
element := enumerable.ElementAtOrDefault(100, -1) // -1
```


### Filtering data

I want to take all students who have the student name contains "an"

```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "An"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).
	Where(func(t Student) bool {
		return strings.Contains(strings.ToLower(t.Name), "an")
	})
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 An}]
```

### Projection operations

#### Select
Select projects values that are based on a transform function.

Example: Take all Ids from students
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "An"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).Select(func(t Student) any {
	return t.Id
})
fmt.Println(enumerable.ToSlice())
// Result: [1 2 3]
```

#### SelectMany
SelectMany projects sequences of values that are based on a transform function and then flattens them into one sequence.

Example: Take all Ids from slice-in-slice students
```go
source := [][]Student{
	{
		{Id: 1, Name: "Anh"},
		{Id: 11, Name: "An"},
	},
	{
		{Id: 2, Name: "Anh"},
		{Id: 21, Name: "An"},
	},
	{
		{Id: 3, Name: "Anh"},
		{Id: 31, Name: "An"},
	},
}

enumerable1 := lingo.AsEnumerable(source).SelectMany(func(ss []Student) []any {
	return lingo.AsEnumerable(ss).Select(func(s Student) any {
		return s.Id
	}).ToSlice()
})
fmt.Println(enumerable1.ToSlice())
// Result: [1 11 2 21 3 31]
```

#### Zip
Zip produces a sequence of tuples with elements from 2 specified sequences.
If resultSelector is nil, the default result is a slice combined with each element

Examples:

Take all Ids from slice-in-slice students
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "An"},
	{Id: 3, Name: "A"},
}
classIds := []int{11, 22, 33}

enumerable := lingo.AsEnumerable(source).Zip(classIds, func(s Student, a any) any {
	return struct {
		Id      int
		Name    string
		ClassId int
	}{
		Id:      s.Id,
		Name:    s.Name,
		ClassId: a.(int),
	}
})
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh 11} {2 An 22} {3 A 33}]
```

In case resultSelector is nil
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "An"},
	{Id: 3, Name: "A"},
}
classIds := []int{11, 22, 33}

enumerable := lingo.AsEnumerable(source).Zip(classIds, nil)
fmt.Println(enumerable.ToSlice())
// Result: [[{1 Anh} 11] [{2 An} 22] [{3 A} 33]]
```

### Aggregation operations
#### Min
Min returns the minimum value in a sequence of values. In this method, comparer is returns whether left is smaller than right or not, if comparer is nill, we will use the default comparer.

Example:
```go
enumerable := lingo.Range(1, 100)
m := enumerable.Min(nil) // 1
```
#### Max
Max returns the minimum value in a sequence of values.
In this method, comparer is returns whether left is greater than right or not, if comparer is nill, we will use the default comparer.

Example:
```go
enumerable := lingo.Range(1, 100)
m := enumerable.Max(nil) // 100
```
#### Sum
Sum computes the sum of a sequence of numeric values.

Example:
```go
enumerable := lingo.Range(1, 100)
sum := enumerable.Sum(nil) // 5050
```
#### Average
Average computes the average of a sequence of numeric values.

Example:
```go
enumerable := lingo.Range(1, 100)
sum := enumerable.Average(nil) // 50.5
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
	nil,
)
fmt.Println(reversed)
// Result: "dog lazy the over jumps fox brown quick the"
```

### Set operations
#### Distinct
Distinct removes duplicate values from a collection.

Example:

Get all unique students
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "An"},
	{Id: 3, Name: "A"},
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "An"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).Distinct()
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 An} {3 A}]
```

#### DistinctBy
DistinctBy removes duplicate values from a collection with keySelector and comparer, comparer can be nil without special comparer.

Examples:

Get all unique students by length of name
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
	{Id: 1, Name: "AKh"},
	{Id: 2, Name: "Lah"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).DistinctBy(func(s Student) any {
	return len(s.Name)
}, nil)
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {3 A}]
```

#### Except
Except returns the set difference, which means the elements of one collection that don't appear in a second collection.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
	{Id: 1, Name: "AKh"},
	{Id: 2, Name: "Lah"},
	{Id: 3, Name: "A"},
}

second := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
}

enumerable := lingo.AsEnumerable(source).Except(lingo.AsEnumerable(second))
fmt.Println(enumerable.ToSlice())
// Result: [{3 A} {1 AKh} {2 Lah} {3 A}]
```
#### Intersect
Intersect returns the set intersection, which means elements that appear in each of two collections.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
	{Id: 1, Name: "AKh"},
	{Id: 2, Name: "Lah"},
	{Id: 3, Name: "A"},
}

second := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "Ab"},
	{Id: 1, Name: "Ah"},
	{Id: 2, Name: "Lbh"},
	{Id: 3, Name: "B"},
}

enumerable := lingo.AsEnumerable(source).Intersect(lingo.AsEnumerable(second))
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 hnA}]
```

#### Union
Union returns the set union, which means unique elements that appear in either of two collections.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

second := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "Ab"},
}

enumerable := lingo.AsEnumerable(source).Union(lingo.AsEnumerable(second))
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 hnA} {3 A} {3 Ab}]
```

### Sorting data
#### OrderBy
OrderBy sorts values in ascending order.

Example:

Sort Student by Student Id
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 5, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).OrderBy(func(s Student) any {
	return s.Id
}, nil)
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {3 A} {5 hnA}]
```

Sort Student by Student Id with special comparer
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 5, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).OrderBy(func(s Student) any {
	return s.Id
}, func(a1, a2 any) bool {
	return a1.(int) >= a2.(int)
})
fmt.Println(enumerable.ToSlice())
// Result: [{5 hnA} {3 A} {1 Anh}]
```

#### OrderByDescending
OrderByDescending sorts values in descending order.

Example:

```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 5, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).OrderByDescending(func(s Student) any {
	return s.Id
}, nil)
fmt.Println(enumerable.ToSlice())
// Result: [{5 hnA} {3 A} {1 Anh}]
```

#### Reverse
Reverse reverses the order of the elements in a collection.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 5, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).Reverse()
fmt.Println(enumerable.ToSlice())
// Result: [[{3 A} {5 hnA} {1 Anh}]]
```

### Quantifier operations
#### All
All determines whether all the elements in a sequence satisfy a condition.

Example:

```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).All(func(s Student) bool {
	return len(s.Name) > 0
})
fmt.Println(enumerable)
// Result: true
```
#### Any
Any determines whether any elements in a sequence satisfy a condition.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).Any(func(s Student) bool {
	return len(s.Name) > 5
})
fmt.Println(enumerable)
// Result: false
```

#### Contains
Contains determines whether a sequence contains a specified element.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).Contains(Student{Id: 5, Name: "A"}, func(s1, s2 Student) bool {
	return s1.Name == s1.Name
})
fmt.Println(enumerable)
// Result: true
```

### Partitioning data
#### Skip
Skip skips elements up to a specified position in a sequence.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).Skip(2)
fmt.Println(enumerable.ToSlice())
// Result: [{3 A}]
```
#### SkipWhile
SkipWhile skips elements based on a predicate function until an element doesn't satisfy the condition.

Example:

```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).SkipWhile(func(s Student) bool {
	return s.Id < 3
})
fmt.Println(enumerable.ToSlice())
// Result: [{3 A}]
```
#### Take
Take takes elements up to a specified position in a sequence.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).Take(2)
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 hnA}]
```
#### TakeWhile
TakeWhile takes elements based on a predicate function until an element doesn't satisfy the condition.

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source).TakeWhile(func(s Student) bool {
	return s.Id < 3
})
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 hnA}]
```

#### Chunk
Chunk splits the elements of a sequence into chunks of a specified maximum size.

Example:
```go
lingo.Chunk(lingo.Range(1, 11), 5).ToSlice()
// 1-5
// 6-10
// 11
```

### Converting data types
#### ToSlice
ToSlice converts the iterator into a slice

Example:
```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

enumerable := lingo.AsEnumerable(source)
fmt.Println(enumerable.ToSlice())
// Result: [{1 Anh} {2 hnA} {3 A}]
```

#### ToMap
ToMap converts the iterator into a map with specific selector

```go
source := []Student{
	{Id: 1, Name: "Anh"},
	{Id: 2, Name: "hnA"},
	{Id: 3, Name: "A"},
}

m := lingo.AsEnumerable(source).ToMap(func(s Student) any {
	return s.Id
}, func(s Student) any {
	return s.Name
})
fmt.Println(m)
// Result: map[1:Anh 2:hnA 3:A]
```

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

### Grouping data
GroupBy groups elements that share a common attribute.

Example: Take all top 1 students group by class.
```go
type args struct {
	keySelector     lingo.SingleSelector[Student]
	elementSelector lingo.SingleSelector[Student]
	resultSelector  lingo.GroupBySelector[any, any]
	getHash         lingo.GetHashCode[any]
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
		return s.ClassId
	},
	elementSelector: func(s Student) any {
		return s
	},
	resultSelector: func(a1 any, a2 []any) any {
		ss := lingo.AsEnumerableTFromSliceAny[Student](a2).OrderByDescending(func(s Student) any {
			return s.Score
		}, nil).ToSlice()
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
	GroupBy(
		arg.keySelector,
		arg.elementSelector,
		arg.resultSelector,
		arg.getHash,
	)
fmt.Println(enumerable.ToSlice())
// Result: [{1 A 1 7} {3 C 2 10}]
```

## Reference
- [LinQ](https://learn.microsoft.com/en-us/dotnet/csharp/linq/)

## Contributors
Thanks to all the people who already contributed!
