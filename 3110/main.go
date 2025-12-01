func scoreOfString(s string) int {
	var score int
	for i := 0; i < len(s)-1; i++ {
		difference := int(s[i]) - int(s[i+1])
		if difference < 0 {
			difference = -difference
		}
		score += difference
	}
	return score
}
