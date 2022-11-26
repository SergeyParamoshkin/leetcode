package mergetwosortedlists_test

import (
	"reflect"
	"testing"

	msl "github.com/SergeyParamoshkin/leetcode/mergetwosortedlists"
)

func Test_mergeTwoLists(t *testing.T) {
	t.Parallel()

	type args struct{ list1, list2 *msl.ListNode }

	tests := []struct {
		name string
		args args
		want *msl.ListNode
	}{
		{name: "empty list", args: args{list1: nil, list2: nil}, want: nil},
		{
			name: "[1,2,4]_[1,3,4]",
			args: args{
				list1: &msl.ListNode{
					Val: 1,
					Next: &msl.ListNode{
						Val:  2,
						Next: &msl.ListNode{Val: 4, Next: nil},
					},
				},
				list2: &msl.ListNode{
					Val: 1,
					Next: &msl.ListNode{
						Val:  3,
						Next: &msl.ListNode{Val: 4, Next: nil},
					},
				},
			}, want: &msl.ListNode{
				Val: 1, Next: &msl.ListNode{
					Val: 1,
					Next: &msl.ListNode{
						Val: 2,
						Next: &msl.ListNode{
							Val: 3,
							Next: &msl.ListNode{
								Val:  4,
								Next: &msl.ListNode{Val: 4, Next: nil},
							},
						},
					},
				},
			},
		},
		{
			name: "[]_[0]",
			args: args{
				list1: &msl.ListNode{Val: 0, Next: nil}, list2: &msl.ListNode{Val: 0, Next: nil},
			}, want: &msl.ListNode{Val: 0, Next: &msl.ListNode{Val: 0, Next: nil}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := msl.MergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
