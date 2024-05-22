package lingo_test

import (
	"testing"

	lingo "github.com/vanh01/lingo"
	"github.com/vanh01/lingo/definition"
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
		keySelector definition.SingleSelector[Student]
		comparer    definition.Comparer[any]
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

func TestExceptBy(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
		second      []any
		keySelector definition.SingleSelector[Student]
		comparer    definition.Comparer[any]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "ExceptBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 3, Name: "Anh1", Level: 2},
				{Id: 3, Name: "Anh3", Level: 2},
			},
			args: args{
				second: []any{
					"An",
					"Anh",
				},
				keySelector: func(s Student) any {
					return s.Name
				},
				comparer: func(a1, a2 any) bool {
					return len(a1.(string)) == len(a2.(string))
				},
			},
			want: []Student{
				{Id: 3, Name: "Anh1", Level: 2},
			},
		},
		{
			name: "ExceptBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 1, Name: "Nam", Level: 1},
			},
			args: args{
				second: []any{
					Student{Id: 2, Name: "An", Level: 2},
					Student{Id: 2, Name: "An", Level: 2},
					Student{Id: 3, Name: "Anh", Level: 2},
				},
				keySelector: func(s Student) any {
					return s
				},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).ExceptBy(lingo.AsEnumerable(tt.args.second), tt.args.keySelector, tt.args.comparer).ToSlice()
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

func TestIntersectBy(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
		second      []any
		keySelector definition.SingleSelector[Student]
		comparer    definition.Comparer[any]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "IntersectBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				second: []any{
					"An",
					"Anh",
					"Anh",
					"An",
					"Anh",
				},
				keySelector: func(s Student) any {
					return s.Name
				},
				comparer: func(a1, a2 any) bool {
					return len(a1.(string)) == len(a2.(string))
				},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
			},
		},
		{
			name: "IntersectBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				second: []any{
					"An",
					"Anh",
					"Anh",
					"An",
					"Anh",
				},
				keySelector: func(s Student) any {
					return s.Name
				},
			},
			want: []Student{
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).IntersectBy(
				lingo.AsEnumerable(tt.args.second),
				tt.args.keySelector,
				tt.args.comparer,
			).ToSlice()
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

func TestUnionBy(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
		second      []Student
		keySelector definition.SingleSelector[Student]
		comparer    definition.Comparer[any]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "UnionBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				second: []Student{
					{Id: 2, Name: "An", Level: 2},
					{Id: 3, Name: "Anh", Level: 2},
					{Id: 4, Name: "An", Level: 2},
					{Id: 5, Name: "Anh", Level: 2},
					{Id: 5, Name: "1", Level: 2},
				},
				keySelector: func(s Student) any {
					return s.Name
				},
				comparer: func(a1, a2 any) bool {
					return len(a1.(string)) == len(a2.(string))
				},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 5, Name: "1", Level: 2},
			},
		},
		{
			name: "UnionBy",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				second: []Student{
					{Id: 2, Name: "An", Level: 2},
					{Id: 3, Name: "Anh", Level: 2},
					{Id: 4, Name: "An", Level: 2},
					{Id: 5, Name: "Anh", Level: 2},
					{Id: 5, Name: "AnhHoa", Level: 2},
				},
				keySelector: func(s Student) any {
					return s.Name
				},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
				{Id: 5, Name: "AnhHoa", Level: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).UnionBy(
				lingo.AsEnumerable(tt.args.second),
				tt.args.keySelector,
				tt.args.comparer,
			).ToSlice()
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

func TestPDistinct(t *testing.T) {
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
			got := lingo.AsEnumerable(tt.source).AsParallel().AsOrdered().Distinct().ToSlice()
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

func TestPExcept(t *testing.T) {
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).AsParallel().AsOrdered().Except(lingo.AsEnumerable(tt.args.second).AsParallel().AsOrdered()).ToSlice()
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

func TestPIntersect(t *testing.T) {
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
				{Id: 3, Name: "Anh", Level: 2},
			}},
			want: []Student{
				{Id: 3, Name: "Anh", Level: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).AsParallel().AsOrdered().Intersect(lingo.AsParallelEnumerable(tt.args.second).AsOrdered()).ToSlice()
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

func TestPUnion(t *testing.T) {
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsParallelEnumerable(tt.source).AsOrdered().Concat(lingo.AsParallelEnumerable(tt.source)).Union(lingo.AsParallelEnumerable(tt.args.second).AsOrdered()).ToSlice()
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
