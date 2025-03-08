package main

import "fmt"

/*
A string S consisting of N characters is considered to be properly nested if any of the following conditions is true:

S is empty;
S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
S has the form "VW" where V and W are properly nested strings.
For example, the string "{[()()]}" is properly nested but "([)()]" is not.

Write a function:

class Solution { public int solution(String S); }

that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.

For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]", the function should return 0, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..200,000];
string S is made only of the following characters: '(', '{', '[', ']', '}' and/or ')'.
*/
type Stack struct {
	stack []string
	size  int
}

func (s *Stack) Push(i byte) {
	s.stack = append(s.stack, string(i))
}
func NewStack(size int) *Stack {
	return &Stack{
		stack: make([]string, size),
		size:  size,
	}
}
func (s *Stack) Pop() string {
	if len(s.stack) > 0 {
		item := s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
		return item
	}
	return ""
}
func Solution(s string) int {
	st := NewStack(len(s))
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "(":
			st.Push(s[i])
		case "{":
			st.Push(s[i])
		case "[":
			st.Push(s[i])
		case "]":
			if st.Pop() != "[" {
				return 0
			}
		case "}":
			if st.Pop() != "{" {
				return 0
			}
		case ")":
			if st.Pop() != "(" {
				return 0
			}
		default:
			return 0
		}
	}
	return 1
}

func main() {
	fmt.Println(Solution("{[()()]}"))
	fmt.Println(Solution("([)()]"))
}
