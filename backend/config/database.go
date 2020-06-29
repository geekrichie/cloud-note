package config

import (
	"fmt"

	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DatabaseInfo struct {
	Host     string `ini:"host"`
	Port     int32  `ini:"port"`
	DbName   string `ini:"db"`
	User     string `ini:"user"`
	Password string `ini:"password"`
}

var databaseInfo = DatabaseInfo{}

func readFile() {
	cfg, err := ini.Load("conf/database.ini")
	if err != nil {
		fmt.Println("解析数据库异常:", err)
		return
	}
	err = cfg.Section("database").MapTo(&databaseInfo)
}

func GetDB() *gorm.DB{
	conn := databaseInfo.User + ":" + databaseInfo.Password + "@tcp(" + databaseInfo.Host + ":" + fmt.Sprintf("%d", databaseInfo.Port) + ")/" + databaseInfo.DbName
	DB, err := gorm.Open("mysql", conn)
	//defer DB.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return DB
}
func init() {
	readFile()
}
