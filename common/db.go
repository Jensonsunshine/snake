package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
)

const MYSQL = "mysql"
const HOST = "localhost"
const PROT = "3306"
const DATABASE = "snake"
const USERNAME = "root"
const PASSWORD = ""
const CHARSET = "utf8"

func InitDB() *gorm.DB {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		USERNAME,
		PASSWORD,
		HOST, /**/
		PROT,
		DATABASE,
		CHARSET,
	)

	db, err := gorm.Open(MYSQL, args)
	//db, err := gorm.Open("mysql", "user:islot@/blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}

	//自动创建数据表
	db.AutoMigrate()
	return db
}

func GetJson(content *gin.Context) {
	data := content.Query("data")
	message := content.Query("message")
	code := content.Query("code")

	content.JSON(http.StatusOK, gin.H{
		data:    data,
		message: message,
		code:    code,
	})
}
