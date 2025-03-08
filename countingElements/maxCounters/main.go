package main

import "fmt"

/*
You are given N counters, initially set to 0, and you have two possible operations on them:

increase(X) − counter X is increased by 1,
max counter − all counters are set to the maximum value of any counter.
A non-empty array A of M integers is given. This array represents consecutive operations:

if A[K] = X, such that 1 ≤ X ≤ N, then operation K is increase(X),
if A[K] = N + 1 then operation K is max counter.
For example, given integer N = 5 and array A such that:

    A[0] = 3
    A[1] = 4
    A[2] = 4
    A[3] = 6
    A[4] = 1
    A[5] = 4
    A[6] = 4
the values of the counters after each consecutive operation will be:

    (0, 0, 1, 0, 0)
    (0, 0, 1, 1, 0)
    (0, 0, 1, 2, 0)
    (2, 2, 2, 2, 2)
    (3, 2, 2, 2, 2)
    (3, 2, 2, 3, 2)
    (3, 2, 2, 4, 2)
The goal is to calculate the value of every counter after all operations.

Write a function:

class Solution { public int[] solution(int N, int[] A); }

that, given an integer N and a non-empty array A consisting of M integers, returns a sequence of integers representing the values of the counters.

Result array should be returned as an array of integers.

For example, given:

    A[0] = 3
    A[1] = 4
    A[2] = 4
    A[3] = 6
    A[4] = 1
    A[5] = 4
    A[6] = 4
the function should return [3, 2, 2, 4, 2], as explained above.

Write an efficient algorithm for the following assumptions:

N and M are integers within the range [1..100,000];
each element of array A is an integer within the range [1..N + 1].
*/
func Solution(A []int, N int) []int {
	//logic here is increment a value of A[i] to 1 considering A[i]<=N
	//if A[i]>=N+1 then all values are updated to max Value of any counter
	//How we approach this problem ? we keep track of this kind of max updates which we will update later
	//at first pass we update as per the rule as we saw all the items then in the next pass we increment any counters value that is less than lazy update by adding lazy update to that
	returnArray := make([]int, N)
	var maxCounterSoFar, lazyUpdate int
	//we are normally incrementing
	for _, j := range A {
		if j <= N {
			if j < lazyUpdate {
				returnArray[j-1] = lazyUpdate + 1
				maxCounterSoFar = max(maxCounterSoFar, returnArray[j-1])
			} else {
				returnArray[j-1]++
				maxCounterSoFar = max(maxCounterSoFar, returnArray[j-1])
			}

		} else {
			lazyUpdate = maxCounterSoFar
		}
	}
	//now we wapply lazy update because the values that are still less than my lazy updates are actually appeared before lazyupdate to be applied
	for i, j := range returnArray {
		if j < lazyUpdate {
			returnArray[i] = lazyUpdate
		}
	}
	return returnArray

}
func main() {
	fmt.Println(Solution([]int{3, 4, 4, 6, 1, 4, 4}, 5))
}
