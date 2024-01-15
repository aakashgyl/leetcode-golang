func isValid(s string) bool {
    st := stack{}
    for i:=0; i< len(s); i++ {
        char := s[i]
        switch char {
            case '{','(','[':
                st.Push(char)
            case '}',')',']':
                rchar := st.Pop()
                if rchar != '{' && char == '}' || rchar != '(' && char == ')' || rchar != '[' && char == ']' {
                    return false
            }
        }
    }
    if len(st) != 0{
        return false
    }
    return true
}

type stack []byte

func (s *stack) Push(ch byte) {
	*s = append(*s, ch)
}

func (s *stack) Pop() (ch byte) {
	if len(*s) == 0 {
		return byte(' ')
	}
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

func (s *stack) Print() {
	for i := 0; i < len(*s); i++ {
		fmt.Print(string((*s)[i]))
		fmt.Print(" -> ")
	}
}
