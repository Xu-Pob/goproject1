package main

import "fmt"

//defer :延迟函数，可以定义多个延时函数，这些函数会放入到一个栈中，当函数执行 到最后时，这些defer语句会按照逆序执行，最后该函数返回

func main() {
	//defer func1()
	//defer func2()
	//defer func3()

	println(DeferFunc1(1))

	//deferPanic3()

	//defer function(1, function(3, 0))
	//defer function(2, function(4, 0))
}
func func1() {
	fmt.Println("A")
}
func func2() {
	fmt.Println("B")
	a()
}
func func3() {
	fmt.Println("C")
}
func a() {
	i := 0
	defer func() {
		i++
		println("a defer1", i)
	}()
	defer func() {
		i++
		println("a defer2", i)
	}()

	i++
}

// 有名函数返回值遇见 defer 情况  1
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

// 未定义返回值
func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

// defer 遇见 panic
// 第一种情况：遇到panic不捕获
func deferPanic1() {
	defer fmt.Println("defer1`")
	defer fmt.Println("defer2`")
	panic("发生异常")
	defer fmt.Println("defer3`")
}

// 第一种情况：遇到panic   捕获
func deferPanic2() {
	defer func() {
		fmt.Println("defer: panic 之前1, 捕获异常")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() { fmt.Println("defer: panic 之前2, 不捕获") }()

	panic("异常内容") //触发defer出栈

	defer func() { fmt.Println("defer: panic 之后, 永远执行不到") }()

}

// defer中含panic
func deferPanic3() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()
	defer func() {
		panic("defer panic1")
	}()
	defer func() {
		panic("defer panic2")
	}()
	panic("panic")
}

func function(index int, value int) int {
	fmt.Println(index)
	return index
}
