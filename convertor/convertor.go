package convertor

func isOperator(ch rune) bool {
	switch ch {
	case '+', '-', '*', '/':
		return true
	}
	return false
}

func isParathesis(ch rune) bool {
	return ch == '(' || ch == ')'
}

func isOperand(ch rune) bool {
	return !(isOperator(ch) || isParathesis(ch))
}
