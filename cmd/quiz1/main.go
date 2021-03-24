package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// First input, sock quantity.
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("ERROR: Input socks quantity is not valid.")
	}
	if n < 1 || n > 100 {
		log.Fatal("ERROR: Socks quantity must be between 1 - 100.")
	}

	// Second input, colour of socks.
	scanner.Scan()
	colours := scanner.Text()

	coloursSlice := strings.Split(colours, " ")
	if len(coloursSlice) != n {
		log.Fatalf("ERROR: Quantity of colours is unequal to the quantity of the socks.\n"+
			"Colour quantity: %d.\nSocks quantity: %d.",
			len(coloursSlice), n)
	}

	ar := make([]int, n)
	for idx, colour := range coloursSlice {

		ar[idx], err = strconv.Atoi(colour)

		if err != nil {
			log.Fatalf("%v is not an integer.", colour)
		}
		if ar[idx] < 1 || ar[idx] > 100 {
			log.Fatal("Colour number must be between 1 - 100.")
		}
	}

	// TODO: refactor the code so n is actually useful in sockMerchant()
	result := sockMerchant(n, ar)

	fmt.Println(result)
}

// SockMerchant takes:
//
// - n  : the number of socks.
//
// - ar : the colours of the socks. Represented with numbers.
//
// Returns an int representing the number of matching pairs of socks that are available.
func sockMerchant(n int, ar []int) int {
	var result int

	// 1. Sort the slice
	// 2. Do a loop
	// 3. Check if the value of i index is the same as the i+1 index
	// 4. If yes, result += 1
	// 5. If no, continue loop.

	sort.Ints(ar)

	for i := 0; i < n-1; i++ {
		if ar[i] == ar[i+1] {
			result += 1
			i++
		}
	}

	return result
}
