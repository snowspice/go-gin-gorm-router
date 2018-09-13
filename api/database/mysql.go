package database
import (
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"fmt"
	tool "github.com/snowspice/go-gin-gorm-router/api/config"

)

var Eloquent *gorm.DB
//var conf *tool.Config

func init() {
	var err error
	var  conf = tool.GetAppConf()
	fmt.Println("my sql value is ---> ",conf.Mysql)

	Eloquent, err = gorm.Open("mysql", conf.Mysql)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	Eloquent.DB().SetMaxIdleConns(10)
	Eloquent.DB().SetMaxOpenConns(100)
	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
}