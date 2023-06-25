package Library

import (
	"fmt"
	"log"
	"time"
)

type CusError struct {
	Code string
	Msg  string
	Time time.Time
}

func (err CusError) Error() string {
	return fmt.Sprintf("Code: %s, Msg: %s, Time: %s\n", err.Code, err.Msg, err.Time.Format("2006-01-02 15:04:05"))
}
func getCusError(code, msg string) error {
	return CusError{Code: code, Msg: msg, Time: time.Now()}
}
func ErrorsCase() {
	a, b := -1, -2
	res, err := sum(a, b)
	log.Println(res, err)
	if err != nil {
		if cusErr, ok := err.(CusError); ok {
			fmt.Printf("打印自定义错误, Code: %s, Msg: %s, Time: %s\n", cusErr.Code, cusErr.Msg, cusErr.Time.Format("2006-01-02 15:04:05"))
		}
	}
}
func sum(a, b int) (int, error) {
	if a <= 0 && b <= 0 {
		return 0, getCusError("500", "不允许数字小于或等于0")
	}
	return a + b, nil
}
