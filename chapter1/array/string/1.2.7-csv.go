package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	CsvRead()
}
func CsvWrite() {
	//创建一个csv文件
	f, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	//
	w := csv.NewWriter(f)
	w.Write([]string{"学号", "姓名", "分数"})
	w.Write([]string{"1", "Barry", "99.5"})
	w.Write([]string{"2", "Shirdon", "100"})
	w.Flush()
}

func CsvRead() {
	//创建一个csv文件
	f, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.Comma = ','
	reader.FieldsPerRecord = -1
	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}
}
