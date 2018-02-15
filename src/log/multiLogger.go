package log

type MultiLogger struct {
	loggers []Logger
}

func NewMultiLogger() *MultiLogger {
	l := MultiLogger{}
	return &l
}

func (l *MultiLogger) Log(m string) error {
	for _, ll := range l.loggers {
		err := ll.Log(m)
		if err != nil {
			return err
		}
	}
	return nil
}
func (l *MultiLogger) LogState(m string) error {
	for _, ll := range l.loggers {
		err := ll.LogState(m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *MultiLogger) Attach(ll Logger) *MultiLogger {
	l.loggers = append(l.loggers, ll)
	return l
}
