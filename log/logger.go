package log

type Logger interface {
	Log(m string) error
	LogState(m string) error
}
