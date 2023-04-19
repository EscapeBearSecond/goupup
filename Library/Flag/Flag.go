package main

import (
	"flag"
	"fmt"
	"strings"
)

// type Value interface {
// 	String() string
// 	Set(string) error
// }

type Account struct {
	Username string
	Password string
}

func (a *Account) String() string {
	return fmt.Sprintf("账号：%s， 密码：%s", a.Username, a.Password)
}
func (a *Account) Set(s string) error {
	slice := strings.Split(s, ",")
	if len(slice) != 2 {
		return fmt.Errorf("参数个数不正确")
	}
	a.Username = slice[0]
	a.Password = slice[1]
	return nil
}

// flag 命令行解析参数
func main() {
	name := flag.String("name", "", "姓名")
	age := flag.Int("age", 18, "年龄")

	var gender string
	flag.StringVar(&gender, "gender", "M", "性别")

	// 将参数值绑定到结构体
	var a Account
	flag.Var(&a, "account", "账户")
	flag.Parse() // 必须调用，调用之后才能输出对应的值
	fmt.Println(a)
	fmt.Println(*name, *age, gender)
}
