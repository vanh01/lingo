package lingo_test

import (
	"testing"

	lingo "github.com/vanh01/lingo"
)

func TestAll(t *testing.T) {
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
		want   bool
	}{
		{
			name: "All",
			source: []Student{
				{Id: 1, Name: "Anh", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Nguyen", Level: 3},
			},
			args: args{p: func(s Student) bool {
				return s.Level > 0
			}},
			want: true,
		},
		{
			name: "All",
			source: []Student{
				{Id: 1, Name: "Anh", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Nguyen", Level: 3},
			},
			args: args{p: func(s Student) bool {
				return s.Level > 1
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).All(tt.args.p)
			if got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny(t *testing.T) {
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
		want   bool
	}{
		{
			name: "Any",
			source: []Student{
				{Id: 1, Name: "Anh", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Nguyen", Level: 3},
			},
			args: args{p: func(s Student) bool {
				return s.Level >= 3
			}},
			want: true,
		},
		{
			name: "Any",
			source: []Student{
				{Id: 1, Name: "Anh", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Nguyen", Level: 3},
			},
			args: args{p: func(s Student) bool {
				return s.Name == "Nguyen"
			}},
			want: true,
		},
		{
			name: "Any",
			source: []Student{
				{Id: 1, Name: "Anh", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Nguyen", Level: 3},
			},
			args: args{p: func(s Student) bool {
				return s.Name == "Nguyen "
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).Any(tt.args.p)
			if got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type Student struct {
		Id    int
		Name  string
		Level int
	}
	type args struct {
		value Student
		c     lingo.Comparer[Student]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   bool
	}{
		{
			name: "Contains",
			source: []Student{
				{Id: 1, Name: "Anh", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Nguyen", Level: 3},
			},
			args: args{
				value: Student{Id: 1},
				c: func(s1, s2 Student) bool {
					return s1.Id == s2.Id
				},
			},
			want: true,
		},
		{
			name: "Contains",
			source: []Student{
				{Id: 1, Name: "Anh", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Nguyen", Level: 3},
			},
			args: args{
				value: Student{Id: 5},
				c: func(s1, s2 Student) bool {
					return s1.Id == s2.Id
				},
			},
			want: false,
		},
		{
			name: "Contains",
			source: []Student{
				{Id: 1, Name: "Anh", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Nguyen", Level: 3},
			},
			args: args{
				value: Student{Id: 1, Name: "Anh", Level: 1},
			},
			want: true,
		},
		{
			name: "Contains",
			source: []Student{
				{Id: 1, Name: "Anh", Level: 1},
				{Id: 2, Name: "An", Level: 2},
				{Id: 3, Name: "Nguyen", Level: 3},
			},
			args: args{
				value: Student{Id: 1, Name: "Anh", Level: 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsEnumerable(tt.source).Contains(tt.args.value, tt.args.c)
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
