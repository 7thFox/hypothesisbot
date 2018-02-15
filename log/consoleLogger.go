package log

import (
	"fmt"
	"strings"
)

type ConsoleLogger struct {
	lastReplaceable bool
	lastOptLength   int
}

func NewConsoleLogger() *ConsoleLogger {
	l := ConsoleLogger{}
	l.lastReplaceable = false
	l.lastOptLength = 0
	return &l
}

func (l *ConsoleLogger) Log(m string) error {
	if l.lastReplaceable {
		fmt.Printf("\r%s\n", m)
	} else {
		fmt.Println(m)
	}

	l.lastReplaceable = false
	l.lastOptLength = 0
	return nil
}

func (l *ConsoleLogger) clearLine() {
	fmt.Printf("\r%s", strings.Repeat(" ", l.lastOptLength))
}

func (l *ConsoleLogger) LogState(m string) error {
	l.clearLine()
	fmt.Printf("\r%s", m)

	l.lastReplaceable = true
	l.lastOptLength = len(m)
	return nil
}
