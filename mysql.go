package codegame

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type db struct {
	Conn *gorm.DB
}

var (
	dataSource string
	DB         db
	lock       = new(sync.RWMutex)
)

func ConnMySQL() {
	// 连接MySQL
	var err error

	dataSource = "root" + ":" + "123456" + "@tcp(" + "172.16.83.86" + ":" + "3306" + ")/" + "codegame" +
		"?charset=" + "utf-8" + "&parseTime=true&loc=Local"

	lock.Lock()
	defer lock.Unlock()
	DB.Conn, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err.Error())
	}

	sqlDB, err := DB.Conn.DB()
	if err != nil {
		panic(err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}
