package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l := 0
	r := len(s) - 1

	for l < r {
		if !isAlphaNumeric(s[l]) {
			l++
			continue
		}
		if !isAlphaNumeric(s[r]) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// above function converts string to lowercase, so uppercase check is not required here
func isAlphaNumeric(char byte) bool {
	if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
		return true
	}
	return false
}
