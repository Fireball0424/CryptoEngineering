package main

import (
	"fmt"
	"math/rand"
)

var mapping map[[4]int]int
var rev map[int][4]int
var count = 0

func generatePermutations(level int, curPerm [4]int, used [4]bool) {
	if level == 4 {
		mapping[curPerm] = count
		rev[count] = curPerm
		count++
		return
	}
	for i := 0; i < 4; i++ {
		if !used[i] {
			used[i] = true
			curPerm[level] = i + 1
			generatePermutations(level+1, curPerm, used)
			curPerm[level] = 0
			used[i] = false
		}
	}
}
func Naive() int {
	card := [4]int{1, 2, 3, 4}
	for i := 0; i < len(card); i++ {
		n := rand.Intn(len(card))
		card[i], card[n] = card[n], card[i]
	}
	return mapping[card]
}

func Knuth() int {
	card := [4]int{1, 2, 3, 4}
	for i := len(card) - 1; i > 0; i-- {
		n := rand.Intn(i + 1)
		card[i], card[n] = card[n], card[i]
	}
	return mapping[card]
}

func calculateVariance(static []int) float64 {
	var mean float64 = 0.0
	var variance float64 = 0.0
	for i := 0; i < len(static); i++ {
		mean += float64(static[i])
	}
	mean /= float64(len(static))
	for i := 0; i < len(static); i++ {
		variance += (float64(static[i]) - mean) * (float64(static[i]) - mean)
	}
	variance /= float64(len(static))
	return variance
}

func main() {
	// build mapping
	mapping = make(map[[4]int]int)
	rev = make(map[int][4]int)
	generatePermutations(0, [4]int{}, [4]bool{false, false, false, false})

	// static

	var staticNaive [24]int
	var staticKnuth [24]int

	const RunTime = 1000000

	for i := 0; i < RunTime; i++ {
		staticNaive[Naive()]++
	}
	for i := 0; i < RunTime; i++ {
		staticKnuth[Knuth()]++
	}

	// output
	println("Naive algorithm: ")
	for i := 0; i < len(staticNaive); i++ {
		fmt.Printf("[%d %d %d %d]: %d \n", rev[i][0], rev[i][1], rev[i][2], rev[i][3], staticNaive[i])
	}
	println("Fisher-Yates shuffle:")
	for i := 0; i < len(staticKnuth); i++ {
		fmt.Printf("[%d %d %d %d]: %d \n", rev[i][0], rev[i][1], rev[i][2], rev[i][3], staticKnuth[i])
	}

	// calculate standard deviation
	varianceNaive := calculateVariance(staticNaive[:])
	varianceKnuth := calculateVariance(staticKnuth[:])
	fmt.Printf("Naive algorithm variance = %.2f\n", varianceNaive)
	fmt.Printf("Fisher-Yates shuffle variance = %.2f\n", varianceKnuth)
}
