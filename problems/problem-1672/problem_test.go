package problem1672

import "testing"

func TestMaximumWealth(t *testing.T) {
	tests := []struct {
		accounts [][]int
		expect   int
	}{
		{accounts: [][]int{{1, 2, 3}, {3, 2, 1}}, expect: 6},
		{accounts: [][]int{{1, 5}, {7, 3}, {3, 5}}, expect: 10},
		{accounts: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, expect: 17},
		// edge case: only one customer, one bank
		{accounts: [][]int{{42}}, expect: 42},
		// edge case: all customers have the same wealth
		{accounts: [][]int{{5, 5}, {10, 0}, {2, 8}}, expect: 10},
	}

	for i, tc := range tests {
		got := MaximumWealth(tc.accounts)
		if got != tc.expect {
			t.Errorf("case %d: got %d, want %d", i+1, got, tc.expect)
		}
	}
}
