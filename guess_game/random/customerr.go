package random

import (
	"time"
	"fmt"
)


// CustomError return custom Error that may happpen
// note that if you want to use the custom error struct
// within your local package it has to be exported But
// if you will only use its fields that need to be exported
// eg use in a json merchal function the struct don't need to be exported
type CustomError struct {
	Param1   string
	Param2   interface{}
	Err			 error
}

func (c *CustomError)	Error() string {
	return fmt.Sprintf("%v Erorr --%v  %v %v\n", time.Now(), c.Param1, c.Param2, c.Err)
}