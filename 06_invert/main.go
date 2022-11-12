package invert

func Run(sentence string) string {
	var inverted string
	for i := len(sentence) - 1; i >= 0; i-- {
		inverted += string(sentence[i])
	}
	return inverted
}
