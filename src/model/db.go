package model

import (
	"conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql" //下划线的意义，引用该包，仅仅是为了调用包的init()函数
	"module/log"
)

var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {
		newDb, err := newDB()
		if err != nil {
			log.Panic(err)
		}
		newDb.DB().SetMaxIdleConns(10)
		newDb.DB().SetMaxOpenConns(100)

		//newDb.SetLogger(orm.Logger{})
		newDb.LogMode(false)
		newDb.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

		db = newDb
	}

	return db
}

func newDB() (*gorm.DB, error) {
	sqlConnection := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", conf.Conf.DB.Host, conf.Conf.DB.UserName, conf.Conf.DB.Pwd, conf.Conf.DB.Port, conf.Conf.DB.Name)
	fmt.Println(sqlConnection)
	db, err := gorm.Open("mssql", sqlConnection)
	if err != nil {
		return nil, err
	}

	return db, nil
}
