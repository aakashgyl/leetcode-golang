package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	top  int
	data []byte
}

func (s *Stack) Pop() (byte, error) {
	if s.top == -1 {
		return 0, errors.New("stack underflow error")
	}

	val := s.data[s.top]
	s.top--
	s.data = s.data[:len(s.data)-1]
	return val, nil
}

func (s *Stack) Push(val byte) {
	s.top++
	s.data = append(s.data, val)
}

func InitStack() *Stack {
	s := Stack{top: -1}
	return &s
}

func isValid(s string) bool {
	stack := InitStack()
	var popVal byte
	var err error

	for i := 0; i < len(s); i++ {
		fmt.Println(i)
		switch s[i] {
		case '(', '{', '[':
			stack.Push(s[i])
		case ')', '}', ']':
			popVal, err = stack.Pop()
			if err != nil {
				return false
			}

			switch popVal {
			case '(':
				if s[i] != ')' {
					return false
				}
			case '{':
				if s[i] != '}' {
					return false
				}
			case '[':
				if s[i] != ']' {
					return false
				}
			}
		}
	}

	if stack.top != -1 {
		return false
	}

	return true
}
