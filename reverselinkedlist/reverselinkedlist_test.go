package reverselinkedlist_test

import (
	"reflect"
	"testing"

	rll "github.com/SergeyParamoshkin/leetcode/reverselinkedlist"
)

func Test_reverseList(t *testing.T) {
	t.Parallel()

	type args struct{ head *rll.ListNode }

	tests := []struct {
		name string
		args args
		want *rll.ListNode
	}{
		{
			name: "1,2->2,1",
			args: args{
				head: &rll.ListNode{
					Val: 1,
					Next: &rll.ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			want: &rll.ListNode{
				Val: 2,
				Next: &rll.ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := rll.ReverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
