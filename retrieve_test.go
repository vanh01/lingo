package lingo_test

import (
	"testing"

	lingo "github.com/vanh01/lingo"
)

func TestFirstOrNil(t *testing.T) {
	type args struct {
		p lingo.Predicate[int]
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   int
	}{
		{
			name:   "FirstOrNil",
			source: []int{1, 2, 3},
			want:   1,
		},
		{
			name:   "FirstOrNil",
			source: []int{1, 2, 3},
			args:   args{p: func(v int) bool { return v > 2 }},
			want:   3,
		},
		{
			name:   "FirstOrNil",
			source: []int{1, 2, 3},
			args:   args{p: func(v int) bool { return v > 5 }},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).FirstOrNil(tt.args.p)
			if got != tt.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestFirstOrDefault(t *testing.T) {
	type args struct {
		p            lingo.Predicate[int]
		defaultValue int
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   int
	}{
		{
			name:   "FirstOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -99,
			},
			want: 1,
		},
		{
			name:   "FirstOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -999,
				p:            func(v int) bool { return v > 2 },
			},
			want: 3,
		},
		{
			name:   "FirstOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -9,
				p:            func(v int) bool { return v > 10 },
			},
			want: -9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).FirstOrDefault(tt.args.defaultValue, tt.args.p)
			if got != tt.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestLastOrNil(t *testing.T) {
	type args struct {
		p lingo.Predicate[int]
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   int
	}{
		{
			name:   "LastOrNil",
			source: []int{1, 2, 3},
			want:   3,
		},
		{
			name:   "LastOrNil",
			source: []int{1, 2, 3},
			args:   args{p: func(v int) bool { return v > 2 }},
			want:   3,
		},
		{
			name:   "LastOrNil",
			source: []int{1, 2, 3},
			args:   args{p: func(v int) bool { return v > 5 }},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).LastOrNil(tt.args.p)
			if got != tt.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestLastOrDefault(t *testing.T) {
	type args struct {
		p            lingo.Predicate[int]
		defaultValue int
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   int
	}{
		{
			name:   "LastOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -99,
			},
			want: 3,
		},
		{
			name:   "LastOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -999,
				p:            func(v int) bool { return v > 2 },
			},
			want: 3,
		},
		{
			name:   "LastOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -9,
				p:            func(v int) bool { return v > 10 },
			},
			want: -9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).LastOrDefault(tt.args.defaultValue, tt.args.p)
			if got != tt.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestElementAtOrNil(t *testing.T) {
	type args struct {
		index int64
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   int
	}{
		{
			name:   "ElementAtOrNil",
			source: []int{1, 2, 3},
			args:   args{index: 1},
			want:   2,
		},
		{
			name:   "ElementAtOrNil",
			source: []int{1, 2, 3},
			args:   args{index: 4},
			want:   0,
		},
		{
			name:   "ElementAtOrNil",
			source: []int{1, 2, 3},
			args:   args{index: 2},
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).ElementAtOrNil(tt.args.index)
			if got != tt.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestElementAtOrDefault(t *testing.T) {
	type args struct {
		index        int64
		defaultValue int
	}
	tests := []struct {
		name   string
		source []int
		args   args
		want   int
	}{
		{
			name:   "ElementAtOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -99,
				index:        0,
			},
			want: 1,
		},
		{
			name:   "ElementAtOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -999,
				index:        2,
			},
			want: 3,
		},
		{
			name:   "ElementAtOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -9,
				index:        9,
			},
			want: -9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).ElementAtOrDefault(tt.args.index, tt.args.defaultValue)
			if got != tt.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestPFirstOrNil(t *testing.T) {
	type args struct {
		p lingo.Predicate[int]
	}
	tests := []struct {
		name      string
		source    []int
		args      args
		want      int
		isOrdered bool
	}{
		{
			name:   "FirstOrNil",
			source: []int{1, 2, 3},
			want:   1,
		},
		{
			name:   "FirstOrNil",
			source: []int{1, 2, 3},
			args:   args{p: func(v int) bool { return v > 2 }},
			want:   3,
		},
		{
			name:   "FirstOrNil",
			source: []int{1, 2, 3},
			args:   args{p: func(v int) bool { return v > 5 }},
			want:   0,
		},
		{
			name:      "FirstOrNil",
			source:    []int{1, 2, 3},
			want:      1,
			isOrdered: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			penum := lingo.AsEnumerable(tt.source).AsParallel()
			if tt.isOrdered {
				penum = penum.AsOrdered()
			}
			got := penum.FirstOrNil(tt.args.p)
			if !lingo.AsEnumerable(tt.source).Append(tt.want).Contains(got) {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestPFirstOrDefault(t *testing.T) {
	type args struct {
		p            lingo.Predicate[int]
		defaultValue int
	}
	tests := []struct {
		name      string
		source    []int
		args      args
		want      int
		isOrdered bool
	}{
		{
			name:   "FirstOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -99,
			},
			want: 1,
		},
		{
			name:   "FirstOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -999,
				p:            func(v int) bool { return v > 2 },
			},
			want: 3,
		},
		{
			name:   "FirstOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -9,
				p:            func(v int) bool { return v > 10 },
			},
			want: -9,
		},
		{
			name:   "FirstOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -999,
				p:            func(v int) bool { return v > 2 },
			},
			want:      3,
			isOrdered: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			penum := lingo.AsEnumerable(tt.source).AsParallel()
			if tt.isOrdered {
				penum = penum.AsOrdered()
			}
			got := penum.FirstOrDefault(tt.args.defaultValue, tt.args.p)
			if !lingo.AsEnumerable(tt.source).Append(tt.want).Contains(got) {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestPLastOrNil(t *testing.T) {
	type args struct {
		p lingo.Predicate[int]
	}
	tests := []struct {
		name      string
		source    []int
		args      args
		want      int
		isOrdered bool
	}{
		{
			name:   "LastOrNil",
			source: []int{1, 2, 3},
			want:   3,
		},
		{
			name:   "LastOrNil",
			source: []int{1, 2, 3},
			args:   args{p: func(v int) bool { return v > 2 }},
			want:   3,
		},
		{
			name:   "LastOrNil",
			source: []int{1, 2, 3},
			args:   args{p: func(v int) bool { return v > 5 }},
			want:   0,
		},
		{
			name:      "LastOrNil",
			source:    []int{1, 2, 3},
			want:      3,
			isOrdered: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			penum := lingo.AsEnumerable(tt.source).AsParallel()
			if tt.isOrdered {
				penum = penum.AsOrdered()
			}
			got := penum.LastOrNil(tt.args.p)
			if !lingo.AsEnumerable(tt.source).Append(tt.want).Contains(got) {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestPLastOrDefault(t *testing.T) {
	type args struct {
		p            lingo.Predicate[int]
		defaultValue int
	}
	tests := []struct {
		name      string
		source    []int
		args      args
		want      int
		isOrdered bool
	}{
		{
			name:   "LastOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -99,
			},
			want: 3,
		},
		{
			name:   "LastOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -999,
				p:            func(v int) bool { return v > 2 },
			},
			want: 3,
		},
		{
			name:   "LastOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -9,
				p:            func(v int) bool { return v > 10 },
			},
			want: -9,
		},
		{
			name:   "LastOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -99,
			},
			want:      3,
			isOrdered: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			penum := lingo.AsEnumerable(tt.source).AsParallel()
			if tt.isOrdered {
				penum = penum.AsOrdered()
			}
			got := penum.LastOrDefault(tt.args.defaultValue, tt.args.p)
			if !lingo.AsEnumerable(tt.source).Append(tt.want).Contains(got) {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestPElementAtOrNil(t *testing.T) {
	type args struct {
		index int64
	}
	tests := []struct {
		name      string
		source    []int
		args      args
		want      int
		isOrdered bool
	}{
		{
			name:   "ElementAtOrNil",
			source: []int{1, 2, 3},
			args:   args{index: 1},
			want:   2,
		},
		{
			name:   "ElementAtOrNil",
			source: []int{1, 2, 3},
			args:   args{index: 4},
			want:   0,
		},
		{
			name:   "ElementAtOrNil",
			source: []int{1, 2, 3},
			args:   args{index: 2},
			want:   3,
		},
		{
			name:      "ElementAtOrNil",
			source:    []int{1, 2, 3},
			args:      args{index: 1},
			want:      2,
			isOrdered: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			penum := lingo.AsEnumerable(tt.source).AsParallel()
			if tt.isOrdered {
				penum = penum.AsOrdered()
			}
			got := penum.ElementAtOrNil(tt.args.index)
			if !lingo.AsEnumerable(tt.source).Append(tt.want).Contains(got) {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestPElementAtOrDefault(t *testing.T) {
	type args struct {
		index        int64
		defaultValue int
	}
	tests := []struct {
		name      string
		source    []int
		args      args
		want      int
		isOrdered bool
	}{
		{
			name:   "ElementAtOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -99,
				index:        0,
			},
			want: 1,
		},
		{
			name:   "ElementAtOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -999,
				index:        2,
			},
			want: 3,
		},
		{
			name:   "ElementAtOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -9,
				index:        9,
			},
			want: -9,
		},
		{
			name:   "ElementAtOrDefault",
			source: []int{1, 2, 3},
			args: args{
				defaultValue: -99,
				index:        0,
			},
			want:      1,
			isOrdered: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			penum := lingo.AsEnumerable(tt.source).AsParallel()
			if tt.isOrdered {
				penum = penum.AsOrdered()
			}
			got := penum.ElementAtOrDefault(tt.args.index, tt.args.defaultValue)
			if !lingo.AsEnumerable(tt.source).Append(tt.want).Contains(got) {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
