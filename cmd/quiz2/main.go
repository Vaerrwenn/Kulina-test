package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// First input, number of steps
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("ERROR: Invalid Input.")
	}
	if n < 2 || n > int(math.Pow10(6)) || n%2 != 0 {
		log.Fatal("ERROR: n must be between 2 and 10^6 and an EVEN number.")
	}

	// Second input, steps description.
	scanner.Scan()
	s := strings.ToUpper(scanner.Text())
	if stepLen := len(strings.Split(s, " ")); stepLen != n {
		log.Fatalf("ERROR: Steps description is unequal to inputted number of steps.\n"+
			"Number of steps: %d\nSteps description: %d", n, stepLen)
	}

	result := countingValleys(n, s)
	fmt.Println(result)
}

// countingValleys counts how much valley. Takes:
//
// - n as number of steps
//
// - s as steps description (must be either U (for upward) or D (for downward)).
//
// Returns how many valleys he passed.
func countingValleys(n int, s string) int {
	/*
		1. Split the s.
		2. Make an int slice.
		3. Loop 4-6
		4. If step == U, append 1 to the int slice.
		5. Else If step == D, append -1 to the int slice.
		6. Else, error
		7. Loop 8-13
		8. If stepVal == -1 AND totalAlt == 0, loop with i == idx + 1 9-11
		9. if stepVals[idx+i] == 1, totalUp += 1, currAlt += 1
		10. else, totalDown += 1, currAlt -=1
		11. if currAlt == 0, break loop 8.
		12.	if totalUp == totalDown, valley += 1
		13. totalAlt += stepVal
		14. if totalAlt != 0, error.
	*/
	var valley int
	steps := strings.Split(s, " ")

	stepValues := make([]int, n)

	for idx, step := range steps {
		if step == "U" {
			stepValues[idx] = 1
		} else if step == "D" {
			stepValues[idx] = -1
		} else {
			log.Fatal("ERROR: Step description must be either U or D.")
		}
	}

	valley = 0
	totalAlt := 0
	for idx, stepValue := range stepValues {
		if stepValue == -1 && totalAlt == 0 && stepValues[idx+1] != 1 {
			totalUp := 0
			totalDown := 0
			currAlt := 0
			for i := idx + 1; i < len(stepValues); i++ {
				if stepValues[i] == 1 {
					totalUp += 1
					currAlt += 1
				} else {
					totalDown += 1
					currAlt -= 1
				}
				if currAlt == 0 {
					break
				}
			}
			if totalUp == totalDown {
				valley += 1
			}
		}
		totalAlt += stepValue
	}

	if totalAlt != 0 {
		log.Fatal("ERROR: Gary/Bill did not finish his hike.")
	}
	return valley
}
