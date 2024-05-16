package lingo_test

import (
	"testing"

	lingo "github.com/vanh01/lingo"
	"github.com/vanh01/lingo/definition"
)

func TestAsLookup(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args struct {
		keySelector     definition.SingleSelectorFull[Student, int]
		elementSelector definition.SingleSelectorFull[Student, any]
	}
	type want struct {
		count int
		item  map[int][]any
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   want
	}{
		{
			name: "AsLookup",
			source: []Student{
				{Id: 1, Name: "Anh", ClassId: 1},
				{Id: 2, Name: "hnA", ClassId: 2},
				{Id: 3, Name: "Abcd", ClassId: 3},
				{Id: 4, Name: "Ank", ClassId: 1},
				{Id: 5, Name: "hnI", ClassId: 2},
				{Id: 6, Name: "A", ClassId: 3},
			},
			args: args{
				keySelector: func(s Student) int {
					return s.ClassId
				},
			},
			want: want{
				count: 3,
				item: map[int][]any{
					1: {
						Student{Id: 1, Name: "Anh", ClassId: 1},
						Student{Id: 4, Name: "Ank", ClassId: 1},
					},
					2: {
						Student{Id: 2, Name: "hnA", ClassId: 2},
						Student{Id: 5, Name: "hnI", ClassId: 2},
					},
					3: {
						Student{Id: 3, Name: "Abcd", ClassId: 3},
						Student{Id: 6, Name: "A", ClassId: 3},
					},
				},
			},
		},
		{
			name: "AsLookup",
			source: []Student{
				{Id: 1, Name: "Anh", ClassId: 1},
				{Id: 2, Name: "hnA", ClassId: 2},
				{Id: 3, Name: "Abcd", ClassId: 3},
				{Id: 4, Name: "Ank", ClassId: 1},
				{Id: 5, Name: "hnI", ClassId: 2},
				{Id: 6, Name: "A", ClassId: 3},
			},
			args: args{
				keySelector: func(s Student) int {
					return s.ClassId
				},
				elementSelector: func(s Student) any {
					return s.Name
				},
			},
			want: want{
				count: 3,
				item: map[int][]any{
					1: {"Anh", "Ank"},
					2: {"hnA", "hnI"},
					3: {"Abcd", "A"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsLookup(lingo.AsEnumerable(tt.source), tt.args.keySelector, tt.args.elementSelector)
			if got.Count != tt.want.count {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
			for group := range got.GetIter() {
				i := 0
				for v := range group.GetIter() {
					if v != tt.want.item[group.Key][i] {
						t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
					}
					i++
				}
			}
		})
	}
}

func TestContainsKey(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args struct {
		keySelector     definition.SingleSelectorFull[Student, int]
		elementSelector definition.SingleSelectorFull[Student, any]
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		key    int
		want   bool
	}{
		{
			name: "ContainsKey",
			source: []Student{
				{Id: 1, Name: "Anh", ClassId: 1},
				{Id: 2, Name: "hnA", ClassId: 2},
				{Id: 3, Name: "Abcd", ClassId: 3},
				{Id: 4, Name: "Ank", ClassId: 1},
				{Id: 5, Name: "hnI", ClassId: 2},
				{Id: 6, Name: "A", ClassId: 3},
			},
			args: args{
				keySelector: func(s Student) int {
					return s.ClassId
				},
			},
			key:  1,
			want: true,
		},
		{
			name: "ContainsKey",
			source: []Student{
				{Id: 1, Name: "Anh", ClassId: 1},
				{Id: 2, Name: "hnA", ClassId: 2},
				{Id: 3, Name: "Abcd", ClassId: 3},
				{Id: 4, Name: "Ank", ClassId: 1},
				{Id: 5, Name: "hnI", ClassId: 2},
				{Id: 6, Name: "A", ClassId: 3},
			},
			args: args{
				keySelector: func(s Student) int {
					return s.ClassId
				},
			},
			key:  0,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsLookup(lingo.AsEnumerable(tt.source), tt.args.keySelector, tt.args.elementSelector)
			if got.ContainsKey(tt.key) != tt.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestGetValue(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args struct {
		keySelector     definition.SingleSelectorFull[Student, int]
		elementSelector definition.SingleSelectorFull[Student, any]
	}
	type want struct {
		count int
		item  map[int][]any
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   want
	}{
		{
			name: "AsLookup",
			source: []Student{
				{Id: 1, Name: "Anh", ClassId: 1},
				{Id: 2, Name: "hnA", ClassId: 2},
				{Id: 3, Name: "Abcd", ClassId: 3},
				{Id: 4, Name: "Ank", ClassId: 1},
				{Id: 5, Name: "hnI", ClassId: 2},
				{Id: 6, Name: "A", ClassId: 3},
			},
			args: args{
				keySelector: func(s Student) int {
					return s.ClassId
				},
			},
			want: want{
				count: 3,
				item: map[int][]any{
					1: {
						Student{Id: 1, Name: "Anh", ClassId: 1},
						Student{Id: 4, Name: "Ank", ClassId: 1},
					},
					2: {
						Student{Id: 2, Name: "hnA", ClassId: 2},
						Student{Id: 5, Name: "hnI", ClassId: 2},
					},
					3: {
						Student{Id: 3, Name: "Abcd", ClassId: 3},
						Student{Id: 6, Name: "A", ClassId: 3},
					},
				},
			},
		},
		{
			name: "AsLookup",
			source: []Student{
				{Id: 1, Name: "Anh", ClassId: 1},
				{Id: 2, Name: "hnA", ClassId: 2},
				{Id: 3, Name: "Abcd", ClassId: 3},
				{Id: 4, Name: "Ank", ClassId: 1},
				{Id: 5, Name: "hnI", ClassId: 2},
				{Id: 6, Name: "A", ClassId: 3},
			},
			args: args{
				keySelector: func(s Student) int {
					return s.ClassId
				},
				elementSelector: func(s Student) any {
					return s.Name
				},
			},
			want: want{
				count: 3,
				item: map[int][]any{
					1: {"Anh", "Ank"},
					2: {"hnA", "hnI"},
					3: {"Abcd", "A"},
					4: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsLookup(lingo.AsEnumerable(tt.source), tt.args.keySelector, tt.args.elementSelector)
			if got.Count != tt.want.count {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
			for k, v := range tt.want.item {
				i := 0
				for _, v1 := range got.GetValue(k) {
					if v1 != v[i] {
						t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
					}
					i++
				}
			}
		})
	}
}

func TestAsPLookup(t *testing.T) {
	type Student struct {
		Id      int
		Name    string
		ClassId int
	}
	type args struct {
		keySelector     definition.SingleSelectorFull[Student, int]
		elementSelector definition.SingleSelectorFull[Student, any]
	}
	type want struct {
		count int
		item  map[int][]any
	}
	tests := []struct {
		name   string
		source []Student
		args   args
		want   want
	}{
		{
			name: "AsLookup",
			source: []Student{
				{Id: 1, Name: "Anh", ClassId: 1},
				{Id: 2, Name: "hnA", ClassId: 2},
				{Id: 3, Name: "Abcd", ClassId: 3},
				{Id: 4, Name: "Ank", ClassId: 1},
				{Id: 5, Name: "hnI", ClassId: 2},
				{Id: 6, Name: "A", ClassId: 3},
			},
			args: args{
				keySelector: func(s Student) int {
					return s.ClassId
				},
			},
			want: want{
				count: 3,
				item: map[int][]any{
					1: {
						Student{Id: 1, Name: "Anh", ClassId: 1},
						Student{Id: 4, Name: "Ank", ClassId: 1},
					},
					2: {
						Student{Id: 2, Name: "hnA", ClassId: 2},
						Student{Id: 5, Name: "hnI", ClassId: 2},
					},
					3: {
						Student{Id: 3, Name: "Abcd", ClassId: 3},
						Student{Id: 6, Name: "A", ClassId: 3},
					},
				},
			},
		},
		{
			name: "AsLookup",
			source: []Student{
				{Id: 1, Name: "Anh", ClassId: 1},
				{Id: 2, Name: "hnA", ClassId: 2},
				{Id: 3, Name: "Abcd", ClassId: 3},
				{Id: 4, Name: "Ank", ClassId: 1},
				{Id: 5, Name: "hnI", ClassId: 2},
				{Id: 6, Name: "A", ClassId: 3},
			},
			args: args{
				keySelector: func(s Student) int {
					return s.ClassId
				},
				elementSelector: func(s Student) any {
					return s.Name
				},
			},
			want: want{
				count: 3,
				item: map[int][]any{
					1: {"Anh", "Ank"},
					2: {"hnA", "hnI"},
					3: {"A", "Abcd"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lingo.AsPLookup(lingo.AsParallelEnumerable(tt.source), tt.args.keySelector, tt.args.elementSelector)
			if got.Count != tt.want.count {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
			for group := range got.GetIter() {
				for value := range group.GetIter() {
					if !lingo.AsEnumerable(tt.want.item[group.Key]).Contains(value) {
						t.Errorf("%s() = %v, want %v", tt.name, group.ToSlice(), tt.want.item[group.Key])
					}
				}
			}
		})
	}
}
