package lingo_test

import (
	"testing"

	lingo "github.com/vanh01/lingo"
	"github.com/vanh01/lingo/definition"
)

func TestOrderBy(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level uint
	}
	type args struct {
		selector definition.SingleSelector[Student]
		comparer definition.Comparer[any]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "OrderBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				selector: func(s Student) any {
					return s.Id
				},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
		},
		{
			name: "OrderBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				selector: func(s Student) any {
					return s.Name
				},
			},
			want: []Student{
				{Id: 2, Name: "An", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 1, Name: "Nam", Level: 1},
			},
		},
		{
			name: "OrderBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 5},
				{Id: 2, Name: "An", Level: 4},
				{Id: 3, Name: "Anh", Level: 1},
			},
			args: args{
				selector: func(s Student) any {
					return s.Level
				},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 3, Name: "Anh", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 2, Name: "An", Level: 4},
				{Id: 3, Name: "Anh", Level: 5},
			},
		},
		{
			name: "OrderBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				selector: func(s Student) any {
					return s
				},
				comparer: func(a1, a2 any) bool {
					return a1.(Student).Id < a2.(Student).Id
				},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).OrderBy(tt.args.selector, tt.args.comparer).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("OrderBy() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("OrderBy() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestOrderByDescending(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
		selector definition.SingleSelector[Student]
		comparer definition.Comparer[any]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "OrderByDescending",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				selector: func(s Student) any {
					return s.Id
				},
			},
			want: []Student{
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 1, Name: "Nam", Level: 1},
			},
		},
		{
			name: "OrderByDescending",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				selector: func(s Student) any {
					return s.Id
				},
				comparer: func(a1, a2 any) bool {
					return a1.(int) > a2.(int)
				},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
		},
		{
			name:   "OrderByDescending",
			source: []Student{},
			args: args{
				selector: func(s Student) any {
					return s.Id
				},
				comparer: func(a1, a2 any) bool {
					return a1.(int) > a2.(int)
				},
			},
			want: []Student{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).OrderByDescending(tt.args.selector, tt.args.comparer).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("OrderByDescending() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("OrderByDescending() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "Reverse",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			want: []Student{
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 1, Name: "Nam", Level: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).Reverse().ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Reverse() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestPOrderBy(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level uint
	}
	type args struct {
		selector definition.SingleSelector[Student]
		comparer definition.Comparer[any]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "OrderBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				selector: func(s Student) any {
					return s.Id
				},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
		},
		{
			name: "OrderBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				selector: func(s Student) any {
					return s.Name
				},
			},
			want: []Student{
				{Id: 2, Name: "An", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 1, Name: "Nam", Level: 1},
			},
		},
		{
			name: "OrderBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 5},
				{Id: 2, Name: "An", Level: 4},
				{Id: 3, Name: "Anh", Level: 1},
			},
			args: args{
				selector: func(s Student) any {
					return s.Level
				},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 3, Name: "Anh", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 2, Name: "An", Level: 4},
				{Id: 3, Name: "Anh", Level: 5},
			},
		},
		{
			name: "OrderBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				selector: func(s Student) any {
					return s
				},
				comparer: func(a1, a2 any) bool {
					return a1.(Student).Id < a2.(Student).Id
				},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).AsParallel().OrderBy(tt.args.selector, tt.args.comparer).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("OrderBy() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("OrderBy() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestPOrderByDescending(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
		selector definition.SingleSelector[Student]
		comparer definition.Comparer[any]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "OrderByDescending",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				selector: func(s Student) any {
					return s.Id
				},
			},
			want: []Student{
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 1, Name: "Nam", Level: 1},
			},
		},
		{
			name: "OrderByDescending",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				selector: func(s Student) any {
					return s.Id
				},
				comparer: func(a1, a2 any) bool {
					return a1.(int) > a2.(int)
				},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
		},
		{
			name:   "OrderByDescending",
			source: []Student{},
			args: args{
				selector: func(s Student) any {
					return s.Id
				},
				comparer: func(a1, a2 any) bool {
					return a1.(int) > a2.(int)
				},
			},
			want: []Student{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsParallelEnumerable(tt.source).OrderByDescending(tt.args.selector, tt.args.comparer).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("OrderByDescending() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("OrderByDescending() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestPReverseOrdered(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "Reverse",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			want: []Student{
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 1, Name: "Nam", Level: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsParallelEnumerable(tt.source).AsOrdered().Reverse().ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Reverse() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestPReverseUnordered(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "Reverse",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsParallelEnumerable(tt.source).Reverse().ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Reverse() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
