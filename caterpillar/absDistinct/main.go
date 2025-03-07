package main

import "fmt"

/*
A non-empty array A consisting of N numbers is given. The array is sorted in non-decreasing order. The absolute distinct count of this array is the number of distinct absolute values among the elements of the array.

For example, consider array A such that:

  A[0] = -5
  A[1] = -3
  A[2] = -1
  A[3] =  0
  A[4] =  3
  A[5] =  6
The absolute distinct count of this array is 5, because there are 5 distinct absolute values among the elements of this array, namely 0, 1, 3, 5 and 6.
*/

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func Solution(A []int) int {
	hashMap := map[int]bool{}
	for _, j := range A {
		hashMap[abs(j)] = true
	}
	return len(hashMap)
}

func main() {
	A := []int{-5, -3, -1, 0, 3, 6}
	fmt.Println(Solution(A))
}
