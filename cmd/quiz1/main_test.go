package main

import (
	"fmt"
	"testing"
)

func TestSockMerchant(t *testing.T) {
	tests := []struct {
		n      int
		ar     []int
		expect int
	}{
		{9, []int{10, 20, 20, 10, 10, 30, 50, 10, 20}, 3},
		{4, []int{3, 1, 5, 4}, 0},
		{7, []int{4, 9, 4, 3, 4, 2, 1}, 1},
		{5, []int{1, 1, 1, 1, 1}, 2},
		{9, []int{3, 4, 5, 1, 1, 5, 7, 3, 9}, 3},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("n = %d, ar = %v", tc.n, tc.ar), func(t *testing.T) {
			result := sockMerchant(tc.n, tc.ar)
			if result != tc.expect {
				t.Fatalf("Result: %d. Expected result: %d. MISMATCH.", result, tc.expect)
			}
			t.Logf("Result: %d. Expected: %d. MATCH.", result, tc.expect)
		})
	}
}
