package parenthesis

import "strings"

func Run(expression string) bool {
	if len(expression) == 0 {
		return true
	} else if countOpening(expression) == countClosing(expression) {
		return true
	}
	return false
}

// each opening par
// check if next par is opening
// if true recurse
// if false check if same type of par
// if true valid
// if false invalid

func validate(opening string, rest string) bool {
	panic("not implemented")
}

func countOpening(expression string) int {
	return strings.Count(expression, "(") + strings.Count(expression, "{") + strings.Count(expression, "[")
}

func countClosing(expression string) int {
	return strings.Count(expression, ")") + strings.Count(expression, "}") + strings.Count(expression, "]")
}
