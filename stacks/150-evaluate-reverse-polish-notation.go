package main

import (
	"errors"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := Stack3{}
	for _, val := range tokens {
		if fn, ok := funcMap[val]; ok {
			// pop 2 values, perform operation and push back on stack
			arg1, err := stack.Pop()
			if err != nil {
				return 0
			}
			arg2, err := stack.Pop()
			if err != nil {
				return 0
			}

			arg1int, err := strconv.Atoi(arg1)
			if err != nil {
				return 0
			}

			arg2int, err := strconv.Atoi(arg2)
			if err != nil {
				return 0
			}

			output := fn(arg2int, arg1int)
			stack.Push(strconv.Itoa(output))
		} else {
			// push on stack
			stack.Push(val)
		}
	}

	result, err := stack.Pop()
	if err != nil {
		return 0
	}

	finalResult, err := strconv.Atoi(result)
	if err != nil {
		return 0
	}

	return finalResult
}

type Stack3 []string

func (s *Stack3) Pop() (string, error) {
	if len(*s) == 0 {
		return "", errors.New("stack underflow")
	}

	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val, nil
}

func (s *Stack3) Push(val string) {
	*s = append(*s, val)
}

var funcMap = map[string]func(a, b int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}
