package main

import "fmt"

func getPeakCoOrdinates(A []int) (int, []int) {
	N := len(A)
	peakCoOrdinate := make([]int, N)
	countOfPeaks := 0
	for i := 1; i < N-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peakCoOrdinate[i]++
			countOfPeaks++
		}
	}
	return countOfPeaks, peakCoOrdinate
}

func Solution(A []int) int {
	noOfPeaks, peakCoOrdinates := getPeakCoOrdinates(A)
	if noOfPeaks == 0 {
		return noOfPeaks
	}
	//initially we take we can break into max of No of peaks number of blocks
	noOfBlcoks := noOfPeaks
	fmt.Println(noOfBlcoks, peakCoOrdinates)
	for ; noOfBlcoks >= 2; noOfBlcoks-- {
		if len(A)%noOfBlcoks == 0 {
			isValid := true
			for i := 0; i < len(A); i += (len(A) / noOfBlcoks) {
				peakfound := false
				for j := i; j < i+(len(A)/noOfBlcoks); j++ {
					if peakCoOrdinates[j] == 1 {
						peakfound = true
					}
				}
				if !peakfound {
					isValid = false
					break
				}

			}
			if isValid {
				return noOfBlcoks
			}
		}
	}
	return 1
}

func main() {
	fmt.Println(Solution([]int{1, 2, 3, 4, 3, 2, 1, 2, 5, 4, 6, 2}))

}
