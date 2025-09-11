package main


func reverseWords3(s string) string {
	result := []rune{}
	word := []rune{}

	for _, r := range s { 
		if r == ' ' {
			for j := len(word) - 1; j >= 0; j-- {
				result = append(result, word[j])
			}
			word = word[:0] 
			result = append(result, ' ')
		} else {
			word = append(word, r)
		}
	}
	
	for j := len(word) - 1; j >= 0; j-- {
		result = append(result, word[j])
	}
	return string(result)
}
