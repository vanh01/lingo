package lingo_test

import (
	"testing"

	lingo "github.com/vanh01/lingo"
)

func TestDistinct(t *testing.T) {
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
			name: "Distinct",
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).Distinct().ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Distinct() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Distinct() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestDistinctBy(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
		keySelector lingo.SingleSelector[Student]
		comparer    lingo.Comparer[any]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "DistinctBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{keySelector: func(t Student) any {
				return t.Level
			}},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
			},
		},
		{
			name: "DistinctBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{keySelector: func(t Student) any {
				return len(t.Name)
			}},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
			},
		},
		{
			name: "DistinctBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 3},
				{Id: 2, Name: "An", Level: 4},
				{Id: 3, Name: "Anh", Level: 4},
			},
			args: args{
				keySelector: func(t Student) any {
					return t
				},
				comparer: func(t1, t2 any) bool {
					return t1.(Student).Level == t2.(Student).Level
				},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 3},
				{Id: 2, Name: "An", Level: 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).DistinctBy(tt.args.keySelector, tt.args.comparer).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("DistinctBy() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("DistinctBy() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestExcept(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
		second []Student
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "Except",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{second: []Student{
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			}},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
		},
		{
			name: "Except",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{second: []Student{
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			}},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).Except(lingo.AsEnumerable(tt.args.second)).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Except() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Except() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
		second []Student
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "Intersect",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{second: []Student{
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 4, Name: "An", Level: 2},
				{Id: 5, Name: "Anh", Level: 2},
			}},
			want: []Student{
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
		},
		{
			name: "Intersect",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{second: []Student{
				{Id: 3, Name: "Anh", Level: 2},
			}},
			want: []Student{
				{Id: 3, Name: "Anh", Level: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).Intersect(lingo.AsEnumerable(tt.args.second)).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Intersect() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestUnion(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
		second []Student
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "Union",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{second: []Student{
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 4, Name: "An", Level: 2},
				{Id: 5, Name: "Anh", Level: 2},
			}},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 4, Name: "An", Level: 2},
				{Id: 5, Name: "Anh", Level: 2},
			},
		},
		{
			name: "Union",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{second: []Student{
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			}},
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
			got := lingo.AsEnumerable(tt.source).Union(lingo.AsEnumerable(tt.args.second)).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Union() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
