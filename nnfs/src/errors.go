package nnfs

import "fmt"

var (
	ERROR_WRONG_PARAMETERS error = fmt.Errorf("Wrong parameters")
	ERROR_EMPTY_PARAMETERS error = fmt.Errorf("Empty parameters")
)
