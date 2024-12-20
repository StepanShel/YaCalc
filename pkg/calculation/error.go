package calculation

import "errors"

var (
	InvalidReqBody     = errors.New("Invalid Request Body")
	DivByZero          = errors.New("Division by zero")
	UnsupportedOp      = errors.New("Unsupported operator")
	InvalExpresInBrack = errors.New("invalid expression in brackets")
	MissBracket        = errors.New("'(' is missing")
	InvalExp           = errors.New("invalid expression")
)
