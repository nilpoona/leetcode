package palindrome_number

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{x: 121},
			want: true,
		},
		{
			args: args{x: -121},
			want: false,
		},
		{
			args: args{x: 10},
			want: false,
		},
		{
			args: args{x: 9},
			want: true,
		},
		{
			args: args{x: 0},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
