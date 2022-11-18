package parenthesis

import (
	"fmt"
	"regexp"
	"strings"
)

func Run(expression string) bool {
	if len(expression) == 0 {
		return true
	} else if countOpening(expression) != countClosing(expression) {
		return false
	}
	return check(trimExp(expression))
}

// each opening par
// check if next par is opening
// if true recurse
// if false check if same type of par
// if true valid
// if false invalid

func check(expression string) bool {
	var s Stack
	fmt.Println(strings.Split(expression, ""))
	for _, ch := range strings.Split(expression, "") {
		fmt.Println(ch)
		if ch == "(" || ch == "[" || ch == "{" {
			s.Push(ch)
			fmt.Printf("Added %v", ch)
		} else if ch == ")" || ch == "]" || ch == "}" {
			if s.Len() == 0 {
				return false
			}
			top := s.Pop()
			if ch == "(" && top != ")" {
				return false
			}
			if ch == "[" && top != "]" {
				return false
			}
			if ch == "{" && top != "}" {
				return false
			}
		}
	}
	return s.Len() == 0
}

func validate(opening string, rest string) bool {
	panic("not implemented")
}

func countOpening(expression string) int {
	return strings.Count(expression, "(") + strings.Count(expression, "{") + strings.Count(expression, "[")
}

func countClosing(expression string) int {
	return strings.Count(expression, ")") + strings.Count(expression, "}") + strings.Count(expression, "]")
}

func trimExp(expression string) string {
	m := regexp.MustCompile("[^})\\][{(]")
	return m.ReplaceAllString(expression, "")
}

// Stack!
type item struct {
	val  string
	next *item
}

type Stack struct {
	top  *item
	size int
}

func (stack *Stack) Len() int {
	return stack.size
}

func (stack *Stack) Push(value string) {
	stack.top = &item{
		val:  value,
		next: stack.top,
	}
	stack.size++
}

func (stack *Stack) Pop() (value string) {
	if stack.Len() > 0 {
		value = stack.top.val
		stack.top = stack.top.next
		stack.size--
		return
	}

	return ""
}
