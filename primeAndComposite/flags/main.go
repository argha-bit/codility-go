package main

import "fmt"

/*A non-empty array A consisting of N integers is given.

A peak is an array element which is larger than its neighbours. More precisely, it is an index P such that 0 < P < N − 1 and A[P − 1] < A[P] > A[P + 1].

For example, the following array A:

    A[0] = 1
    A[1] = 5
    A[2] = 3
    A[3] = 4
    A[4] = 3
    A[5] = 4
    A[6] = 1
    A[7] = 2
    A[8] = 3
    A[9] = 4
    A[10] = 6
    A[11] = 2
has exactly four peaks: elements 1, 3, 5 and 10.

You are going on a trip to a range of mountains whose relative heights are represented by array A, as shown in a figure below. You have to choose how many flags you should take with you. The goal is to set the maximum number of flags on the peaks, according to certain rules.

*/
//logic is first find the indices of peaks , then set flag on the first one and
func getPeaks(A []int) []int {
	peaks := []int{}
	N := len(A)
	for i := 1; i < N-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}
	return peaks
}

func Solution(A []int) int {
	peaks := getPeaks(A)
	//we will say we can put flags on all the peaks
	flags := len(peaks)
	if flags == 0 {
		return 0
	}
	//Now we validate by n=binary search
	minPeaks, maxpeaks := 1, flags
	for minPeaks <= maxpeaks {
		mid := (minPeaks + maxpeaks) / 2
		placed := 1
		flagsPlaces := peaks[0]
		for i := 1; i < len(peaks); i++ {
			if peaks[i] > flagsPlaces+mid {
				placed++
			}
			if placed == mid {
				break
			}
		}
		if placed == mid {
			flags = mid
			//move right
			minPeaks = mid + 1
		} else {
			//move left
			maxpeaks = mid - 1
		}
	}
	return flags
}

func main() {
	fmt.Println(Solution([]int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}))
}
