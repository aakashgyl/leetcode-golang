package main

import "fmt"

type Key [26]byte

func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[Key][]string)
	var output [][]string
	for _, str := range strs {
		strKey := getKey(str)
		groupMap[strKey] = append(groupMap[strKey], str)
	}

	for _, group := range groupMap {
		output = append(output, group)
	}

	return output
}

func getKey(str string) (key Key) {
	for _, char := range str {
		key[char-'a'] = key[char-'a'] + 1
	}
	return key
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
