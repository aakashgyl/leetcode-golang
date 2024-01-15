package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	srcMap := make(map[rune]int)
	for _, char := range s {
		srcMap[char] = srcMap[char] + 1 // srcMap[s] will be 0 if its a new char
	}

	targetMap := make(map[rune]int)
	for _, char := range t {
		if srcMap[char] == 0 { // found char present in t but not in s
			return false
		}

		charCount := targetMap[char] + 1
		if charCount > srcMap[char] { // occurence in t is more than in s
			return false
		}
		targetMap[char] = targetMap[char] + 1
	}
	return true
}
