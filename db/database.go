package db

import (
	"fmt"

	"github.com/funayoseyoshito/yakiniku-image-logo/lib"
	"github.com/jinzhu/gorm"
	//gormで依存がある為
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

//GetConnection データベースコネクションを返却します
func GetConnection() *gorm.DB {
	if db == nil {
		databaseInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
			lib.Config.Database.User,
			lib.Config.Database.Password,
			lib.Config.Database.Host,
			lib.Config.Database.Port,
			lib.Config.Database.Name)

		var err error
		db, err = gorm.Open("mysql", databaseInfo)
		if err != nil {
			fmt.Println(err)
			panic("failed to connect database")
		}
	}
	return db
}
