func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    letterCount := make(map[byte]int)
    
    // process s
    for _, char := range []byte(s) {
        if count, ok := letterCount[char]; ok {
            letterCount[char] = count + 1
        } else {
            letterCount[char] = 1   
        }
    }

    // decrement count
    for _, char := range []byte(t) {
        if count, ok := letterCount[char]; ok {
            if count == 0 { // if char occurs more in t
                return false
            }
            letterCount[char] = count - 1
        } else {    // char present in t but not in s
            return false
        }
    }
    
    // all counts should be zero
    for _, count := range letterCount {
        if count != 0 {
            return false
        }
    }
    
    return true
}