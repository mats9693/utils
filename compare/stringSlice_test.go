package cmp

import "testing"

func TestCompareOnStringSlice(t *testing.T) {
	var testCase = []struct {
		In     []string
		Expect []string
	}{
		{[]string{"ab", "cd", "ef"}, []string{"cd", "ab", "ef"}},
	}

	for i := range testCase {
		if !CompareOnStringSlice(testCase[i].In, testCase[i].Expect) {
			t.Errorf("compare on string slice failed.")
		}
	}
}
