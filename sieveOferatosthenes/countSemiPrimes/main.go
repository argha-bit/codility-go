package main

import "fmt"

func getPrimeNumbersInARange(N int) []int {
	numbers := make([]bool, N+1)
	for i := range numbers {
		numbers[i] = true
	}
	numbers[0], numbers[1] = false, false
	// eratosthenes algorithm
	i := 2
	for i*i < N {
		if numbers[i] {
			k := i * i
			for k <= N {
				numbers[k] = false
				k += i
			}
		}
		i += 1
	}
	primeNumbers := []int{}
	for i, j := range numbers {
		if j {
			primeNumbers = append(primeNumbers, i)
		}
	}

	return primeNumbers
}
func getPrimeFactors(N int) []int {
	primeFactors := make([]int, N+1)

	i := 2
	for i*i < N {
		//if this is a prime factor mark all of its multiples denoting it is the smallest prime number
		if primeFactors[i] == 0 {
			k := i * i
			for k <= N {
				primeFactors[k] = i
				k += 2
			}
		}
		i++
	}
	return primeFactors
}
func Solution(N int, A, B []int) []int {
	primes := getPrimeNumbersInARange(N)
	fmt.Println("primes in range", primes)
	//calculate semiPrimes
	semiPrimerange := make([]int, N+1)
	for i := 0; i < len(primes); i++ {
		for j := 0; j < len(primes); j++ {
			//we need to find semi primes only in given range
			if primes[i]*primes[j] <= N {
				semiPrimerange[primes[i]*primes[j]] = 1
			} else {
				break
			}
		}
	}
	//create cumulative sum of semiprimes till a range this will help us understand how many numbers are present till a range
	cumulative := make([]int, N+1)
	for i := 1; i < len(cumulative); i++ {
		cumulative[i] += cumulative[i-1]
		if semiPrimerange[i] == 1 {
			cumulative[i]++
		}
	}

	fmt.Println("semiPrimeRange", semiPrimerange)
	fmt.Println("cumulative", cumulative)
	returnArray := make([]int, len(A))
	for i := range A {
		returnArray[i] = cumulative[B[i]] - cumulative[A[i]-1]
	}
	return returnArray
}
func main() {
	fmt.Println(Solution(26, []int{1, 4, 16}, []int{26, 10, 20}))
}
