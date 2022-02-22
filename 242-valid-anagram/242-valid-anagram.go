func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    letterCount := make([]int, 26)
    
    for i := range s {
        letterCount[s[i]-'a']++
        letterCount[t[i]-'a']--
    }
    
    for _, val := range letterCount {
        if val != 0 {
            return false
        }
    }
    
    return true
}