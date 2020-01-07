package cmp

import "testing"

func TestCompareOnList(t *testing.T) {
	var testCase = []struct {
		In     *ListNode
		Expect *ListNode
	}{
		{MakeList(1, 2, 3, 4, 5), MakeList(1, 2, 3, 4, 5)},
	}

	for i := range testCase {
		if !CompareOnList(testCase[i].In, testCase[i].Expect) {
			t.Errorf("compare on list failed.")
		}
	}
}
