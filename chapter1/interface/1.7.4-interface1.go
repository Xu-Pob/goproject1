package main

import (
	"fmt"
	"reflect"
)

// 接口的类型断言
func main() {
	//var a interface{} = func(a int) string {
	//	return fmt.Sprintf("d:%d", a)
	//}
	var a1 interface{} = new(int)
	a2 := a1.(*int)
	*a2 = 1

	fmt.Println(reflect.TypeOf(a1))
	switch b := a1.(type) {
	case nil:
		println("nil")
	case *int:
		println(*b)
	case func(int2 int) string:
		println(b(3))
	case fmt.Stringer:
		fmt.Println(b)
	default:
		println("unknown")
	}
}
