package main

import (
	"fmt"
	"math"
)

/*
You are given integers K, M and a non-empty array A consisting of N integers. Every element of the array is not greater than M.

You should divide this array into K blocks of consecutive elements. The size of the block is any integer between 0 and N. Every element of the array should belong to some block.

The sum of the block from X to Y equals A[X] + A[X + 1] + ... + A[Y]. The sum of empty block equals 0.

The large sum is the maximal sum of any block.

For example, you are given integers K = 3, M = 5 and array A such that:

	A[0] = 2
	A[1] = 1
	A[2] = 5
	A[3] = 1
	A[4] = 2
	A[5] = 2
	A[6] = 2

The array can be divided, for example, into the following blocks:

[2, 1, 5, 1, 2, 2, 2], [], [] with a large sum of 15;
[2], [1, 5, 1, 2], [2, 2] with a large sum of 9;
[2, 1, 5], [], [1, 2, 2, 2] with a large sum of 8;
[2, 1], [5, 1], [2, 2, 2] with a large sum of 6.// this is wrong we can make [1,2,2],[1,2,2],[5]
The goal is to minimize the large sum. In the above example, 6 is the minimal large sum.
*/
func minimalLargeSum(a []int, K int) int {
	//as we are trying to partion the array where least amount can be one element in to the array to all the element in the array
	//the sum will range between the largest element in the array to the sum of the total array
	//so the logic we have to build is the value of the minimal large sum will be between the largest element of the array and sum of all elements of the array
	//If we can split the array pertaining to the sum produced in K number of blocks then that is a probable value
	// sort.Ints(a)
	var left, right int
	for _, j := range a {
		left = max(left, j)
		right += j
	}
	sum := math.MaxInt
	for left <= right {
		mid := (left + right) / 2
		if canSplit(a, mid, K) {
			fmt.Println("mid is ", mid)
			sum = min(sum, mid)
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	if sum == math.MaxInt {
		return 0
	}
	return sum

}
func canSplit(a []int, mid, K int) bool {
	var sum int
	blocks := 1

	for i := 0; i < len(a); {
		sum += a[i]
		if sum > mid {
			blocks++
			sum = 0
		} else {
			i++
		}
		if blocks > K {
			return false
		}
	}
	return true
}
func main() {
	A := []int{2, 1, 5, 1, 2, 2, 2}
	fmt.Println(minimalLargeSum(A, 3))
}
