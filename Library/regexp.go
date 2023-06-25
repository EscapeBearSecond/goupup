package Library

import (
	"fmt"
	"regexp"
)

func RegexpCase() {
	reg := regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	fmt.Println(reg.MatchString("abc[123445]"))

	bytes := reg.FindAll([]byte("eft[1231]1312"), -1)
	fmt.Println(string(bytes[0]))
}
