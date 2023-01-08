package removesmiles_test

import (
	"testing"

	"github.com/SergeyParamoshkin/leetcode/removesmiles"
)

func TestRemoveSmiles(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "ok", args: args{s: "Я работаю в Гугле :-)))"}, want: "Я работаю в Гугле "},
		{name: "okok", args: args{s: "Aaaaa!!!!! :-))(())"}, want: "Aaaaa!!!!! (())"},
		{name: "okok", args: args{s: "везет :-) а я туда собеседование завалил:-(("}, want: "везет  а я туда собеседование завалил"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removesmiles.RemoveSmiles(tt.args.s); got != tt.want {
				t.Errorf("RemoveSmiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
