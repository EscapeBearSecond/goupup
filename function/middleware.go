package function

import (
	"errors"
	"log"
)

func Sum(a, b int) (int, error) {
	if a <= 0 && b <= 0 {
		err := errors.New("both numbers must be greater than 0")
		return 0, err
	}
	sum := a + b
	return sum, nil
}
type MyMiddleFun func(int, int) (int, error)
func LogMiddleWare(next MyMiddleFun) MyMiddleFun {
	return func(a, b int) (int, error) {
		log.Printf("日志中间件，a=%d, b=%d", a, b)
		return next(a, b)
	}
}
// 声明receiver为函数类型的的方法，即函数类型对象的方法
func (sum MyMiddleFun) Accumulation(list ...int) (int, error) {
	s := 0
	var err error
	for _, v := range list {
		s, err = sum(s, v)
		if err != nil  {
			return s, err
		}
	}
	return s, err
}