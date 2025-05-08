package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todos struct {
	// gorm.Model
	ID         int        `json:"id"`
	Title      string     `json:"title"`
	Status     bool       `json:"status"`
	Created_at *time.Time `json:"createAt"`
	Updated_at *time.Time `json:"updatedAt"`
	Deleted_at *time.Time `json:"deletedAt"`
}

var DB *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}

// 配置操作数据库的表名称
func (table *Todos) TableName() string {
	return "todos"
}
