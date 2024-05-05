package lingo_test

import (
	"strings"
	"testing"

	lingo "github.com/vanh01/lingo"
)

func TestSkip(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
		number int
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "Skip",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				number: 2,
			},
			want: []Student{
				{Id: 3, Name: "Anh", Level: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).Skip(tt.args.number).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Skip() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Skip() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestSkipWhile(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
		p lingo.Predicate[Student]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "SkipWhile",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				p: func(s Student) bool {
					return s.Id < 3
				},
			},
			want: []Student{
				{Id: 3, Name: "Anh", Level: 2},
			},
		},
		{
			name: "SkipWhile",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				p: func(s Student) bool {
					return strings.Contains(s.Name, "N")
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
			got := lingo.AsEnumerable(tt.source).SkipWhile(tt.args.p).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("SkipWhile() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("SkipWhile() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestTake(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
		number int
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "Take",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				number: 2,
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).Take(tt.args.number).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Take() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Take() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestTakeWhile(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
		p lingo.Predicate[Student]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   []Student
	}{
		{
			name: "SkipWhile",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh", Level: 2},
			},
			args: args{
				p: func(s Student) bool {
					return s.Id < 3
				},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
			},
		},
		{
			name: "SkipWhile",
			source: []Student{
				{Id: 1, Name: "Nam", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Anh N", Level: 2},
			},
			args: args{
				p: func(s Student) bool {
					return strings.Contains(s.Name, "N")
				},
			},
			want: []Student{
				{Id: 1, Name: "Nam", Level: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).TakeWhile(tt.args.p).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("SkipWhile() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("SkipWhile() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestChunk(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name   string
		source lingo.Enumerable[int]
		args   args
		want   [][]int
	}{
		{
			name:   "Chunk",
			source: lingo.Range(1, 2),
			args:   args{size: 3},
			want: [][]int{
				{1, 2},
			},
		},
		{
			name:   "Chunk",
			source: lingo.Range(1, 3),
			args:   args{size: 3},
			want: [][]int{
				{1, 2, 3},
			},
		},
		{
			name:   "Chunk",
			source: lingo.Range(1, 4),
			args:   args{size: 3},
			want: [][]int{
				{1, 2, 3},
				{4},
			},
		},
		{
			name:   "Chunk",
			source: lingo.Range(1, 6),
			args:   args{size: 3},
			want: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.Chunk(tt.source, tt.args.size).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
			for i := range tt.want {
				if len(got[i]) != len(tt.want[i]) {

					for j := range tt.want[i] {
						if got[i][j] != tt.want[i][j] {
							t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
						}
					}
				}
			}
		})
	}
}
