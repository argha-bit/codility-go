package main

import (
	"fmt"
	"math"
	"sort"
)

/*
Let A be a non-empty array consisting of N integers.

The abs sum of two for a pair of indices (P, Q) is the absolute value |A[P] + A[Q]|, for 0 ≤ P ≤ Q < N.

For example, the following array A:

	A[0] =  1
	A[1] =  4
	A[2] = -3

has pairs of indices (0, 0), (0, 1), (0, 2), (1, 1), (1, 2), (2, 2).
The abs sum of two for the pair (0, 0) is A[0] + A[0] = |1 + 1| = 2.
The abs sum of two for the pair (0, 1) is A[0] + A[1] = |1 + 4| = 5.
The abs sum of two for the pair (0, 2) is A[0] + A[2] = |1 + (−3)| = 2.
The abs sum of two for the pair (1, 1) is A[1] + A[1] = |4 + 4| = 8.
The abs sum of two for the pair (1, 2) is A[1] + A[2] = |4 + (−3)| = 1.
The abs sum of two for the pair (2, 2) is A[2] + A[2] = |(−3) + (−3)| = 6.

Write a function:

func Solution(A []int) int

that, given a non-empty array A consisting of N integers, returns the minimal abs sum of two for any pair of indices in this array.

For example, given the following array A:

	A[0] =  1
	A[1] =  4
	A[2] = -3

the function should return 1, as explained above.

Given array A:

	A[0] = -8
	A[1] =  4
	A[2] =  5
	A[3] =-10
	A[4] =  3

the function should return |(−8) + 5| = 3.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000,000..1,000,000,000].
*/
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Solution(A []int) int {
	//The Notion here is we first sort the array now the minimum positive integer is 1 so if our value is too large then we reduce right and if it is too less we increase left
	sort.Ints(A)
	var minPositiveInteger, left, right int
	right = len(A) - 1
	minPositiveInteger = math.MaxInt
	for left <= right {
		sum := abs(A[left] + A[right])
		if sum == 1 {
			return sum
		} else if sum > 0 {
			left++
		} else {
			right--
		}
		minPositiveInteger = min(sum, minPositiveInteger)
	}
	return minPositiveInteger
}
func main() {
	A := []int{-8, 4, 5, -10, 3}
	fmt.Println(Solution(A))
	fmt.Println(Solution([]int{1, 4, -3}))
}
