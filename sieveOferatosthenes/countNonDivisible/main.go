package main

import (
	"fmt"
	"math"
)

func getFactors(N int) []int {
	factors := []int{}
	if N == 1 {
		return []int{N}
	}
	for i := 1; i*i < N; i++ {
		if N%i == 0 {
			if i != N/i {
				factors = append(factors, i, N/i)
			} else {
				factors = append(factors, i)
			}
		}
	}
	return factors
}
func Solution(A []int) []int {
	//lets first get the highest Number in this array
	N := len(A)
	returnArray := make([]int, N)
	maxNumber := math.MinInt
	for _, j := range A {
		maxNumber = max(maxNumber, j)
	}
	//letss find out occurences of each number in the array
	occurences := make([]int, maxNumber+1)
	for _, j := range A {
		occurences[j]++
	}
	//lets count occurences of factors which if we subtract from number of elements we will have non divisors
	for i, j := range A {
		factors := getFactors(j)
		countOfDivisibles := 0
		for _, k := range factors {
			countOfDivisibles += occurences[k]
		}
		fmt.Println("Factors for ", j, ":", factors, " and divisibles are ", countOfDivisibles)
		nonDivisible := N - countOfDivisibles
		returnArray[i] = nonDivisible
	}
	return returnArray
}

func main() {
	fmt.Println(Solution([]int{3, 1, 2, 3, 6}))
}
