package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{1, 2}, []int{2, 1}},
	}
	for _, c := range cases {
		got := toSlice(Reverse(fromSlice(c.in)))
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Reverse(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
