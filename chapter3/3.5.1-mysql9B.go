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

	stat, err := db.Prepare(`UPDATE users set UserName = ? where id=?`)
	if err != nil {
		log.Fatal(err)
	}
	ret, err := stat.Exec("赵六", 3)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := ret.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("insert success, the id is %d.\n", rowsAffected)
}
