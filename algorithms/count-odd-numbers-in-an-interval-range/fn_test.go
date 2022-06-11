package count_odd_numbers_in_an_interval_range

import (
	"testing"
)

func Test_countOdds(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				low:  1,
				high: 3,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				low:  0,
				high: 3,
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				low:  3,
				high: 7,
			},
			want: 3,
		},
		{
			name: "4",
			args: args{
				low:  8,
				high: 10,
			},
			want: 1,
		},
		{
			name: "5",
			args: args{
				low:  1,
				high: 50,
			},
			want: 25,
		},
		{
			name: "6",
			args: args{
				low:  0,
				high: 1000000000,
			},
			want: 1000000000 / 2,
		},
		{
			name: "7",
			args: args{
				low:  111,
				high: 112,
			},
			want: 1,
		},
		{
			name: "8",
			args: args{
				low:  11,
				high: 11,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOdds(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("countOdds() = %v, want %v", got, tt.want)
			}
		})
	}
}
