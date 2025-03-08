package main

import "fmt"

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
func Solution(A, B []int) int {
	var count int
	st := NewStack(len(A))
	for i, j := range A {
		if B[i] == 1 {
			st.Push(j)
		} else {
			if len(st.stack) == 0 {
				count++
			} else {
				var eaten bool
				for len(st.stack) != 0 {
					item := st.Pop()
					if item > j {
						eaten = true
						st.Push(item)
						break
					}
				}
				if !eaten {
					count++
				}
			}
		}
	}
	count += len(st.stack)
	return count
}

func main() {
	fmt.Println(Solution([]int{4, 3, 2, 1, 5}, []int{0, 1, 0, 0, 0}))
}
