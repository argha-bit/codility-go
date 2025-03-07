package main

import (
	"fmt"
)

/*
The Fibonacci sequence is defined using the following recursive formula:

	F(0) = 0
	F(1) = 1
	F(M) = F(M - 1) + F(M - 2) if M >= 2

A small frog wants to get to the other side of a river. The frog is initially located at one bank of the river (position −1) and wants to get to the other bank (position N). The frog can jump over any distance F(K), where F(K) is the K-th Fibonacci number. Luckily, there are many leaves on the river, and the frog can jump between the leaves, but only in the direction of the bank at position N.

The leaves on the river are represented in an array A consisting of N integers. Consecutive elements of array A represent consecutive positions from 0 to N − 1 on the river. Array A contains only 0s and/or 1s:

0 represents a position without a leaf;
1 represents a position containing a leaf.
The goal is to count the minimum number of jumps in which the frog can get to the other side of the river (from position −1 to position N). The frog can jump between positions −1 and N (the banks of the river) and every position containing a leaf.

For example, consider array A such that:

	A[0] = 0
	A[1] = 0
	A[2] = 0
	A[3] = 1
	A[4] = 1
	A[5] = 0
	A[6] = 1
	A[7] = 0
	A[8] = 0
	A[9] = 0
	A[10] = 0

The frog can make three jumps of length F(5) = 5, F(3) = 2 and F(5) = 5.
*/
func genFib(limit int) []int {

	//creating fib with 1, 2 because duplicate 1 and 0 has no sense
	fibArray := []int{1, 2}
	for i := 2; i < limit; i++ {
		fibArray = append(fibArray, fibArray[i-1]+fibArray[i-2])
	}
	return fibArray
}

func Solution(A []int) int {
	//the idea is we have to find from any point all the possible jumps to reach exactly N else we will return -1
	fibArray := genFib(len(A) + 1)
	//frog is initially at -1
	queue := []int{-1}
	var jump int
	N := len(A)
	seen := map[int]bool{}

	for len(queue) > 0 {
		pos := queue[0]
		jump++
		if len(queue) == 1 {
			queue = []int{}
		} else {
			queue = queue[1:]
		}
		for _, j := range fibArray {
			nextPos := pos + j
			if nextPos == N {
				return jump
			} else if nextPos > N {
				break
			}
			if seen[nextPos] {
				continue
			}
			if A[nextPos] == 1 {
				seen[nextPos] = true
				queue = append(queue, nextPos)
			}
		}
	}
	return -1
}
func main() {
	A := []int{0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0}
	fmt.Println(genFib(len(A) + 1))
	fmt.Println(Solution(A))

}
