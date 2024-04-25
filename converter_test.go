package lingo_test

import (
	"testing"

	lingo "github.com/vanh01/lingo"
)

func TestToSlice(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   []int
	}{
		{
			name:   "ToSlice",
			source: []int{1, 2, 3},
			want:   []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
			for k, v := range tt.want {
				if got[k] != v {
					t.Errorf("ToSlice() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestToMap(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   map[int]int
	}{
		{
			name:   "ToMap",
			source: []int{1, 2, 3},
			want: map[int]int{
				1: 1,
				2: 2,
				3: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).ToMap(func(i int) any { return i }, func(i int) any { return i })
			if len(got) != len(tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
			for k, v := range tt.want {
				if got[k] != v {
					t.Errorf("ToMap() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestSliceAnyToT(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name   string
		source []any
		args   args
		want   []int
	}{
		{
			name:   "SliceAnyToT",
			source: []any{1, 2, 3},
			want:   []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.SliceAnyToT[int](tt.source)
			if len(got) != len(tt.want) {
				t.Errorf("SliceAnyToT() = %v, want %v", got, tt.want)
			}
			for k, v := range tt.want {
				if got[k] != v {
					t.Errorf("SliceAnyToT() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestSliceTToAny(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   []any
	}{
		{
			name:   "SliceTToAny",
			source: []int{1, 2, 3},
			want:   []any{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.SliceTToAny[int](tt.source)
			if len(got) != len(tt.want) {
				t.Errorf("SliceTToAny() = %v, want %v", got, tt.want)
			}
			for k, v := range tt.want {
				if got[k] != v {
					t.Errorf("SliceTToAny() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
