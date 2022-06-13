package subtract_the_product_and_sum_of_digits_of_an_integer

import "testing"

func Test_subtractProductAndSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{n: 234},
			want: 15,
		},
		{
			name: "2",
			args: args{n: 4421},
			want: 21,
		},
		{
			name: "3",
			args: args{n: 1},
			want: 0,
		},
		{
			name: "4",
			args: args{n: 100000},
			want: -1,
		},
		{
			name: "5",
			args: args{n: 9901},
			want: -19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtractProductAndSum(tt.args.n); got != tt.want {
				t.Errorf("subtractProductAndSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
