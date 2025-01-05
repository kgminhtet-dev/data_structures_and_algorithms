package convertor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsOperator(t *testing.T) {
	testcases := []struct {
		input    rune
		expected bool
	}{
		{'+', true},
		{'-', true},
		{'*', true},
		{'/', true},
		{'(', false},
		{')', false},
	}

	for _, tc := range testcases {
		result := isOperator(tc.input)

		assert.Equal(t, tc.expected, result)
	}
}

func TestIsOperand(t *testing.T) {
	testcases := []struct {
		input    rune
		expected bool
	}{
		{'a', true},
		{'b', true},
		{'*', false},
		{'+', false},
		{'(', false},
	}

	for _, tc := range testcases {
		result := isOperand(tc.input)
		assert.Equal(t, tc.expected, result)
	}
}

func TestArrayStack(t *testing.T) {
	t.Run("Push", func(t *testing.T) {
		stack := NewArrayStack[int]()
		stack.Push(10)

		currentValue, _ := stack.Peek()
		assert.Equal(t, 10, *currentValue)
		assert.Equal(t, 1, stack.Size())
	})

	t.Run("Pop", func(t *testing.T) {
		stack := NewArrayStack[int]()
		stack.Push(10)
		stack.Push(20)
		stack.Push(30)
		assert.Equal(t, 3, stack.size)
		elem, _ := stack.Pop()

		assert.Equal(t, 30, *elem)
		assert.Equal(t, 2, stack.size)
	})

}

func TestToPostfixConvertor(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"a+b",
			"a+b",
			"ab+",
		},
		{
			"a*b",
			"a*b",
			"ab*",
		},
		{
			"a+b-c",
			"a+b-c",
			"ab+c-",
		},
		{
			"a*b+c",
			"a*b+c",
			"ab*c+",
		},
		{
			"a+b*c",
			"a+b*c",
			"abc*+",
		},
		{
			"a+b-c+d",
			"a+b-c+d",
			"ab+c-d+",
		},
		{
			"a*b+c*d",
			"a*b+c*d",
			"ab*cd*+",
		},
		{
			"a/b+c*d",
			"a/b+c*d",
			"ab/cd*+",
		},
		{
			"(a+b)*c",
			"(a+b)*c",
			"ab+c*",
		},
		{
			"a*(b+c)",
			"a*(b+c)",
			"abc+*",
		},
		{
			"(a+b)*(c-d)",
			"(a+b)*(c-d)",
			"ab+cd-*",
		},
		{
			"(a*b)/(c+d)",
			"(a*b)/(c+d)",
			"ab*cd+/",
		},
	}

	t.Parallel()
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := ToPosfixConvertor(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
