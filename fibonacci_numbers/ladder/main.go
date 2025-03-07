package main

import "fmt"

/*
You have to climb up a ladder. The ladder has exactly N rungs, numbered from 1 to N. With each step, you can ascend by one or two rungs. More precisely:

with your first step you can stand on rung 1 or 2,
if you are on rung K, you can move to rungs K + 1 or K + 2,
finally you have to stand on rung N.
Your task is to count the number of different ways of climbing to the top of the ladder.

For example, given N = 4, you have five different ways of climbing, ascending by:

1, 1, 1 and 1 rung,
1, 1 and 2 rungs,
1, 2 and 1 rung,
2, 1 and 1 rungs, and
2 and 2 rungs.
Given N = 5, you have eight different ways of climbing, ascending by:

1, 1, 1, 1 and 1 rung,
1, 1, 1 and 2 rungs,
1, 1, 2 and 1 rung,
1, 2, 1 and 1 rung,
1, 2 and 2 rungs,
2, 1, 1 and 1 rungs,
2, 1 and 2 rungs, and
2, 2 and 1 rung.
The number of different ways can be very large, so it is sufficient to return the result modulo 2P, for a given integer P.
*/
func createExpoofTwo() []int {
	a := []int{1}
	for i := 1; i <= 30; i++ {
		a = append(a, 2*a[i-1])
	}
	return a
}
func createFibonacci(A []int) []int {
	fibArray := []int{0, 1}
	for i := 2; i <= len(A)+1; i++ {
		fibArray = append(fibArray, fibArray[i-1]+fibArray[i-2])
	}
	return fibArray
}
func Solution(A, B []int) []int {
	powOfTwo := createExpoofTwo()
	fib := createFibonacci(A)
	result := []int{}
	for i := 0; i < len(A); i++ {
		result = append(result, fib[A[i]+1]%powOfTwo[B[i]])
	}
	return result
}
func main() {
	A := []int{4, 4, 5, 5, 1}
	B := []int{3, 2, 4, 3, 1}
	fmt.Println(Solution(A, B))

}
