package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	//db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
	conn := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", ".", 1433, "GoDb", "sa", "1234.com")
	db, err := sql.Open("mssql", conn)
	if err != nil {
		fmt.Println(err)
	}
	//通常用于长时间判断后看连接是否是OK 的（由于超时、断开连接或者其他原因）
	err = db.Ping()
	if err != nil {
		//错误处理逻辑
	}
	defer db.Close()
}
