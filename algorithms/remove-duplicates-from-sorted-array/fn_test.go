package remove_duplicates_from_sorted_array

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{nums: []int{0}},
			want: []int{0},
		},
		{
			args: args{nums: []int{1, 1, 2}},
			want: []int{1, 2},
		},
		{
			args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4}},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 5}},
			want: []int{0, 1, 2, 3, 4, 5},
		},

		{
			args: args{nums: []int{-100, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 5}},
			want: []int{-100, 0, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.nums)
			got := removeDuplicates(tt.args.nums)
			if got != len(tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			t.Log(tt.args.nums)
			t.Log("---------------------------------------------")
			for i := 0; i < got; i++ {
				if tt.args.nums[i] != tt.want[i] {
					t.Errorf("removeDuplicates() = %+v, want %+v", tt.args.nums, tt.want)
				}
			}
		})
	}
}
