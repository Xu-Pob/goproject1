package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"os"
)

func main() {

	conn := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", ".", 1433, "GoDb", "sa", "1234.com")
	db, err := sql.Open("mssql", conn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	file, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	var (
		id       int
		username string
	)
	//查询数据库
	rows, err := db.Query("SELECT id,username from GoDb.dbo.users where id =?;", 1)

	if err != nil {
		//log.Fatal(err) 默认会在错误消息前添加时间戳和日志级别，以及程序的调用栈信息//相当于print后，os.exit
		//go run XX.go 2> error.log 生成日志
		log.Print(err)
		log.Printf("ERR:%v", err)
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		//通过scan方法赋值
		err := rows.Scan(&id, &username)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, username)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

}
