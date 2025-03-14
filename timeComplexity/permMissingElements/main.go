package main

import (
	"fmt"
	"sort"
)

/*
An array A consisting of N different integers is given. The array contains integers in the range [1..(N + 1)], which means that exactly one element is missing.

Your goal is to find that missing element.

Write a function:

class Solution { public int solution(int[] A); }

that, given an array A, returns the value of the missing element.

For example, given array A such that:

	A[0] = 2
	A[1] = 3
	A[2] = 1
	A[3] = 5

the function should return 4, as it is the missing element.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
the elements of A are all distinct;
each element of array A is an integer within the range [1..(N + 1)].
*/
func Solution(A []int) int {
	N := len(A)
	missingElem := N + 1
	sort.Ints(A)
	for i, j := range A {
		if i+1 != j {
			missingElem = i + 1
		}
	}
	return missingElem
}
func main() {
	fmt.Println(Solution([]int{2, 3, 1, 5}))
}
