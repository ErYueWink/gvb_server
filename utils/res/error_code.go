package res

// ErrorCode 自定义错误类型
type ErrorCode int

const (
	SettingsError ErrorCode = 1001
	ArgumentError ErrorCode = 1002
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
		ArgumentError: "参数错误",
	}
)
