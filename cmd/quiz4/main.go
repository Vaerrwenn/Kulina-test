package main

import "fmt"

func main() {
	// Initialize 100 lamps. All lamps are off.
	lamps := [100]int{}

	// Loop for the trips
	for trip := 1; trip <= 100; trip++ {
		// Loop for the lamps
		for idx := 0; idx < len(lamps); idx++ {
			if idx%trip == 0 {
				if lamps[idx] == 0 {
					lamps[idx] = 1
				} else {
					lamps[idx] = 0
				}
			}
		}
	}

	totalLampsOn := 0
	for _, lamp := range lamps {
		if lamp == 1 {
			totalLampsOn += 1
		}
	}

	fmt.Println(totalLampsOn)
}
