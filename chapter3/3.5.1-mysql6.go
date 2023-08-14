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
	var name string

	err = db.QueryRow("SELECT id,username from GoDb.dbo.users where id =?", 4).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			//结果没有行
			log.Println("mei有hang")
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(name)
}
