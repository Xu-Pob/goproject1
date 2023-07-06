package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"os"
	"sync"
)

var (
	tables = []string{"users", "orders"}
	count  = len(tables)
	//ch     = make(chan bool, count)
)

func main() {
	//db, err := sql.Open("mysql", "root:a123456@tcp(127.0.0.1:3306)/chapter1?charset=utf8")
	conn := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", ".", 1433, "GoDb", "sa", "1234.com")
	db, err := sql.Open("mssql", conn)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(2, err)
		}
		db.Close()
	}()
	if err != nil {
		//fmt.Println(1, err)
		panic(err.Error())
	}

	var wg sync.WaitGroup

	for _, table := range tables {
		wg.Add(1)
		//开启查询表的协程
		go SqlQuery(db, table, &wg)

	}
	//for i := 0; i < count; i++ {
	//<-ch
	//}
	wg.Wait()
	fmt.Println("连接成功，并完成")
}

// SqlQuery 运行sql
func SqlQuery(db *sql.DB, table string, group *sync.WaitGroup) {
	fmt.Println("开始处理：", table)
	rows, _ := db.Query(fmt.Sprintf("select *from %s", table))

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range scanArgs {
		scanArgs[i] = &values[i]
	}
	var totalValues [][]string
	for rows.Next() {
		var s []string
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		for _, v := range values {
			s = append(s, string(v))
		}
		totalValues = append(totalValues, s)
	}
	//fmt.Println(totalValues)
	//ch <- true
	exportToCSV(table+".csv", columns, totalValues)
	group.Done()
}
func exportToCSV(file string, column []string, totalValues [][]string) {
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	f.WriteString("\xEF\xBB\xBF")
	defer f.Close()
	write := csv.NewWriter(f)
	for i, v := range totalValues {
		if i == 0 {
			write.Write(column)
		}
		write.Write(v)
	}
	write.Flush()
	fmt.Println("处理完毕", file)
}
