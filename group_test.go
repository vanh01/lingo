package lingo_test

import (
	"testing"

	lingo "github.com/vanh01/lingo"
	"github.com/vanh01/lingo/definition"
)

func TestGroupBy(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args struct {
		keySelector     definition.SingleSelector[Student]
		elementSelector definition.SingleSelector[Student]
		resultSelector  definition.GroupBySelector[any, any]
		getHash         definition.GetHashCode[any]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "GroupBy",
			source: []Student{
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
			},
			args: args{
				keySelector: func(s Student) any {
					return s.ClassId
				},
				elementSelector: func(s Student) any {
					return s
				},
				resultSelector: func(a1 any, a2 []any) any {
					ss := lingo.AsEnumerableTFromSliceAny[Student](a2).OrderBy(func(s Student) any {
						return s.Id
					}, nil).ToSlice()
					s := ss[0]
					return Student{
						Id:      s.Id,
						Name:    s.Name,
						ClassId: a1.(int),
					}
				},
			},
			want: []Student{
				{Id: 1, Name: "A", ClassId: 1},
				{Id: 2, Name: "B", ClassId: 2},
			},
		},
		{
			name: "GroupBy",
			source: []Student{
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
			},
			args: args{
				keySelector: func(s Student) any {
					return s.ClassId
				},
				elementSelector: func(s Student) any {
					return s
				},
				resultSelector: func(a1 any, a2 []any) any {
					ss := lingo.AsEnumerableTFromSliceAny[Student](a2).OrderBy(func(s Student) any {
						return s.Id
					}, nil).ToSlice()
					s := ss[0]
					return Student{
						Id:      s.Id,
						Name:    s.Name,
						ClassId: s.ClassId,
					}
				},
				getHash: func(a any) any {
					return 1
				},
			},
			want: []Student{
				{Id: 1, Name: "A", ClassId: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).
				GroupBy(
					tt.args.keySelector,
					tt.args.elementSelector,
					tt.args.resultSelector,
					tt.args.getHash,
				).
				OrderBy(func(a any) any { return a.(Student).Id }, nil).
				ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("GroupBy() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

// ParallelEnumerable

func TestPGroupBy(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args struct {
		keySelector     definition.SingleSelector[Student]
		elementSelector definition.SingleSelector[Student]
		resultSelector  definition.GroupBySelector[any, any]
		getHash         definition.GetHashCode[any]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "GroupBy",
			source: []Student{
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
			},
			args: args{
				keySelector: func(s Student) any {
					return s.ClassId
				},
				elementSelector: func(s Student) any {
					return s
				},
				resultSelector: func(a1 any, a2 []any) any {
					ss := lingo.AsEnumerableTFromSliceAny[Student](a2).OrderBy(func(s Student) any {
						return s.Id
					}, nil).ToSlice()
					s := ss[0]
					return Student{
						Id:      s.Id,
						Name:    s.Name,
						ClassId: a1.(int),
					}
				},
			},
			want: []Student{
				{Id: 1, Name: "A", ClassId: 1},
				{Id: 2, Name: "B", ClassId: 2},
			},
		},
		{
			name: "GroupBy",
			source: []Student{
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
			},
			args: args{
				keySelector: func(s Student) any {
					return s.ClassId
				},
				elementSelector: func(s Student) any {
					return s
				},
				resultSelector: func(a1 any, a2 []any) any {
					ss := lingo.AsEnumerableTFromSliceAny[Student](a2).OrderBy(func(s Student) any {
						return s.Id
					}, nil).ToSlice()
					s := ss[0]
					return Student{
						Id:      s.Id,
						Name:    s.Name,
						ClassId: s.ClassId,
					}
				},
				getHash: func(a any) any {
					return 1
				},
			},
			want: []Student{
				{Id: 1, Name: "A", ClassId: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).
				AsParallel().
				GroupBy(
					tt.args.keySelector,
					tt.args.elementSelector,
					tt.args.resultSelector,
					tt.args.getHash,
				).
				AsEnumerable().
				OrderBy(func(a any) any { return a.(Student).Id }, nil).
				ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("GroupBy() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
