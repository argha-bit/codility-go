package main

import (
	"fmt"
)

//the way to count factors are loop from 1 to sqrt of n and if n%i==0 then i is a factor if i and N/i is different then there are two factors if i^2=N then only one factor

func Solution(N int) int {
	if N == 1 {
		return 1
	}
	var factorCount int
	for i := 1; i*i < N; i++ {
		if N%i == 0 {
			//there is a factor
			if i == N/i {
				factorCount++
				fmt.Print(i, " ")
			} else {
				factorCount += 2
				fmt.Print(i, " ", N/i, " ")
			}
		}
	}
	return factorCount
}

func main() {
	fmt.Println(Solution(24))
}
