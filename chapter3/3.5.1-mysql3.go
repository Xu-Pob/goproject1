package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

//事务的使用

func main() {

	conn := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", ".", 1433, "GoDb", "sa", "1234.com")
	db, err := sql.Open("mssql", conn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//开启事务
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	//回滚事务(?)
	defer tx.Rollback()
	var (
		id       int
		username string
	)
	//查询数据库
	rows, err := db.Query("SELECT id,username from GoDb.dbo.users where id =?;", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &username)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, username)
	}

	//提交事务 (?)
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}
