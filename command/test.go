package command

import (
	"fmt"

	"github.com/7thFox/hypothesisbot/sender"
)

// Test sends a test message
type Test struct {
}

func (t Test) Execute(s sender.Sender) error {
	fmt.Println("exec")
	return s.Say("Hello World!")
}

func NewTest() *Test {
	this := new(Test)
	return this
}
