package lingo_test

import (
	"testing"

	lingo "github.com/vanh01/lingo"
)

func TestAnd(t *testing.T) {
	type args struct {
		p     lingo.Predicate[int]
		right lingo.Predicate[int]
	}
	tests := []struct {
		name   string
		source int
		args   args
		want   bool
	}{
		{
			name:   "And",
			source: 5,
			args: args{
				p:     func(v int) bool { return v > 2 },
				right: func(i int) bool { return i < 10 },
			},
			want: true,
		},
		{
			name:   "And",
			source: 2,
			args: args{
				p:     func(v int) bool { return v > 2 },
				right: func(i int) bool { return i < 10 },
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.p.And(tt.args.right)(tt.source)
			if got != tt.want {
				t.Errorf("And() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOr(t *testing.T) {
	type args struct {
		p     lingo.Predicate[int]
		right lingo.Predicate[int]
	}
	tests := []struct {
		name   string
		source int
		args   args
		want   bool
	}{
		{
			name:   "Or",
			source: 5,
			args: args{
				p:     func(v int) bool { return v > 2 },
				right: func(i int) bool { return i < 10 },
			},
			want: true,
		},
		{
			name:   "Or",
			source: 2,
			args: args{
				p:     func(v int) bool { return v > 2 },
				right: func(i int) bool { return i < 10 },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.p.Or(tt.args.right)(tt.source)
			if got != tt.want {
				t.Errorf("Or() = %v, want %v", got, tt.want)
			}
		})
	}
}
