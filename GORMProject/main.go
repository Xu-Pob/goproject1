package main

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)
import "gorm.io/driver/sqlserver"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type DeTel struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString `gorm:"column:MemberNumber"` //自定义列名
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type GORMUser struct {
	ID        uint   // 列名是 `id`
	Name      string // 列名是 `name`
	Age       int
	Birthday  time.Time // 列名是 `birthday`
	CreatedAt time.Time // 列名是 `created_at`
}

// TableName 会将 DeTel 的表名重写为 `DdTel`   //动态表名官方文档
func (d DeTel) TableName() string {
	return "DdTel"
}

func main() {
	conn := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", ".", 1433, "GoDb", "sa", "1234.com")

	db, err := gorm.Open(sqlserver.Open(conn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema(数据类型)
	//db.AutoMigrate(&GORMUser{})

	//parse, err := time.Parse("2006-01-02", "2006-01-02")
	//if err != nil {
	//	return
	//}

	//user2 := GORMUser{Name: "jinzhu", Age: 18, Birthday: time.Now()}

	//result := db.Create(&user2)
	//fmt.Println(result.Error)
	//fmt.Println(result.RowsAffected)

	//批量新增
	//users := []GORMUser{
	//	{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
	//	{Name: "Jackson", Age: 19, Birthday: time.Now()},
	//}
	////db.Create(&users)
	////创建指定字段
	//db.Select("Name", "Age", "CreatedAt").Create(&users)

	var users = []GORMUser{}
	for i := 0; i <= 1000; i++ {
		users = append(users, GORMUser{Name: fmt.Sprintf("pob_%d", i)})
	}
	// batch size 100
	//db.CreateInBatches(users, 100)

	//db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
	//	CreateBatchSize: 1000,
	//})

	Query(db)

	// 想要修改该值，您可以使用 `Update`
	//db.First(&user2, "Name =?", "jinzhu1")
	//db.Model(&user2).Update("Name", "jinzhu2")

}

// Query First() 主键升序第一个,  Last()主键升序最后一个  Take()默认查询第一个
func Query(db *gorm.DB) {
	var user map[string]interface{}
	//result := db.Model(&GORMUser{}).Where(map[string]string{
	//	"name": "Jinzhu",
	//}).First(&user)

	result := db.Table("gorm_users").First(&user)
	//db.Take(&user)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	fmt.Println(user)
	fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))

}
