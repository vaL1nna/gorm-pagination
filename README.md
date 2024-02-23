# Gorm Pagination

## Installation

```bash
go get -u github.com/val1nna/gorm-pagination
```

## Usage

```go
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	pagination "github.com/val1nna/gorm-pagination"
)

type User struct {
	ID       int    `gorm:"primary_key"`
	Username string `gorm:"type:varchar(100);not null"`
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	var users []User

	paginator := pagination.Paginate(&pagination.Option{
		DB:      db,
		Page:    1,
		Limit:   10,
		ShowSQL: true,
	}, &users)

	fmt.Println(paginator)
}
```