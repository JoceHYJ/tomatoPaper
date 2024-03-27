package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"tomatoPaper/common/config"
)

var (
	GormDB *gorm.DB
)

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
	GormDB, err = gorm.Open(mysql.Open(url), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		//panic(err)
		fmt.Println("数据库连接失败:", err)
	}
	if GormDB.Error != nil {
		panic(GormDB.Error)
	}
	sqlDB, err := GormDB.DB()
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpen)
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdle)
	return nil
}
