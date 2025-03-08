package main

import "fmt"

/*
	A non-empty array A consisting of N integers is given.

A triplet (X, Y, Z), such that 0 ≤ X < Y < Z < N, is called a double slice.

The sum of double slice (X, Y, Z) is the total of A[X + 1] + A[X + 2] + ... + A[Y − 1] + A[Y + 1] + A[Y + 2] + ... + A[Z − 1].

For example, array A such that:

    A[0] = 3
    A[1] = 2
    A[2] = 6
    A[3] = -1
    A[4] = 4
    A[5] = 5
    A[6] = -1
    A[7] = 2
contains the following example double slices:

double slice (0, 3, 6), sum is 2 + 6 + 4 + 5 = 17,
double slice (0, 3, 7), sum is 2 + 6 + 4 + 5 − 1 = 16,
double slice (3, 4, 5), sum is 0.
The goal is to find the maximal sum of any double slice.

Write a function:

class Solution { public int solution(int[] A); }

that, given a non-empty array A consisting of N integers, returns the maximal sum of any double slice.

For example, given:

    A[0] = 3
    A[1] = 2
    A[2] = 6
    A[3] = -1
    A[4] = 4
    A[5] = 5
    A[6] = -1
    A[7] = 2
the function should return 17, because no double slice of array A has a sum of greater than 17.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [3..100,000];
each element of array A is an integer within the range [−10,000..10,000].
*/

//AnaLysis
// 3,2,6,-1,4,5,-1,2

//The Understanding behind this is sum of all elements in Xto Z with out the value of X and also y is a index between Xto Z which is not taken
//So the possible slices are 3,2,6,-1,4,5,-1,2 (0,2,7)is 2+(-1+4+5+-1)
// the logic we will apply is kadanes
//if we start from left our max sum will look like
//[0,2,8,7,11,16,15] from left to right
//[16,14,8,9,5,0,0] from right to left
// Now notice any time if we take y =2 then ,max sum will be 10
func Solution(A []int) int {
	N := len(A)
	LR := make([]int, len(A))
	RL := make([]int, len(A))
	// LR[0] = 0
	// RL[len(A)-1] = 0
	sum := 0
	for i := 1; i < N-1; i++ {
		sum += A[i]
		if sum < 0 {
			sum = 0
		}
		LR[i] = sum
	}
	sum = 0
	for i := N - 2; i >= 1; i-- {
		sum += A[i]
		if sum < 0 {
			sum = 0
		}
		RL[i] = sum
	}
	var maxSum int
	for i := 1; i < N-2; i++ {
		maxSum = max(maxSum, LR[i]+RL[i+2])
	}
	fmt.Println(maxSum)
	return maxSum
}
func main() {
	Solution([]int{3, 2, 6, -1, 4, 5, -1, 2})
}
