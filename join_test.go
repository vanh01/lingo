package lingo_test

import (
	"testing"

	lingo "github.com/vanh01/lingo"
	"github.com/vanh01/lingo/definition"
)

func TestJoin(t *testing.T) {
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
	type StudentClass struct {
		Id        int
		Name      string
		ClassId   int
		ClassName string
		Total     int
	}
	type args struct {
		inner            []any
		innerKeySelector definition.SingleSelector[any]
		outerKeySelector definition.SingleSelector[Student]
		resultSelector   definition.CombinationSelector[Student, any]
		comparer         definition.Comparer[any]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []StudentClass
	}{
		{
			name: "Where",
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
				inner: []any{
					Class{Id: 1, Name: "Class 1", Total: 20},
					Class{Id: 2, Name: "Class 2", Total: 10},
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
			},
			want: []StudentClass{
				{Id: 1, Name: "A", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 2, Name: "B", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 3, Name: "C", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 4, Name: "D", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 5, Name: "E", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 6, Name: "F", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 7, Name: "G", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 8, Name: "H", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 9, Name: "J", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 10, Name: "K", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 12, Name: "Q", ClassId: 2, ClassName: "Class 2", Total: 10},
			},
		},
		{
			name: "Where",
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
				inner: []any{
					Class{Id: 1, Name: "Class 1", Total: 20},
					Class{Id: 2, Name: "Class 2", Total: 10},
				},
				outerKeySelector: func(s Student) any { return s },
				innerKeySelector: func(c any) any { return c },
				resultSelector: func(s Student, c any) any {
					return StudentClass{
						Id:        s.Id,
						Name:      s.Name,
						ClassId:   s.ClassId,
						ClassName: c.(Class).Name,
						Total:     c.(Class).Total,
					}
				},
				comparer: func(a1, a2 any) bool {
					return a1.(Student).ClassId == a2.(Class).Id
				},
			},
			want: []StudentClass{
				{Id: 1, Name: "A", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 2, Name: "B", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 3, Name: "C", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 4, Name: "D", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 5, Name: "E", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 6, Name: "F", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 7, Name: "G", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 8, Name: "H", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 9, Name: "J", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 10, Name: "K", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 12, Name: "Q", ClassId: 2, ClassName: "Class 2", Total: 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).
				Join(
					lingo.AsEnumerable(tt.args.inner),
					tt.args.outerKeySelector,
					tt.args.innerKeySelector,
					tt.args.resultSelector,
					tt.args.comparer).
				ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Join() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestPJoin(t *testing.T) {
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
	type StudentClass struct {
		Id        int
		Name      string
		ClassId   int
		ClassName string
		Total     int
	}
	type args struct {
		inner            []any
		innerKeySelector definition.SingleSelector[any]
		outerKeySelector definition.SingleSelector[Student]
		resultSelector   definition.CombinationSelector[Student, any]
		comparer         definition.Comparer[any]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []StudentClass
	}{
		{
			name: "Where",
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
				inner: []any{
					Class{Id: 1, Name: "Class 1", Total: 20},
					Class{Id: 2, Name: "Class 2", Total: 10},
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
			},
			want: []StudentClass{
				{Id: 1, Name: "A", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 2, Name: "B", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 3, Name: "C", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 4, Name: "D", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 5, Name: "E", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 6, Name: "F", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 7, Name: "G", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 8, Name: "H", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 9, Name: "J", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 10, Name: "K", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 12, Name: "Q", ClassId: 2, ClassName: "Class 2", Total: 10},
			},
		},
		{
			name: "Where",
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
				inner: []any{
					Class{Id: 1, Name: "Class 1", Total: 20},
					Class{Id: 2, Name: "Class 2", Total: 10},
				},
				outerKeySelector: func(s Student) any { return s },
				innerKeySelector: func(c any) any { return c },
				resultSelector: func(s Student, c any) any {
					return StudentClass{
						Id:        s.Id,
						Name:      s.Name,
						ClassId:   s.ClassId,
						ClassName: c.(Class).Name,
						Total:     c.(Class).Total,
					}
				},
				comparer: func(a1, a2 any) bool {
					return a1.(Student).ClassId == a2.(Class).Id
				},
			},
			want: []StudentClass{
				{Id: 1, Name: "A", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 2, Name: "B", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 3, Name: "C", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 4, Name: "D", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 5, Name: "E", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 6, Name: "F", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 7, Name: "G", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 8, Name: "H", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 9, Name: "J", ClassId: 1, ClassName: "Class 1", Total: 20},
				{Id: 10, Name: "K", ClassId: 2, ClassName: "Class 2", Total: 10},
				{Id: 12, Name: "Q", ClassId: 2, ClassName: "Class 2", Total: 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsParallelEnumerable(tt.source).AsOrdered().
				Join(
					lingo.AsParallelEnumerable(tt.args.inner).AsOrdered(),
					tt.args.outerKeySelector,
					tt.args.innerKeySelector,
					tt.args.resultSelector,
					tt.args.comparer).
				ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Join() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
