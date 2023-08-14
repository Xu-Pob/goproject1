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
	stmt, err := db.Prepare("SELECT username from GoDb.dbo.users where id =?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(1)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var s sql.NullString
		err = rows.Scan(&s)
		if s.Valid {
			fmt.Println(s.String)
		} else {
			//NULL值
		}
	}
}
