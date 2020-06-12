package chapter16

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/badgerodon/collections/stack"
	"github.com/labstack/gommon/log"
)

type operator string

var (
	operatorSubtract = operator("-")
	operatorAdd      = operator("+")
	operatorMultiply = operator("*")
	operatorDivide   = operator("/")
	operatorBlank    = operator(" ")
)

type calculator struct{}

func newCalculator() *calculator {
	return &calculator{}
}

// Compute returns the result of computing a string
func (c calculator) Compute(sequence string) (result float64) {
	numberStack := stack.New()
	operatorStack := stack.New()

	for i := 0; i < len(sequence); i++ {
		value := c.parseNextNumber(sequence, i)
		numberStack.Push(float64(value))

		// move to the operator

		i += len(strconv.Itoa(value))
		if i >= len(sequence) {
			break
		}

		// get operator, collapse top as needed, push operator
		op := c.parseNextOperator(sequence, i)
		c.collapseTop(op, numberStack, operatorStack)
		operatorStack.Push(op)
	}

	// fo final collapse
	c.collapseTop(operatorBlank, numberStack, operatorStack)
	if numberStack.Len() == 1 && operatorStack.Len() == 0 {
		return numberStack.Pop().(float64)
	}

	return result
}

func (c calculator) collapseTop(futureTop operator, numberStack *stack.Stack, operatorStack *stack.Stack) {
	for operatorStack.Len() >= 1 && numberStack.Len() >= 2 {
		if c.priorityOfOperator(futureTop) <= c.priorityOfOperator(operatorStack.Peek().(operator)) {
			second := numberStack.Pop().(float64)
			first := numberStack.Pop().(float64)
			op := operatorStack.Pop().(operator)
			collapsed := c.applyOperator(first, op, second)
			numberStack.Push(collapsed)
		} else {
			break
		}
	}
}

func (c calculator) applyOperator(first float64, op operator, second float64) float64 {
	switch op {
	case operatorAdd:
		return first + second
	case operatorSubtract:
		return first - second
	case operatorMultiply:
		return first * second
	case operatorDivide:
		return first / second
	}
	return second
}
func (c calculator) priorityOfOperator(op operator) int {
	switch op {
	case operatorAdd:
		return 1
	case operatorSubtract:
		return 1
	case operatorMultiply:
		return 2
	case operatorDivide:
		return 2
	case operatorBlank:
		return 0

	}

	return 0
}

func (c calculator) parseNextOperator(sequence string, index int) operator {
	if index < len(sequence) {
		switch sequence[index] {
		case '+':
			return operatorAdd
		case '-':
			return operatorSubtract
		case '*':
			return operatorMultiply
		case '/':
			return operatorDivide
		}
	}
	return operatorBlank
}

func (c calculator) parseNextNumber(sequence string, index int) (value int) {
	var sb strings.Builder
	for index < len(sequence) && unicode.IsDigit(rune(sequence[index])) {
		sb.WriteString(string(sequence[index]))
		index++
	}

	val, err := strconv.Atoi(sb.String())
	if err != nil {
		log.Fatalf(err.Error())
	}

	return val
}
