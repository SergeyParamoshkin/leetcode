package mergetwoschedule_test

import (
	"reflect"
	"testing"

	"github.com/SergeyParamoshkin/leetcode/mergetwoschedule"
)

func TestMerge(t *testing.T) {
	t.Parallel()

	type args struct {
		a []mergetwoschedule.Point
		b []mergetwoschedule.Point
	}

	tests := []struct {
		name string
		args args
		want []mergetwoschedule.Point
	}{
		{name: "ok", args: args{
			a: []mergetwoschedule.Point{{1, 2}, {5, 1}},
			b: []mergetwoschedule.Point{{2, 4}, {3, 6}, {9, 7}},
		}, want: []mergetwoschedule.Point{{1, 2}, {2, 6}, {3, 8}, {5, 7}, {9, 8}}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := mergetwoschedule.Merge(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
