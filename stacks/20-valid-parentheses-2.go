package main

import (
	"errors"
	"fmt"
)

var closeToOpenMap = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

type Stack2 []byte

func (s *Stack2) Pop() (byte, error) {
	if len(*s) == 0 {
		return ' ', errors.New("stack underflow")
	}

	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val, nil
}

func (s *Stack2) Push(val byte) {
	*s = append(*s, val)
}

func isValid2(s string) bool {
	stack := Stack2{}

	for i := 0; i < len(s); i++ {
		if val, ok := closeToOpenMap[s[i]]; ok { // current char is close symbol
			popVal, err := stack.Pop()
			if err != nil || val != popVal {
				return false
			}
		} else { // current char is open symbol
			stack.Push(s[i])
		}
	}

	// stack is not empty
	if len(stack) != 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(isValid2("()"))
}
