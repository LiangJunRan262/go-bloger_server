package enum

type LogLevelType int

const (
	LogInfoLevel LogLevelType = iota + 1
	LogWarnLevel
	LogErrLevel
)

func (l LogLevelType) String() string {
	switch l {
	case LogInfoLevel:
		return "info"
	case LogWarnLevel:
		return "warn"
	case LogErrLevel:
		return "err"
	default:
		return "unknown"
	}
}
