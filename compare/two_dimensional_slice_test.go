package cmp

import "testing"

func TestCompareOnErWeiInt(t *testing.T) {
	var testCase = []struct {
		In     [][]int
		Expect [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 8, 9}, {1, 2, 3}, {6, 5, 4}}},
	}

	for i := range testCase {
		if !CompareOnTwoDimensionalSlice(testCase[i].In, testCase[i].Expect) {
			t.Errorf("compare on er-wei int failed.")
		}
	}
}
