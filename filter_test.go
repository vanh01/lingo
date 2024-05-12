package lingo_test

import (
	"testing"

	lingo "github.com/vanh01/lingo"
)

func TestWhere(t *testing.T) {
	type args struct {
		p lingo.Predicate[int]
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   []int
	}{
		{
			name:   "Where",
			source: []int{1, 2, 3},
			args:   args{p: func(v int) bool { return v > 2 }},
			want:   []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).Where(tt.args.p).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Where() = %v, want %v", got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Where() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

// ParallelEnumerable

func TestPWhere(t *testing.T) {
	type args struct {
		p lingo.Predicate[int]
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   []int
	}{
		{
			name:   "Where",
			source: []int{1, 2, 3},
			args:   args{p: func(v int) bool { return v > 2 }},
			want:   []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsParallelEnumerable(tt.source).Where(tt.args.p).AsEnumerable().OrderBy(func(i int) any { return i }).ToSlice()
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
