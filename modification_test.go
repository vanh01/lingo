package lingo_test

import (
	"testing"

	lingo "github.com/vanh01/lingo"
)

func TestAppend(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   []int
	}{
		{
			name:   "Append",
			source: []int{1, 2, 3},
			args:   args{num: 1},
			want:   []int{1, 2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).Append(tt.args.num).ToSlice()
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

func TestAppendRange(t *testing.T) {
	type args struct {
		num []int
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   []int
	}{
		{
			name:   "AppendRange",
			source: []int{1, 2, 3},
			args:   args{num: []int{2, 3, 4}},
			want:   []int{1, 2, 3, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).AppendRange(lingo.AsEnumerable(tt.args.num)).ToSlice()
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

func TestPrepend(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   []int
	}{
		{
			name:   "Prepend",
			source: []int{1, 2, 3},
			args:   args{num: 1},
			want:   []int{1, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).Prepend(tt.args.num).ToSlice()
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

func TestPrependRange(t *testing.T) {
	type args struct {
		num []int
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   []int
	}{
		{
			name:   "PrependRange",
			source: []int{1, 2, 3},
			args:   args{num: []int{2, 3, 4}},
			want:   []int{2, 3, 4, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).PrependRange(lingo.AsEnumerable(tt.args.num)).ToSlice()
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

func TestClear(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   []int
	}{
		{
			name:   "Clear",
			source: []int{1, 2, 3},
			want:   []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).Clear().ToSlice()
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

func TestInsert(t *testing.T) {
	type args struct {
		index int
		num   int
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   []int
	}{
		{
			name:   "Insert",
			source: []int{1, 2, 3},
			args: args{
				index: 1,
				num:   99,
			},
			want: []int{1, 99, 2, 3},
		},
		{
			name:   "Insert",
			source: []int{1, 2, 3},
			args: args{
				index: 0,
				num:   99,
			},
			want: []int{99, 1, 2, 3},
		},
		{
			name:   "Insert",
			source: []int{1, 2, 3},
			args: args{
				index: 3,
				num:   99,
			},
			want: []int{1, 2, 3, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).Insert(tt.args.index, tt.args.num).ToSlice()
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

func TestRemove(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   []int
	}{
		{
			name:   "Remove",
			source: []int{1, 2, 3},
			args: args{
				num: 1,
			},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).Remove(tt.args.num, nil).ToSlice()
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

func TestRemoveAt(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   []int
	}{
		{
			name:   "RemoveAt",
			source: []int{1, 2, 3},
			args: args{
				index: 1,
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).RemoveAt(tt.args.index).ToSlice()
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

func TestRemoveRange(t *testing.T) {
	type args struct {
		index int
		count int
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   []int
	}{
		{
			name:   "RemoveAt",
			source: []int{1, 2, 3},
			args: args{
				index: 1,
				count: 2,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).RemoveRange(tt.args.index, tt.args.count).ToSlice()
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
