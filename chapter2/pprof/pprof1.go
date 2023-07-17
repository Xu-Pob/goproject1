package main

import (
	"log"
	"net/http"
	"os"
)

// 性能分析
func main() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

	go func() {
		log.Println(http.ListenAndServe("8802", nil))
	}()

	for {
		Sum("This is Test")
	}
}
func Sum(str string) string {
	data := []byte(str)
	sData := string(data)
	var sum = 0
	for i := 0; i < 10000; i++ {
		sum += 1
	}
	return sData
}
