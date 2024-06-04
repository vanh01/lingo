package lingo_test

import (
	"math"
	"reflect"
	"testing"

	lingo "github.com/vanh01/lingo"
	"github.com/vanh01/lingo/definition"
)

func TestMin(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args[T any] struct {
		comparer definition.Comparer[T]
		source   []T
		want     T
	}
	tests := []struct {
		name string
		args args[any]
	}{
		{
			name: "Min",
			args: args[any]{
				source: lingo.SliceTToAny([]int{3, 2, 1, 4, 5}),
				want:   1,
			},
		},
		{
			name: "Min",
			args: args[any]{
				comparer: func(a1, a2 any) bool {
					return a1.(float64) < a2.(float64)
				},
				source: lingo.SliceTToAny([]float64{1.3, 2.2, 3, 4, 5}),
				want:   1.3,
			},
		},
		{
			name: "Min",
			args: args[any]{
				comparer: func(a1, a2 any) bool {
					return a1.(Student).ClassId < a2.(Student).ClassId
				},
				source: lingo.SliceTToAny([]Student{
					{Id: 1, Name: "1", ClassId: 2},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: 2, Name: "12", ClassId: 1},
					{Id: 8, Name: "13", ClassId: 7},
				}),
				want: Student{Id: 2, Name: "12", ClassId: 1},
			},
		},
		{
			name: "Min",
			args: args[any]{
				comparer: func(a1, a2 any) bool {
					return a1.(float64) < a2.(float64)
				},
				source: lingo.SliceTToAny([]float64{1.3, 2.2, 3, 4, 5}),
				want:   1.3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.args.source).Min(tt.args.comparer)
			if got != tt.args.want {
				t.Errorf("Min() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestMinBy(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args[T any] struct {
		selector definition.SingleSelector[T]
		comparer definition.Comparer[any]
		source   []T
		want     T
	}
	tests := []struct {
		name string
		args args[any]
	}{
		{
			name: "MinBy",
			args: args[any]{
				selector: func(a any) any {
					return a.(Student).Id
				},
				source: lingo.SliceTToAny([]Student{
					{Id: 1, Name: "1", ClassId: 2},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: -2, Name: "12", ClassId: 1},
					{Id: 8, Name: "13", ClassId: 7},
				}),
				want: Student{Id: -2, Name: "12", ClassId: 1},
			},
		},
		{
			name: "MinBy",
			args: args[any]{
				selector: func(a any) any {
					return a.(Student).Id
				},
				source: lingo.SliceTToAny([]Student{
					{Id: 1, Name: "1", ClassId: 2},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: 2, Name: "12", ClassId: 1},
					{Id: 8, Name: "13", ClassId: 7},
				}),
				want: Student{Id: 1, Name: "1", ClassId: 2},
			},
		},
		{
			name: "MinBy",
			args: args[any]{
				selector: func(a any) any {
					return a.(Student).Name
				},
				comparer: func(a1, a2 any) bool {
					return len(a1.(string)) < len(a2.(string))
				},
				source: lingo.SliceTToAny([]Student{
					{Id: 1, Name: "1", ClassId: 2},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: -2, Name: "", ClassId: 1},
					{Id: 8, Name: "13", ClassId: 7},
				}),
				want: Student{Id: -2, Name: "", ClassId: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.args.source).MinBy(tt.args.selector, tt.args.comparer)
			if got != tt.args.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.args.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args[T any] struct {
		comparer definition.Comparer[T]
		source   []T
		want     T
	}
	tests := []struct {
		name string
		args args[any]
	}{
		{
			name: "Max",
			args: args[any]{
				source: lingo.SliceTToAny([]int{3, 2, 1, 4, 5}),
				want:   5,
			},
		},
		{
			name: "Max",
			args: args[any]{
				source: lingo.SliceTToAny([]float64{11.3, 2.2, 3, 4, 5}),
				want:   11.3,
			},
		},
		{
			name: "Max",
			args: args[any]{
				comparer: func(a1, a2 any) bool {
					return a1.(Student).ClassId > a2.(Student).ClassId
				},
				source: lingo.SliceTToAny([]Student{
					{Id: 1, Name: "1", ClassId: 9},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: 2, Name: "12", ClassId: 11},
					{Id: 8, Name: "13", ClassId: 7},
				}),
				want: Student{Id: 2, Name: "12", ClassId: 11},
			},
		},
		{
			name: "Max",
			args: args[any]{
				source: lingo.SliceTToAny([]float32{1.3, 8.2, 3, 4, 5}),
				want:   float32(8.2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.args.source).Max(tt.args.comparer)
			if got != tt.args.want {
				t.Errorf("Max() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestMaxBy(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args[T any] struct {
		selector definition.SingleSelector[T]
		comparer definition.Comparer[T]
		source   []T
		want     T
	}
	tests := []struct {
		name string
		args args[any]
	}{
		{
			name: "MaxBy",
			args: args[any]{
				selector: func(a any) any {
					return a.(Student).Id
				},
				source: lingo.SliceTToAny([]Student{
					{Id: 1, Name: "1", ClassId: 2},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: 22, Name: "12", ClassId: 1},
					{Id: 8, Name: "13", ClassId: 7},
				}),
				want: Student{Id: 22, Name: "12", ClassId: 1},
			},
		},
		{
			name: "MaxBy",
			args: args[any]{
				selector: func(a any) any {
					return a.(Student).Id
				},
				source: lingo.SliceTToAny([]Student{
					{Id: 10, Name: "1", ClassId: 2},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: 2, Name: "12", ClassId: 1},
					{Id: 8, Name: "13", ClassId: 7},
				}),
				want: Student{Id: 10, Name: "1", ClassId: 2},
			},
		},
		{
			name: "MaxBy",
			args: args[any]{
				selector: func(a any) any {
					return a.(Student).Name
				},
				comparer: func(a1, a2 any) bool {
					return len(a1.(string)) > len(a2.(string))
				},
				source: lingo.SliceTToAny([]Student{
					{Id: 1, Name: "1", ClassId: 2},
					{Id: 3, Name: "111", ClassId: 6},
					{Id: -2, Name: "", ClassId: 1},
					{Id: 8, Name: "13", ClassId: 7},
				}),
				want: Student{Id: 3, Name: "111", ClassId: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.args.source).MaxBy(tt.args.selector, tt.args.comparer)
			if got != tt.args.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.args.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args[T any] struct {
		selector definition.SingleSelector[T]
		source   []T
		want     float64
	}
	tests := []struct {
		name string
		args args[any]
	}{
		{
			name: "Sum",
			args: args[any]{
				source: lingo.SliceTToAny([]int{3, 2, 1, 4, 5}),
				want:   15,
			},
		},
		{
			name: "Sum",
			args: args[any]{
				source: lingo.SliceTToAny([]float64{11.3, 2.2, 3, 4, 5}),
				want:   25.5,
			},
		},
		{
			name: "Sum",
			args: args[any]{
				selector: func(a any) any {
					return a.(Student).ClassId
				},
				source: lingo.SliceTToAny([]Student{
					{Id: 1, Name: "1", ClassId: 9},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: 2, Name: "12", ClassId: 11},
					{Id: 8, Name: "13", ClassId: 7},
				}),
				want: 14,
			},
		},
		{
			name: "Sum",
			args: args[any]{
				source: lingo.SliceTToAny([]float32{1.3, 8.2, 3, 4, 5}),
				want:   21.5,
			},
		},
		{
			name: "Sum",
			args: args[any]{
				source: lingo.SliceTToAny([]uint32{1, 2, 9, 4, 5}),
				want:   21,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.args.source).Sum(tt.args.selector)
			if reflect.ValueOf(got).Kind() == reflect.Float64 {
				if math.Abs(got.(float64)-tt.args.want) < 1/10000 {
					t.Errorf("%s = %v, want %v", tt.name, got, tt.args.want)
				}
			} else if got == tt.args.want {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.args.want)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args[T any] struct {
		selector definition.SingleSelector[T]
		source   []T
		want     float64
	}
	tests := []struct {
		name string
		args args[any]
	}{
		{
			name: "Average",
			args: args[any]{
				source: lingo.SliceTToAny([]int{3, 2, 1, 4, 5}),
				want:   3,
			},
		},
		{
			name: "Average",
			args: args[any]{
				source: lingo.SliceTToAny([]float64{11.3, 2.2, 3, 4, 5}),
				want:   5.1,
			},
		},
		{
			name: "Average",
			args: args[any]{
				selector: func(a any) any {
					return a.(Student).Id
				},
				source: lingo.SliceTToAny([]Student{
					{Id: 1, Name: "1", ClassId: 9},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: 2, Name: "12", ClassId: 11},
					{Id: 8, Name: "13", ClassId: 7},
				}),
				want: 3.5,
			},
		},
		{
			name: "Average",
			args: args[any]{
				source: lingo.SliceTToAny([]float32{1.3, 8.2, 3, 4, 5}),
				want:   4.3,
			},
		},
		{
			name: "Average",
			args: args[any]{
				source: lingo.SliceTToAny([]uint32{1, 2, 9, 4, 5}),
				want:   4.2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.args.source).Average(tt.args.selector)
			if math.Abs(got-tt.args.want) < 1/10000 {
				t.Errorf("Average() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	type args struct {
		source []int
		want   float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Count",
			args: args{
				source: []int{1, 2, 3, 4, 5},
				want:   5,
			},
		},
		{
			name: "Count",
			args: args{
				source: []int{},
				want:   0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.args.source).Count()
			if got != int64(tt.args.want) {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.args.want)
			}
		})
	}
}

func TestAggregate(t *testing.T) {
	type args struct {
		source     []int
		seed       int
		accmulator definition.Accumulator[any, int]
		selector   definition.SingleSelector[any]
		want       int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Aggregate",
			args: args{
				source: []int{1, 2, 3, 4, 5},
				seed:   0,
				accmulator: func(a1 any, a2 int) any {
					if a2%2 == 0 {
						return a1.(int) + 1
					}
					return a1
				},
				selector: func(a any) any {
					return a.(int) * a.(int)
				},
				want: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.args.source).Aggregate(tt.args.seed, tt.args.accmulator, tt.args.selector)
			if got != tt.args.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.args.want)
			}
		})
	}
}

// ParellelEnumerable

func TestPMinBy(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args[T any] struct {
		selector definition.SingleSelector[T]
		comparer definition.Comparer[any]
		source   []T
		want     T
	}
	tests := []struct {
		name string
		args args[Student]
	}{
		{
			name: "MinBy",
			args: args[Student]{
				selector: func(a Student) any {
					return a.Id
				},
				source: []Student{
					{Id: 1, Name: "1", ClassId: 2},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: -2, Name: "12", ClassId: 1},
					{Id: 8, Name: "13", ClassId: 7},
				},
				want: Student{Id: -2, Name: "12", ClassId: 1},
			},
		},
		{
			name: "MinBy",
			args: args[Student]{
				selector: func(a Student) any {
					return a.Id
				},
				source: []Student{
					{Id: 1, Name: "1", ClassId: 2},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: 2, Name: "12", ClassId: 1},
					{Id: 8, Name: "13", ClassId: 7},
				},
				want: Student{Id: 1, Name: "1", ClassId: 2},
			},
		},
		{
			name: "MinBy",
			args: args[Student]{
				selector: func(a Student) any {
					return a.Name
				},
				comparer: func(a1, a2 any) bool {
					return len(a1.(string)) < len(a2.(string))
				},
				source: []Student{
					{Id: 1, Name: "1", ClassId: 2},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: -2, Name: "", ClassId: 1},
					{Id: 8, Name: "13", ClassId: 7},
				},
				want: Student{Id: -2, Name: "", ClassId: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsParallelEnumerable(tt.args.source).MinBy(tt.args.selector, tt.args.comparer)
			if got != tt.args.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.args.want)
			}
		})
	}
}

func TestPMaxBy(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args[T any] struct {
		selector definition.SingleSelector[T]
		comparer definition.Comparer[any]
		source   []T
		want     T
	}
	tests := []struct {
		name string
		args args[Student]
	}{
		{
			name: "MaxBy",
			args: args[Student]{
				selector: func(a Student) any {
					return a.Id
				},
				source: []Student{
					{Id: 1, Name: "1", ClassId: 2},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: 22, Name: "12", ClassId: 1},
					{Id: 8, Name: "13", ClassId: 7},
				},
				want: Student{Id: 22, Name: "12", ClassId: 1},
			},
		},
		{
			name: "MaxBy",
			args: args[Student]{
				selector: func(a Student) any {
					return a.Id
				},
				source: []Student{
					{Id: 10, Name: "1", ClassId: 2},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: 2, Name: "12", ClassId: 1},
					{Id: 8, Name: "13", ClassId: 7},
				},
				want: Student{Id: 10, Name: "1", ClassId: 2},
			},
		},
		{
			name: "MaxBy",
			args: args[Student]{
				selector: func(a Student) any {
					return a.Name
				},
				comparer: func(a1, a2 any) bool {
					return len(a1.(string)) > len(a2.(string))
				},
				source: []Student{
					{Id: 1, Name: "1", ClassId: 2},
					{Id: 3, Name: "111", ClassId: 6},
					{Id: -2, Name: "", ClassId: 1},
					{Id: 8, Name: "13", ClassId: 7},
				},
				want: Student{Id: 3, Name: "111", ClassId: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.args.source).AsParallel().MaxBy(tt.args.selector, tt.args.comparer)
			if got != tt.args.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.args.want)
			}
		})
	}
}

func TestPSum(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args[T any] struct {
		selector definition.SingleSelector[T]
		source   []T
		want     float64
	}
	tests := []struct {
		name string
		args args[any]
	}{
		{
			name: "Sum",
			args: args[any]{
				source: lingo.SliceTToAny([]int{3, 2, 1, 4, 5}),
				want:   15,
			},
		},
		{
			name: "Sum",
			args: args[any]{
				source: lingo.SliceTToAny([]float64{11.3, 2.2, 3, 4, 5}),
				want:   25.5,
			},
		},
		{
			name: "Sum",
			args: args[any]{
				selector: func(a any) any {
					return a.(Student).ClassId
				},
				source: lingo.SliceTToAny([]Student{
					{Id: 1, Name: "1", ClassId: 9},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: 2, Name: "12", ClassId: 11},
					{Id: 8, Name: "13", ClassId: 7},
				}),
				want: 14,
			},
		},
		{
			name: "Sum",
			args: args[any]{
				source: lingo.SliceTToAny([]float32{1.3, 8.2, 3, 4, 5}),
				want:   21.5,
			},
		},
		{
			name: "Sum",
			args: args[any]{
				source: lingo.SliceTToAny([]uint32{1, 2, 9, 4, 5}),
				want:   21,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsParallelEnumerable(tt.args.source).Sum(tt.args.selector)
			if reflect.ValueOf(got).Kind() == reflect.Float64 {
				if math.Abs(got.(float64)-tt.args.want) < 1/10000 {
					t.Errorf("%s = %v, want %v", tt.name, got, tt.args.want)
				}
			} else if got == tt.args.want {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.args.want)
			}
		})
	}
}

func TestPAverage(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args[T any] struct {
		selector definition.SingleSelector[T]
		source   []T
		want     float64
	}
	tests := []struct {
		name string
		args args[any]
	}{
		{
			name: "Average",
			args: args[any]{
				source: lingo.SliceTToAny([]int{3, 2, 1, 4, 5}),
				want:   3,
			},
		},
		{
			name: "Average",
			args: args[any]{
				source: lingo.SliceTToAny([]float64{11.3, 2.2, 3, 4, 5}),
				want:   5.1,
			},
		},
		{
			name: "Average",
			args: args[any]{
				selector: func(a any) any {
					return a.(Student).Id
				},
				source: lingo.SliceTToAny([]Student{
					{Id: 1, Name: "1", ClassId: 9},
					{Id: 3, Name: "11", ClassId: 6},
					{Id: 2, Name: "12", ClassId: 11},
					{Id: 8, Name: "13", ClassId: 7},
				}),
				want: 3.5,
			},
		},
		{
			name: "Average",
			args: args[any]{
				source: lingo.SliceTToAny([]float32{1.3, 8.2, 3, 4, 5}),
				want:   4.3,
			},
		},
		{
			name: "Average",
			args: args[any]{
				source: lingo.SliceTToAny([]uint32{1, 2, 9, 4, 5}),
				want:   4.2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.args.source).AsParallel().Average(tt.args.selector)
			if math.Abs(got-tt.args.want) < 1/10000 {
				t.Errorf("Average() = %v, want %v", got, tt.args.want)
			}
		})
	}
}
