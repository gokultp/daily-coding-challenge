package main

import "testing"

func Test_largestSumNonAdjascent(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name:"test", args:args{data: []int{2, 4, 6, 2, 5}}, want: 13},
		{name:"test", args:args{data: []int{2, 4, -1, 2, 5}}, want: 9},
		{name:"test", args:args{data: []int{5, 1 ,1 ,5}}, want: 10},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumNonAdjascent(tt.args.data); got != tt.want {
				t.Errorf("largestSumNonAdjascent() = %v, want %v", got, tt.want)
			}
		})
	}
}
