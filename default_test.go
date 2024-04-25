package lingo

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Hello",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "Hello",
			args: args{s: "abc"},
			want: "abc",
		},
		{
			name: "Hello",
			args: args{s: "aa"},
			want: "aa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.s); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
