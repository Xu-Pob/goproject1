package main

// 输出什么
func main() {
	var x *struct {
		s [][32]byte
	}

	println(len(x.s[99]))
}
