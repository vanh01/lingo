package lingo_test

import (
	"testing"

	lingo "github.com/vanh01/lingo"
)

func TestConcat(t *testing.T) {
	type args struct {
		source []int
		second []int
		want   []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Concat",
			args: args{
				source: []int{1, 2, 3},
				second: []int{4, 5, 6},
				want:   []int{1, 2, 3, 4, 5, 6},
			},
		},
		{
			name: "Concat",
			args: args{
				source: []int{},
				second: []int{1, 2, 3},
				want:   []int{1, 2, 3},
			},
		},
		{
			name: "Concat",
			args: args{
				source: []int{1, 2, 3},
				second: []int{},
				want:   []int{1, 2, 3},
			},
		},
		{
			name: "Concat",
			args: args{
				source: []int{},
				second: []int{},
				want:   []int{},
			},
		},
		{
			name: "Concat",
			args: args{
				want: []int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.args.source).Concat(lingo.AsEnumerable(tt.args.second)).ToSlice()
			if len(got) != len(tt.args.want) {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.args.want)
			}
			for i := range tt.args.want {
				if got[i] != tt.args.want[i] {
					t.Errorf("%s() = %v, want %v", tt.name, got[i], tt.args.want[i])
				}
			}
		})
	}
}

func TestEmpty(t *testing.T) {
	type args struct {
		second []int
		want   []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Empty",
			args: args{
				second: []int{1, 2, 3},
				want:   []int{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.Empty[int]().Concat(lingo.AsEnumerable(tt.args.second)).ToSlice()
			if len(got) != len(tt.args.want) {
				t.Errorf("%s() = %v", tt.name, got)
			}
			for i := range tt.args.want {
				if got[i] != tt.args.want[i] {
					t.Errorf("%s() = %v, want %v", tt.name, got[i], tt.args.want[i])
				}
			}
		})
	}
}

func TestRange(t *testing.T) {
	type args struct {
		start int
		end   int
		want  []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Range",
			args: args{
				start: 1,
				end:   3,
				want:  []int{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.Range(tt.args.start, tt.args.end).ToSlice()
			if len(got) != len(tt.args.want) {
				t.Errorf("%s() = %v", tt.name, got)
			}
			for i := range tt.args.want {
				if got[i] != tt.args.want[i] {
					t.Errorf("%s() = %v, want %v", tt.name, got[i], tt.args.want[i])
				}
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	type args struct {
		value int
		times int
		want  []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Range",
			args: args{
				value: 1,
				times: 3,
				want:  []int{1, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.Repeat(tt.args.value, tt.args.times).ToSlice()
			if len(got) != len(tt.args.want) {
				t.Errorf("%s() = %v", tt.name, got)
			}
			for i := range tt.args.want {
				if got[i] != tt.args.want[i] {
					t.Errorf("%s() = %v, want %v", tt.name, got[i], tt.args.want[i])
				}
			}
		})
	}
}

func TestGetIter(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   []int
	}{
		{
			name:   "GetIter",
			source: []int{1, 2, 3, 4, 5, 6},
			want:   []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := 0
			for value := range lingo.AsEnumerable(tt.source).GetIter() {
				if value != tt.want[i] {
					t.Errorf("%s() = %v, want %v", tt.name, value, tt.want[i])
				}
				i++
			}
		})
	}
}

func TestAsEnumerableFromChannel(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   []int
	}{
		{
			name:   "AsEnumerableFromChannel",
			source: []int{1, 2, 3, 4, 5, 6},
			want:   []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := 0
			for _, value := range lingo.AsEnumerableFromChannel(lingo.AsEnumerable(tt.source).GetIter()).ToSlice() {
				if value != tt.want[i] {
					t.Errorf("%s() = %v, want %v", tt.name, value, tt.want[i])
				}
				i++
			}
		})
	}
}
