func findWordsContaining(words []string, x byte) []int {
	var s []int
	for i := 0; i < len(words); i++ {
		currentWord := words[i]
		for j := 0; j < len(currentWord); j++ {
			if currentWord[j] == x {
				s = append(s, i)
				break
			}
		}
	}
	return s
}
