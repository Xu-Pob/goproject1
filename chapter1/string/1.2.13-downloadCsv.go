package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/down", Welcome)
	err := http.ListenAndServe(":8070", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	//定义导出的文件名
	filename := "exportUsers.csv"

	//定义一个二维数组
	column := [][]string{
		{"手机号", "用户UID", "Email", "用户名"},
		{"18888888888", "2", "barry@163.com", "barry"},
		{"18888888889", "3", "wangwu@163.com", "wangwu"},
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("content-disPosition", "attachment; filename=\""+filename+"\"")
	GenerateCsv(filename, column)

	//读取文件
	file, err := os.Open(filename)

	content, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("Read File Err:", err.Error())
	} else {
		w.Write(content)
	}
	//w.Write(b)
}

// s生成文件
func GenerateCsv(filePath string, data [][]string) {
	fp, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("创建文件[\"+filePath+\"]句柄失败,%v", err)
	}
	defer fp.Close()
	fp.WriteString("\xEF\xBB\xBF") //写入UTF-8 BOM
	w := csv.NewWriter(fp)
	w.WriteAll(data)
	w.Flush()
}
