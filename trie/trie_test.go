package trie

import "testing"

func Test_node_IsExist(t *testing.T) {
	type args struct {
		word string
	}

	root := newNode()
	root.Insert("hello")

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_is_exist",
			args: args{word: "hello"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := root.IsExist(tt.args.word); got != tt.want {
				t.Errorf("IsExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
