package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"tomatoPaper/api/entity"
	"tomatoPaper/common/config"
)

var Db *gorm.DB

// SetupDBLink 处理数据库连接
func SetupDBLink() error {
	var err error
	var dbConfig = config.Config.Db
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset,
		dbConfig.Loc)
	//url = fmt.Sprintf("root:010729@tcp(127.0.0.1:3306)/tomato_paper?charset=utf8mb4&parseTime=True&loc=Local")
	Db, err = gorm.Open(mysql.Open(url), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		//panic(err)
		fmt.Println("数据库连接失败:", err)
	}
	err = Db.AutoMigrate(entity.Users{})
	if err != nil {
		fmt.Println("创建数据库表格失败:", err)
	}

	if Db.Error != nil {
		panic(Db.Error)
	}
	sqlDB, err := Db.DB()
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpen)
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdle)
	return nil
}
