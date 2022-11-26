package addtwonumbers_test

import (
	"reflect"
	"testing"

	a "github.com/SergeyParamoshkin/leetcode/add-two-numbers"
)

func Test_AddTwoNumbers(t *testing.T) {
	t.Parallel()

	type args struct {
		l1 *a.ListNode
		l2 *a.ListNode
	}

	tests := []struct {
		name string
		args args
		want *a.ListNode
	}{
		{
			name: "okok",
			args: args{
				l1: &a.ListNode{
					Val: 2,
					Next: &a.ListNode{
						Val: 4,
						Next: &a.ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &a.ListNode{
					Val: 5,
					Next: &a.ListNode{
						Val: 6,
						Next: &a.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &a.ListNode{
				Val: 7,
				Next: &a.ListNode{
					Val: 0,
					Next: &a.ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := a.AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
