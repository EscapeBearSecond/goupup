package Library

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func EncodingCase() {
	type User struct {
		ID   int64
		Name string
		Age  int
	}
	u := User{
		ID:   1,
		Name: "Jack",
		Age:  19,
	}
	// json序列化
	bytes, err := json.Marshal(u)
	fmt.Println(bytes, err)
	u1 := User{}
	// 反序列化
	err = json.Unmarshal(bytes, &u1)
	fmt.Println(u1, err)
	// base64编码解码
	str := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(str)
	bytes1, err := base64.StdEncoding.DecodeString(str)
	fmt.Println(bytes1, err)
}
