package mergetwosortedlists

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
	}{
		{name: "empty list", args: args{}, want: nil},
		{
			name: "[1,2,4]_[1,3,4]",
			args: args{
				list1: &ListNode{Val: 1,
					Next: &ListNode{Val: 2,
						Next: &ListNode{Val: 4}}},
				list2: &ListNode{Val: 1,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 4}}}},
			want: &ListNode{Val: 1,
				Next: &ListNode{Val: 1,
					Next: &ListNode{Val: 2,
						Next: &ListNode{Val: 3,
							Next: &ListNode{Val: 4,
								Next: &ListNode{Val: 4}}}}}}},
		{
			name: "[]_[0]",
			args: args{
				list1: &ListNode{},
				list2: &ListNode{Val: 0}},
			want: &ListNode{Val: 0, Next: &ListNode{}},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
