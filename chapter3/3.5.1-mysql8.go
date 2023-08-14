package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"time"
)

func main() {
	conn := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", ".", 1433, "GoDb", "sa", "1234.com")
	db, err := sql.Open("mssql", conn)
	if err != nil {
		log.Fatal(err)
	}
	go func(db *sql.DB) {
		mt := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-mt.C:
				stat := db.Stats()
				fmt.Println(stat.MaxOpenConnections)
				er := fmt.Errorf("monitor db conn(%p): maxopen(%d), open(%d), use(%d), idle(%d), "+
					"wait(%d), idleClose(%d), lifeClose(%d), totalWait(%v)",
					db,
					stat.MaxOpenConnections, stat.OpenConnections,
					stat.InUse, stat.Idle,
					stat.WaitCount, stat.MaxIdleClosed,
					stat.MaxLifetimeClosed, stat.WaitDuration)
				fmt.Println(er)
			}
		}
	}(db)
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
