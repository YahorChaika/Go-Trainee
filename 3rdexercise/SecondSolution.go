package main

func reverseWords2(s string) string {
	b := []byte(s)
	start := -1
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			if start != -1 {
				reverseBytes(b, start, i-1)
				start = -1
			}
		} else if start == -1 {
			start = i
		}
	}
	return string(b)
}

func reverseBytes(b []byte, l, r int) {
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}
