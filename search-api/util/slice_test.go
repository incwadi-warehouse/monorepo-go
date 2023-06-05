package util

import "testing"

func TestContains(t *testing.T) {
	type args struct {
		needle string
		haystack   []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test products",
			args{
				"products",
				[]string{"products"},
			},
			true,
		},
		{
			"test products",
			args{
				"product",
				[]string{"products"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.needle, tt.args.haystack); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
