package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("aab"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	max := 0
	uniqueMap := make(map[byte]bool)

	l := 0
	r := 1
	uniqueMap[s[l]] = true

	for r < len(s) {
		if uniqueMap[s[r]] == false {
			uniqueMap[s[r]] = true
			r++
		} else {
			currCount := len(uniqueMap)
			max = getMaxNew(max, currCount)

			for s[l] != s[r] {
				delete(uniqueMap, s[l])
				l++
			}
			l++
			r++
		}
	}

	return getMaxNew(max, len(uniqueMap))
}

func getMaxNew(a, b int) int {
	if a > b {
		return a
	}
	return b
}
