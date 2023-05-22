package main

// 类库  hmap 原理
// map的初始化
// 1，make     2,字面量
// select,for 阻塞的方法，main方法不会退出，协程也不会退出；类似休眠的效果，不会占用空间  for 会占用空间
func main() {
	//m := make(map[string]int, 10)
	//fmt.Println(m)
	//字面量
	//has := map[string]int{
	//	"a": 1,
	//	"b": 2,
	//}
	//fmt.Println(has["a"])
	xiechengReadWrite()
}
func xiechengReadWrite() {
	m := map[int]int{}
	go func() {
		for {
			_ = m[1]
		}
	}()
	go func() {
		for {
			m[2] = 1
		}
	}()
	//select,for 阻塞的方法，main方法不会退出，协程也不会退出；类似休眠的效果，不会占用空间  for 会占用空间
	select {}
}
