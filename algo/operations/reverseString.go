package op

func ReverseString(s string) string {
	slen := len(s)
	var ret = make([]byte, slen)
	for i := 0; i < slen; i++ {
		ret[i] = s[slen-i-1]
	}
	return string(ret)
}

// BUG() ,No extra space
func ReverseString_nospace(s string) string {
	slen := len(s)
	for i := 0; i < slen/2; i++ {
		s[i], s[slen-1-i] = s[slen-1-i], s[i]
	}
	return s
}
