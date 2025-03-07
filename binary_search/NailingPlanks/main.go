package main

import (
	"fmt"
	"math"
	"sort"
)

func findEarliestNail(nails [][]int, A, B int) int {
	left, right := 0, len(nails)-1
	minIndex := math.MaxInt
	result := -1

	for left <= right {
		mid := (left + right) / 2
		//Found match
		if nails[mid][0] >= A && nails[mid][0] <= B {
			minIndex = min(minIndex, nails[mid][1])
			result = mid
			//search for more smaller index that can satisfy the requirement
			right = mid - 1
		} else if nails[mid][0] < A {
			//Now if nail is very smaller than mid we need to search right of mid
			left = mid + 1
		} else {
			//Now if nail is too big than mid we need to search left of mid
			right = mid - 1
		}
	}
	// Update if there is any other nail in the range with lesser index that we missed
	if result > -1 {
		for i := result + 1; i < len(nails); i++ {
			if nails[i][0] <= B {
				minIndex = min(nails[i][1], minIndex)
			} else {
				return minIndex
			}
		}
	}
	return minIndex
}

func minNailsRequired(A []int, B []int, C []int) int {
	N, M := len(A), len(C)

	// Step 1: Store nails with their original indices and sort by position
	nails := make([][]int, M)
	for i := 0; i < M; i++ {
		nails[i] = []int{C[i], i} // Store nail position and its original index
	}
	// fmt.Println(nails)
	sort.Slice(nails, func(i, j int) bool { return nails[i][0] < nails[j][0] })
	// fmt.Println(nails)

	maxNailIndex := -1

	// Step 2: Process each plank
	for i := 0; i < N; i++ {
		earliestNail := findEarliestNail(nails, A[i], B[i])
		// fmt.Println(earliestNail, A[i], B[i])
		if earliestNail == math.MaxInt {
			return -1 // If no nail found, return -1
		}
		maxNailIndex = max(maxNailIndex, earliestNail)
	}

	return maxNailIndex + 1 // Convert index to count
}

func main() {
	A := []int{1, 4, 5, 8}
	B := []int{4, 5, 9, 10}
	C := []int{4, 6, 7, 10, 2}
	fmt.Println(minNailsRequired(A, B, C))
}
