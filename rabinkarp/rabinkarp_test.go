package rabinkarp

import "testing"

func TestRK(t *testing.T) {
	type args struct {
		src     string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "rk_test",
			args: args{
				src:     "hello world",
				pattern: "world",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RK(tt.args.src, tt.args.pattern); got != tt.want {
				t.Errorf("RK() = %v, want %v", got, tt.want)
			}
		})
	}
}

var pattern = "world"

func BenchmarkRK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RK("hello world", pattern)
	}
}

func BenchmarkXXHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		XXHash("hello world")
	}
}
