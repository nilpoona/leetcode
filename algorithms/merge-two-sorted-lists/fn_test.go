package merge_two_sorted_list

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
		wantList []int
	}{
		{
			args: args{
				list1: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				list2: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			wantList: []int{1,1,2,3,4,4},
		},
		{
			args: args{
				list1: &ListNode{
					Val:  2,
					Next: nil,
				},
				list2: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			wantList: []int{1, 2},
		},

		{
			args: args{
				list1: &ListNode{
					Val:  50,
					Next: &ListNode{
						Val:  60,
						Next: &ListNode{
							Val:  60,
						},
					},
				},
				list2: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
			wantList: []int{3, 50, 60, 60},
		},

		{
			args: args{
				list1: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{
							Val:  3,
						},
					},
				},
				list2: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
			wantList: []int{3, 3, 3, 3},
		},
		{
			args: args{
				list1: &ListNode{
					Val:  2,
					Next: &ListNode{
						Val:  5,
					},
				},
				list2: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{
							Val:  4,
						},
					},
				},
			},
			wantList: []int{1, 2, 3, 4, 5},
		},

		{
			args: args{
				list1: &ListNode{
					Val:  2,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{
							Val:  2,
							Next: &ListNode{
								Val:  5,
							},
						},
					},
				},
				list2: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{
							Val:  4,
						},
					},
				},
			},
			wantList: []int{1, 2, 2, 2, 3, 4, 5},
		},
		{
			args: args{
				list1: nil,
				list2: nil,
			},
			wantList: []int{},
		},

		{
			args: args{
				list1: nil,
				list2: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			wantList: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.list1, tt.args.list2)
			l := []int{}
			for {
				if got == nil {
					break
				}
				l = append(l, got.Val)
				got = got.Next
			}
			if !reflect.DeepEqual(l, tt.wantList) {
				t.Errorf("mergeTwoLists() = %v, want %v", l, tt.wantList)
			}
		})
	}
}
