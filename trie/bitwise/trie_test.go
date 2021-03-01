package bitwise

import (
	"fmt"
	"testing"
)

func Test_node_Insert(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_bitwise_insert",
			args: args{word: "hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := newNode()
			root.Insert("hello")
		})
	}
}

func Test_node_IsExist(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_bitwise_isExist",
			args: args{word: "hello"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := newNode()
			root.Insert("hello")
			root.Insert("world")
			if got := root.IsExist(tt.args.word); got != tt.want {
				t.Errorf("IsExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNode_Insert(b *testing.B) {
	r := newNode()
	for i := 0; i < b.N; i++ {
		r.Insert("hello")
	}
}

var root = func() *node {
	r := newNode()
	for i := 0; i < 100; i++ {
		r.Insert(fmt.Sprintf("hello%d", i))
	}
	return r
}()

func BenchmarkNode_IsExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		root.IsExist("hello0")
	}
}
