package average_salary_excluding_the_minimum_and_maximum_salary

import (
	"testing"
)

func Test_average(t *testing.T) {
	type args struct {
		salary []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1",
			args: args{salary: []int{4000, 3000, 1000, 2000}},
			want: 2500,
		},
		{
			name: "2",
			args: args{salary: []int{3000, 1000, 2000}},
			want: 2000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.args.salary); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
