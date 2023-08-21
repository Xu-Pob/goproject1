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

	ret, err := db.Exec(`insert into GoDb.dbo.users(username)values(?)`, "王五")
	if err != nil {
		log.Fatal(err)
	}
	//应该是sqlserver不能用，mysql可能可以
	addID, err := ret.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("insert success, the id is %d.\n", addID)
}
