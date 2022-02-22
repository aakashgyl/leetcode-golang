func isPalindrome(s string) bool {
    strLen := len(s)
    start := 0
    end := strLen -1
    
    s = strings.ToLower(s)
    
    for start < end {
        fmt.Println(start, end)
        if !isAlphaNum(s[start]) {
            start++
            continue
        }
        
        if !isAlphaNum(s[end]) {
            end--
            continue
        } 
        
        if s[start] != s[end] {
            fmt.Println(s[start], s[end])
            return false
        } else {
            start++
            end--
        }
    }
    return true
}

func isAlphaNum(char byte) bool {
    if char >= 48 && char <= 57 || char >= 65 && char <= 90 || char >= 97 && char <= 122 {
        return true
    }
    
    // if char >= '0' && char <= '9' || char >= 'a' && char >= 'z' || char >= 'A' && char <= 'Z' {
    //     return true
    // } 
    return false
}