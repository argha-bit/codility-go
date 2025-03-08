package main

import "fmt"

/*
You are going to build a stone wall. The wall should be straight and N meters long, and its thickness should be constant; however, it should have different heights in different places. The height of the wall is specified by an array H of N positive integers. H[I] is the height of the wall from I to I+1 meters to the right of its left end. In particular, H[0] is the height of the wall's left end and H[N−1] is the height of the wall's right end.

The wall should be built of cuboid stone blocks (that is, all sides of such blocks are rectangular). Your task is to compute the minimum number of blocks needed to build the wall.

Write a function:

class Solution { public int solution(int[] H); }

that, given an array H of N positive integers specifying the height of the wall, returns the minimum number of blocks needed to build it.

For example, given array H containing N = 9 integers:

  H[0] = 8    H[1] = 8    H[2] = 5
  H[3] = 7    H[4] = 9    H[5] = 8
  H[6] = 7    H[7] = 4    H[8] = 8
the function should return 7. The figure shows one possible arrangement of seven blocks.
*/
//Idea is very simple, if you see two stones of same size then u dont need a new stone ,
//if you see a stone that is bigger than the last stone you have added then u need to add a new stone complete the step
//if you see a stone whose height is less then you then u definitely can keep you current stone so you need remove the older stone until you get a stone lesser than it or you have no more stone and add count

type Stack struct {
	stack []int
	size  int
}

func (s *Stack) Push(i int) {
	s.stack = append(s.stack, i)
}
func NewStack(size int) *Stack {
	return &Stack{
		stack: []int{},
		size:  size,
	}
}
func (s *Stack) Pop() int {
	if len(s.stack) > 0 {
		item := s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
		return item
	}
	return 0
}

func Solution(A []int) int {
	s := NewStack(len(A))
	s.Push(A[0])
	count := 1
	for i := 1; i < len(A); i++ {
		newBlock := false
		for len(s.stack) != 0 {
			item := s.Pop()
			if item > A[i] {
				newBlock = true
				continue
			} else if item == A[i] {
				newBlock = false
				s.Push(item)
				s.Push(A[i])
				break
			} else {
				s.Push(item)
				s.Push(A[i])
				newBlock = true
				break
			}
		}
		if len(s.stack) == 0 {
			s.Push(A[i])
			newBlock = true
		}
		if newBlock {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(Solution([]int{8, 8, 5, 7, 9, 8, 7, 4, 8}))
}
