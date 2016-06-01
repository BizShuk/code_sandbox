package backtracking

func IsValid_Parenthese(s string) bool {
	stack := []byte{}
	validmap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		_s := s[i]
		if _s == '{' || _s == '[' || _s == '(' {
			stack = append(stack, _s)
		}

		if _s == '}' || _s == ']' || _s == ')' {
			if len(stack) == 0 || validmap[_s] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
