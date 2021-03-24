package main

import (
	"fmt"
	"testing"
)

func TestCountingValley(t *testing.T) {
	tests := []struct {
		n      int
		s      string
		expect int
	}{
		{8, "U D D D U D U U", 1},
		{2, "U D", 0},
		{4, "U U D D", 0},
		{20, "U D D D D U U U U D D D U U D U D U D U", 2},
		{14, "U U D D D D U U D D D U U U", 2},
		{4, "D D U U", 1},
	}

	for _, testCase := range tests {
		t.Run(fmt.Sprintf("n = %d.\ns = %s.", testCase.n, testCase.s), func(t *testing.T) {
			valley := countingValleys(testCase.n, testCase.s)
			if valley != testCase.expect {
				t.Fatalf("Valley: %d. Expect: %d. MISMATCH.", valley, testCase.expect)
			}
			t.Logf("Valley: %d. Expect: %d. MATCH.", valley, testCase.expect)
		})
	}
}
