package main

import (
	"fmt"
)

// Tony is interface
type Tony interface {
	Ifeoma() string
}

type together struct {
	text string
}

// New must return an interface of Tony 
// That can only happen if the struct (together) as an implementation of Ifeoma
// also not you mustn't use text:text in a struct you can just pass in the value
func New(text string) Tony {
	return &together{text}
}

// Ifeoma is a method
func (t *together) Ifeoma() string{
	return t.text
}
func (t *together) String() string{
	return fmt.Sprintf("%v", t.text)
}
func main() {
	holder := New("Loves each other")
	fmt.Println(holder)
	holder2 := &together{"another way"}
	factor(holder2)
}

func factor(t Tony) {
	fmt.Println(t.Ifeoma())
}