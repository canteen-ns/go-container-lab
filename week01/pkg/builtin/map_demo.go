package builtin

func CountChars(s string) map[byte]int {
	m := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	return m
}
