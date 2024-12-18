package calculation

import (
	"errors"
	"strconv"
	"unicode"
)

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

func applyOperation(a, b float64, op rune) (float64, error) {
	switch op {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	}
	return 0, errors.New("invalid operation")
}

func Calc(expression string) (float64, error) {
	var nums []float64
	var ops []rune

	i := 0
	n := len(expression)

	for i < n {
		if unicode.IsSpace(rune(expression[i])) {
			i++
			continue
		}

		if unicode.IsDigit(rune(expression[i])) {
			start := i
			for i < n && (unicode.IsDigit(rune(expression[i])) || expression[i] == '.') {
				i++
			}
			num, err := strconv.ParseFloat(expression[start:i], 64)
			if err != nil {
				return 0, err
			}
			nums = append(nums, num)
			continue
		}

		if expression[i] == '(' {
			ops = append(ops, '(')
		} else if expression[i] == ')' {
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				if len(nums) < 2 {
					return 0, errors.New("invalid expression")
				}
				a := nums[len(nums)-2]
				b := nums[len(nums)-1]
				opsLast := ops[len(ops)-1]
				ops = ops[:len(ops)-1]
				nums = nums[:len(nums)-2]
				result, err := applyOperation(a, b, opsLast)
				if err != nil {
					return 0, err
				}
				nums = append(nums, result)
			}
			if len(ops) == 0 {
				return 0, errors.New("mismatched parentheses")
			}
			ops = ops[:len(ops)-1]
		} else {
			for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(rune(expression[i])) {
				if len(nums) < 2 {
					return 0, errors.New("invalid expression")
				}
				a := nums[len(nums)-2]
				b := nums[len(nums)-1]
				opsLast := ops[len(ops)-1]
				ops = ops[:len(ops)-1]
				nums = nums[:len(nums)-2]
				result, err := applyOperation(a, b, opsLast)
				if err != nil {
					return 0, err
				}
				nums = append(nums, result)
			}
			ops = append(ops, rune(expression[i]))
		}
		i++
	}

	for len(ops) > 0 {
		if len(nums) < 2 {
			return 0, errors.New("invalid expression")
		}
		a := nums[len(nums)-2]
		b := nums[len(nums)-1]
		opsLast := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		nums = nums[:len(nums)-2]
		result, err := applyOperation(a, b, opsLast)
		if err != nil {
			return 0, err
		}
		nums = append(nums, result)
	}

	if len(nums) != 1 {
		return 0, errors.New("invalid expression")
	}

	return nums[0], nil
}
