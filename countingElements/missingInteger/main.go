package main

import (
	"fmt"
	"sort"
)

/*
This is a demo task.

Write a function:

class Solution { public int solution(int[] A); }

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].
*/
func Solution(a []int) int {
	lowestPositiveMissingInteger := 1
	sort.Ints(a)
	for i := 0; i < len(a)-1; i++ {
		if a[i] == 1 {
			lowestPositiveMissingInteger = -1
		}
		if a[i]+1 != a[i+1] && a[i] != a[i+1] {
			if a[i] > 0 {
				lowestPositiveMissingInteger = a[i] + 1
			}
			break
		}

	}
	if lowestPositiveMissingInteger == -1 {
		lowestPositiveMissingInteger = len(a) + 1
	}
	return lowestPositiveMissingInteger
}

func main() {
	fmt.Println(Solution([]int{1, 3, 6, 4, 1, 2}))
	fmt.Println(Solution([]int{1, 2, 3}))
	fmt.Println(Solution([]int{-1, -3}))

}
