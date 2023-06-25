package generic

// MyStruct 泛型结构体
type MyStruct[T interface{ *int | *string }] struct {
	Name string
	Data T
}
