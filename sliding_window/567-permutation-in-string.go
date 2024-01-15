package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Map := make(map[byte]int)
	s2Map := make(map[byte]int)

	// create map of s1
	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]] = s1Map[s1[i]] + 1
	}

	// create map of s2 (element count should be same as s1)
	for i := 0; i < len(s1); i++ {
		s2Map[s2[i]] = s2Map[s2[i]] + 1
	}

	for i := 0; i < len(s2)-len(s1)+1; i++ {
		// check for substring
		isSubstr := true
		for char, _ := range s1Map {
			if s2Map[char] < s1Map[char] {
				isSubstr = false
			}
		}
		if isSubstr == true {
			return true
		}

		// update s2Map
		s2Map[s2[i]] = s2Map[s2[i]] - 1
		if (i + len(s1)) < len(s2) {
			s2Map[s2[i+len(s1)]] = s2Map[s2[i+len(s1)]] + 1
		}
	}

	return false
}
