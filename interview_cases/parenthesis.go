package interview_cases

//https://leetcode.com/problems/valid-parentheses/description/
//Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
//determine if the input string is valid.
//An input string is valid if:
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
//Every close bracket has a corresponding open bracket of the same type.

var squares = map[rune]rune{
	'[': ']',
	'(': ')',
	'{': '}',
}

func isValidParentheses(str string) bool {
	var stack = make([]rune, 0)
	chars := []rune(str)
	for i := 0; i < len(chars); i++ {
		switch chars[i] {
		case '[', '(', '{':
			stack = append(stack, chars[i])
		default:
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if chars[i] != squares[top] {
				return false
			}
		}
	}
	return len(stack) == 0
}
