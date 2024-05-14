package lingo_test

import (
	"fmt"
	"reflect"
	"testing"

	lingo "github.com/vanh01/lingo"
	"github.com/vanh01/lingo/definition"
)

func TestSelect(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type StudentF struct {
		Id    int
		NameF string
	}
	type args struct {
		s definition.SingleSelector[Student]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []StudentF
	}{
		{
			name: "Select",
			source: []Student{
				{Id: 1, Name: "Anh", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Nguyen", Level: 3},
			},
			args: args{s: func(t Student) any {
				return StudentF{
					Id:    t.Id,
					NameF: t.Name + " " + fmt.Sprint(t.Level),
				}
			}},
			want: []StudentF{
				{Id: 1, NameF: "Anh 1"},
				{Id: 2, NameF: "An 2"},
				{Id: 3, NameF: "Nguyen 3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerableTFromAny[StudentF](lingo.AsEnumerable(tt.source).Select(tt.args.s)).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				g := got[i]
				if g.Id != tt.want[i].Id || g.NameF != tt.want[i].NameF {
					t.Errorf("Select() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestSelectMany(t *testing.T) {
	type Class struct {
		Id          int
		StudentName []string
		Total       int
	}
	type args struct {
		s definition.SingleSelector[Class]
	}
	tests := []struct {
		name   string
		source []Class
		args   args
		want   []string
	}{
		{
			name: "SelectMany",
			source: []Class{
				{Id: 1, StudentName: []string{"Anh", "An", "A"}, Total: 1},
				{Id: 2, StudentName: []string{"Nguy", "Nguyen", "Viet"}, Total: 2},
				{Id: 3, StudentName: []string{"Guy", "Vi", "Ngen"}, Total: 3},
			},
			args: args{s: func(t Class) any {
				res := make([]any, len(t.StudentName))
				for i := range t.StudentName {
					res[i] = t.StudentName[i]
				}
				return res
			}},
			want: []string{"Anh", "An", "A", "Nguy", "Nguyen", "Viet", "Guy", "Vi", "Ngen"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).SelectMany(tt.args.s).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("SelectMany() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("SelectMany() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestZip(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type Class struct {
		Name string
	}
	type args struct {
		first          lingo.Enumerable[any]
		resultSelector definition.CombinationSelector[Student, any]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []any
	}{
		{
			name: "Zip",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{first: lingo.AsEnumerableAnyFromSliceT([]int{8, 9, 10})},
			want: []any{
				[]any{Student{Id: 1, Name: "Nam", Level: 1}, 8},
				[]any{Student{Id: 2, Name: "An", Level: 2}, 9},
				[]any{Student{Id: 3, Name: "Anh", Level: 2}, 10},
			},
		},
		{
			name: "Zip",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				first: lingo.AsEnumerableAnyFromT(lingo.AsEnumerable([]int{8, 9, 10})),
				resultSelector: func(s Student, k any) any {
					return []any{s, k.(int) - 1}
				},
			},
			want: []any{
				[]any{Student{Id: 1, Name: "Nam", Level: 1}, 7},
				[]any{Student{Id: 2, Name: "An", Level: 2}, 8},
				[]any{Student{Id: 3, Name: "Anh", Level: 2}, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).Zip(tt.args.first, tt.args.resultSelector).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Zip() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				gotVal := reflect.ValueOf(got[i])
				v1, v2 := gotVal.Index(0).Interface(), gotVal.Index(1).Interface()
				wantVal := reflect.ValueOf(tt.want[i])
				w1, w2 := wantVal.Index(0).Interface(), wantVal.Index(1).Interface()
				if v2 != w2 {
					t.Errorf("Zip() = %v, want %v", got, tt.want)
				}
				va1 := v1.(Student)
				wa1 := w1.(Student)
				if va1.Id != wa1.Id || va1.Name != wa1.Name || va1.Level != wa1.Level {
					t.Errorf("Zip() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

// ParellelEnumerable

func TestPSelect(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type StudentF struct {
		Id    int
		NameF string
	}
	type args struct {
		s definition.SingleSelector[Student]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []StudentF
	}{
		{
			name: "Select",
			source: []Student{
				{Id: 1, Name: "Anh", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Nguyen", Level: 3},
			},
			args: args{s: func(t Student) any {
				return StudentF{
					Id:    t.Id,
					NameF: t.Name + " " + fmt.Sprint(t.Level),
				}
			}},
			want: []StudentF{
				{Id: 1, NameF: "Anh 1"},
				{Id: 2, NameF: "An 2"},
				{Id: 3, NameF: "Nguyen 3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerableTFromAny[StudentF](lingo.AsParallelEnumerable(tt.source).Select(tt.args.s).AsEnumerable()).OrderBy(func(sf StudentF) any {
				return sf.Id
			}).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
			for i := range tt.want {
				g := got[i]
				if g.Id != tt.want[i].Id || g.NameF != tt.want[i].NameF {
					t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
				}
			}
		})
	}
}

func TestPSelectMany(t *testing.T) {
	type Class struct {
		Id          int
		StudentName []string
		Total       int
	}
	type args struct {
		s definition.SingleSelector[Class]
	}
	tests := []struct {
		name   string
		source []Class
		args   args
		want   []string
	}{
		{
			name: "SelectMany",
			source: []Class{
				{Id: 1, StudentName: []string{"Anh", "An", "A"}, Total: 1},
				{Id: 2, StudentName: []string{"Nguy", "Nguyen", "Viet"}, Total: 2},
				{Id: 3, StudentName: []string{"Guy", "Vi", "Ngen"}, Total: 3},
			},
			args: args{s: func(t Class) any {
				res := make([]any, len(t.StudentName))
				for i := range t.StudentName {
					res[i] = t.StudentName[i]
				}
				return res
			}},
			want: []string{"A", "An", "Anh", "Guy", "Ngen", "Nguy", "Nguyen", "Vi", "Viet"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).AsParallel().SelectMany(tt.args.s).AsEnumerable().OrderBy(func(a any) any { return a }).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
				}
			}
		})
	}
}

func TestPZip(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type Class struct {
		Name string
	}
	type args struct {
		first          lingo.ParallelEnumerable[any]
		resultSelector definition.CombinationSelector[Student, any]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []any
	}{
		{
			name: "Zip",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{first: lingo.AsEnumerableAnyFromSliceT([]int{8, 9, 10}).AsParallel()},
			want: []any{
				[]any{Student{Id: 1, Name: "Nam", Level: 1}, 8},
				[]any{Student{Id: 2, Name: "An", Level: 2}, 9},
				[]any{Student{Id: 3, Name: "Anh", Level: 2}, 10},
			},
		},
		{
			name: "Zip",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				first: lingo.AsEnumerableAnyFromT(lingo.AsEnumerable([]int{8, 9, 10})).AsParallel(),
				resultSelector: func(s Student, k any) any {
					return []any{s, k.(int) - 1}
				},
			},
			want: []any{
				[]any{Student{Id: 1, Name: "Nam", Level: 1}, 7},
				[]any{Student{Id: 2, Name: "An", Level: 2}, 8},
				[]any{Student{Id: 3, Name: "Anh", Level: 2}, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsParallelEnumerable(tt.source).
				Zip(tt.args.first, tt.args.resultSelector).
				AsEnumerable().
				OrderBy(func(a any) any {
					return a.([]any)[1]
				}).
				ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
			for i := range tt.want {
				gotVal := reflect.ValueOf(got[i])
				v1, v2 := gotVal.Index(0).Interface(), gotVal.Index(1).Interface()
				wantVal := reflect.ValueOf(tt.want[i])
				w1, w2 := wantVal.Index(0).Interface(), wantVal.Index(1).Interface()
				if v2 != w2 {
					t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
				}
				va1 := v1.(Student)
				wa1 := w1.(Student)
				if va1.Id != wa1.Id || va1.Name != wa1.Name || va1.Level != wa1.Level {
					t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
				}
			}
		})
	}
}
