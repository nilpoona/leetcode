package string_to_integer_atoi

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{s: "123"},
			want: 123,
		},
		{
			args: args{s: "1234567890"},
			want: 1234567890,
		},
		{
			args: args{s: "+77373"},
			want: 77373,
		},
		{
			args: args{s: "-88833883"},
			want: -88833883,
		},
		{
			args: args{s: "9223372036854775808"},
			want: 2147483647,
		},
		{
			args: args{s: "-91283472332"},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
