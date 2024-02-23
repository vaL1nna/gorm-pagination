package pagination

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

var Db *gorm.DB

type User struct {
	ID       int    `gorm:"primary_key"`
	Username string `gorm:"type:varchar(100);not null"`
}

func connectDb() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
}

func TestPaginate(t *testing.T) {
	connectDb()

	var users []User

	paginator := Paginate(&Option{
		DB:      Db,
		Page:    1,
		Limit:   10,
		ShowSQL: true,
	}, &users)

	t.Log(paginator)
}
