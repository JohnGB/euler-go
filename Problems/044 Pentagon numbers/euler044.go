/*
Pentagonal numbers are generated by the formula, Pn=n(3n−1)/2. The first ten
pentagonal numbers are:

1, 5, 12, 22, 35, 51, 70, 92, 117, 145, ...

It can be seen that P4 + P7 = 22 + 70 = 92 = P8. However, their difference, 70
− 22 = 48, is not pentagonal.

Find the pair of pentagonal numbers, Pj and Pk, for which their sum and
difference are pentagonal and D = |Pk − Pj| is minimised; what is the value of
D?
*/

package main

import (
	"fmt"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func abs(n int) int {
	if n > 0 {
		return n
	} else {
		return -n
	}
}

func makePentagonMap(lim int) map[int]bool {

	pentMap := make(map[int]bool)
	for n := 1; n <= lim; n++ {
		pentMap[n*(3*n-1)/2] = true
	}
	return pentMap
}

func pent(n int) int {
	return n * (3*n - 1) / 2
}

func findLowestDiffPentPair() int {
	defer timeTrack(time.Now(), "findLowestDiffPentPair()")

	const LIMIT = 5000

	minDifference := 9999999999 // large number

	pentMap := makePentagonMap(LIMIT + 100)

	// k will be smaller than j
	for k := 1; k <= LIMIT; k++ {
		for j := k + 1; j <= LIMIT; j++ {
			if pentMap[pent(j)-pent(k)] == false {
				continue
			} else if pentMap[pent(k)+pent(j)] == false {
				continue
			} else if (pent(j) - pent(k)) < minDifference {
				minDifference = (pent(j) - pent(k))
			}
		}
	}
	return minDifference

}

func main() {
	// p(makePentagonMap(20))
	// p(abs(1 - 9))
	p(findLowestDiffPentPair())

}