package roman_to_integer

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{s: "III"},
			want:  3,
		},
		{
			args: args{s: "VI"},
			want:  6,
		},
		{
			args: args{s: "L"},
			want:  50,
		},
		{
			args: args{s: "LVIII"},
			want:  58,
		},
		{
			args: args{s: "XCIX"},
			want:  99,
		},
		{
			args: args{s: "C"},
			want:  100,
		},
		{
			args: args{s: "CD"},
			want:  400,
		},
		{
			args: args{s: "D"},
			want:  500,
		},
		{
			args: args{s: "MCMXCIV"},
			want:  1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
