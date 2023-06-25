package generic

import "fmt"

type ToString interface {
	String() string
}
type keyAble[T comparable] interface {
	any
	GetKey() T
}

func (u user) String() string {
	return fmt.Sprintf("ID: %d, name: %s, age: %d", u.ID, u.Name, u.Age)
}

func (addr address) String() string {
	return fmt.Sprintf("ID: %d, province: %s, city: %s", addr.ID, addr.Province, addr.City)
}

func (addr address) GetKey() int {
	return addr.ID
}

func (u user) GetKey() int64 {
	return u.ID
}
func listToMap[k comparable, T keyAble[k]](list []T) map[k]T {
	mp := make(GenericMap[k, T], len(list))
	for _, data := range list {
		mp[data.GetKey()] = data
	}
	return mp
}
func InterfaceCase() {
	userList := []keyAble[int64]{
		user{
			Name: "nick",
			Age:  19,
			ID:   1,
		},
		user{
			Name: "lucy",
			Age:  20,
			ID:   2,
		},
	}
	addrList := []keyAble[int]{
		address{
			ID:       1,
			Province: "广西",
			City:     "桂林",
		},
		address{
			ID:       2,
			Province: "江苏",
			City:     "连云港",
		},
	}
	userMp := listToMap[int64, keyAble[int64]](userList)
	addrMp := listToMap[int](addrList)
	fmt.Println(userMp)
	fmt.Println(addrMp)
}
