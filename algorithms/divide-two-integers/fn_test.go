package divide_two_integers

import (
	"fmt"
	"testing"
)

func Test_divide(t *testing.T) {
	fmt.Println(3 << 0)
	fmt.Println(3 << 1)
	fmt.Println(3 << 2)
	fmt.Println(1 << 2)

	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: struct {
				dividend int
				divisor  int
			}{dividend: 10, divisor: 3},
			want: 3,
		},

		{
			args: struct {
				dividend int
				divisor  int
			}{dividend: -7, divisor: -3},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
