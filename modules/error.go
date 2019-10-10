package modules

const (
	SUCCESS = 200
	ERROR = 500
	INVALID_PARAMS = 400
)

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
}
// 定义一个Error接口
type Error interface {
	error
	Status() int
}

//定义异常信息
type StatusError struct {
	Code int
	Err error
}

func (s StatusError) Error() string{
	return s.Err.Error()
}
func (s StatusError) Status() int{
	return s.Code
}


