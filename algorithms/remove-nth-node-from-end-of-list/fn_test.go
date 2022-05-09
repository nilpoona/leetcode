package remove_nth_node_from_end_of_list

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	genListNode := func(nums []int) *ListNode {
		var prev, first *ListNode
		for i, num := range nums {
			node := &ListNode{
				Val: num,
			}
			if prev != nil {
				prev.Next = node
			}
			if i == 0 {
				first = node
			}
			prev = node
		}
		return first
	}
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				head: genListNode([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "2",
			args: args{
				head: genListNode([]int{1}),
				n:    1,
			},
			want: []int{},
		},
		{
			name: "3",
			args: args{
				head: genListNode([]int{1, 2}),
				n:    1,
			},
			want: []int{1},
		},
		{
			name: "4",
			args: args{
				head: genListNode([]int{1, 2}),
				n:    2,
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNthFromEnd(tt.args.head, tt.args.n)
			l := []int{}
			for {
				if got == nil {
					break
				}
				l = append(l, got.Val)
				got = got.Next
			}
			if !reflect.DeepEqual(l, tt.want) {
				t.Errorf("removeNthFromEnd() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
