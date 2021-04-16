package main

func main() {
	smallestSubsquence("cbacdcbc")
}

func smallestSubsequence(s string) string {

	var left [26]int
	for i := range s {
		left[s[i]-'a']++
	}

	var stack []byte
	var instack [26]bool

	for i := range s {
		c := s[i]
		if !instack[c-'a'] {

		}

		left[c-a]--
	}

	return string(stack)
}
