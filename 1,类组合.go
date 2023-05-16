package main

func main() {

}

type People struct {
	name string
	tel  string
}

func (p People) walk() {

}

type Man struct {
	People
}

type alive interface {
	walk()
	jamp()
}
