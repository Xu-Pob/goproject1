package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

func main() {
	conn := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", ".", 1433, "GoDb", "sa", "1234.com")
	db, err := sql.Open("mssql", conn)
	if err != nil {
		log.Fatal(err)
	}
	//预编译
	stmt, err := db.Prepare("SELECT id,username from GoDb.dbo.users1 where id =?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(1)
	//在此会检测sql语句问题，
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id       string
			username string
		)
		rows.Scan(&id, &username)
		fmt.Println(id, ":", username)
		break // 如果行没有关闭，则可能会导致内存泄漏
	}
	//(?)  为什么写两个
	//if err = rows.Close(); err != nil {
	//	log.Println(err)
	//}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
