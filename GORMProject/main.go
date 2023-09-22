package main

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

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
	//conn := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", ".", 1433, "GoDb", "sa", "1234.com")

	db, err := gorm.Open(mysql.Open("root:1234.com@tcp(127.0.0.1:3306)/mysqlgodb"), &gorm.Config{})

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
		users = append(users, GORMUser{Name: fmt.Sprintf("pob_%d", i),
			Birthday: time.Now(),
		})
	}
	// batch size 100
	//db.CreateInBatches(users, 100)

	//db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
	//	CreateBatchSize: 1000,
	//})

	UpdateBatchExample(db)

	// 想要修改该值，您可以使用 `Update`
	//db.First(&user2, "Name =?", "jinzhu1")
	//db.Model(&user2).Update("Name", "jinzhu2")

}

// Query First() 主键升序第一个,  Last()主键升序最后一个  Take()默认查询第一个
func Query(db *gorm.DB) {
	//var user map[string]interface{}
	//result := db.Model(&GORMUser{}).Where(map[string]string{
	//	"name": "Jinzhu",
	//}).First(&user)

	//db.Model(&GORMUser{}).First(&user)

	//user为 map 类型时 ， First和last不工作，  result.Error  -->  model value required
	//result := db.Table("gorm_users").First(&user)

	var user []GORMUser
	result := db.Table("gorm_users").Last(&user)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	fmt.Println(user)
	fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
}

func QueryById(db *gorm.DB) {
	//var user GORMUser
	//result := db.First(&user, 10)
	//SELECT * FROM gorm_users WHERE id = 10;

	//result := db.First(&user, "name = ?", "pob_4")
	//SELECT * FROM users WHERE name = “pob_4”;

	//var user []map[string]interface{}
	//result := db.Model(&GORMUser{}).Find(&user, []int{1, 2, 3})

	var user []GORMUser
	result := db.Find(&user, []int{1, 2, 3})
	// SELECT * FROM gorm_users WHERE id IN (1,2,3);   user用结构体切片或数组接收

	//var user = GORMUser{ID: 10}
	//result := db.First(&user)

	//var user GORMUser
	//result := db.Model(&GORMUser{ID: 10}).First(&user)
	////SELECT * FROM gorm_users WHERE id = 10;

	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	fmt.Println(user)
	fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
}
func QueryWhere(db *gorm.DB) {
	var user GORMUser
	//select -》选择特定值
	result := db.Select("name").Not("name = ?", "jinzhu").First(&user)

	rows, _ := db.Table("gorm_users").Select("COALESCE(age,?)", 42).Rows()

	for rows.Next() {
		var ctr sql.RawBytes
		err := rows.Scan(&ctr)
		if err != nil {
			return
		}
		fmt.Println(string(ctr))

	}

	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	fmt.Println(user)
	fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
}
func QueryOrderBy(db *gorm.DB) {
	var user []GORMUser

	result := db.Order("age desc, name").Find(&user) //db.Order("age desc").Order("name").Find(&users)
	//result := db.Clauses(clause.OrderBy{
	//	Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	//}).Find(&user)
	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	fmt.Println(user)
	fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
}

func QueryLimitOffset(db *gorm.DB) {
	var user []GORMUser

	result := db.Limit(3).Offset(3).Find(&user) //result := db.Limit(3).Find(&user)

	//result := db.Offset(10).Find(&user).Offset(-1).Find(&user)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	fmt.Println(user)
	fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
}

func QueryGroupBy(db *gorm.DB) {
	type result struct {
		Name  string
		Date  time.Time
		Total int
	}
	var ret result
	reaa := db.Model(&GORMUser{}).Select("name, sum(age) as total").Where("name LIKE ?", "pob%").Group("name").Take(&ret)

	//var ret []result
	//reaa := db.Model(&GORMUser{}).Select("name, sum(age) as total").Group("name").Having("name like ?", "pob%").Find(&ret)

	fmt.Println(reaa.RowsAffected)
	fmt.Println(reaa.Error)
	fmt.Println(ret)
	fmt.Println(errors.Is(reaa.Error, gorm.ErrRecordNotFound))
}
func QueryDistinct(db *gorm.DB) {
	var users []GORMUser
	result := db.Debug().Distinct("name", "age").Order("name, age desc").Find(&users)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	fmt.Println(users)
	fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
}
func UpdateBatchExample(db *gorm.DB) {
	//result := db.Model(&GORMUser{}).Update("name", "jinzhu")
	// WHERE conditions required

	//result := db.Exec("UPDATE gorm_users SET name = ?", "jinzhu")
	result := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&GORMUser{}).Update("name", "jinzhu1")
	// UPDATE users SET `name` = "jinzhu"

	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)

}
