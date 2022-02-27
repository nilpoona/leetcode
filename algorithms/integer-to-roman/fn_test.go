package integer_to_roman

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{num: 3},
			want: "III",
		},
		{
			args: args{num: 58},
			want: "LVIII",
		},
		{
			args: args{num: 1994},
			want: "MCMXCIV",
		},
		{
			args: args{num: 1},
			want: "I",
		},
		{
			args: args{num: 5},
			want: "V",
		},
		{
			args: args{num: 10},
			want: "X",
		},
		{
			args: args{num: 50},
			want: "L",
		},
		{
			args: args{num: 100},
			want: "C",
		},
		{
			args: args{num: 500},
			want: "D",
		},
		{
			args: args{num: 1000},
			want: "M",
		},
		{
			args: args{num: 4},
			want: "IV",
		},
		{
			args: args{num: 9},
			want: "IX",
		},
		{
			args: args{num: 40},
			want: "XL",
		},
		{
			args: args{num: 90},
			want: "XC",
		},
		{
			args: args{num: 400},
			want: "CD",
		},
		{
			args: args{num: 900},
			want: "CM",
		},
		{
			args: args{num: 3999},
			want: "MMMCMXCIX",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
