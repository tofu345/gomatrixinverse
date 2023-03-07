/*
 * Title: Q5 Practise Sheet 7 Overkill
 * Author: Oluwatofunmi Yinka-Adebimpe
 * Date: 26 February 2023 21:35
 */

package main

import "testing"

func Test_matrixInverse(t *testing.T) {
	type args struct {
		mat *[3][3]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"tests", args{&[3][3]int{
			{12, 34, 56},
			{3490, 3104, -100},
			{20, 19, 18},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrixInverse(tt.args.mat)
		})
	}
}
