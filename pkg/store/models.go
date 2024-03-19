package store

import (
	"fmt"
	"github.com/youth/product/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func Setup() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)

	//conf := mysql.Config{
	//	DSN: dsn,
	//	DefaultStringSize:         256,
	//	DisableDatetimePrecision:  true,
	//	DontSupportRenameIndex:    true,
	//	DontSupportRenameColumn:   true,
	//	SkipInitializeWithVersion: false,
	//}
	//db, err := gorm.Open(mysql.New(conf) , &gorm.Config{})
	//if err != nil {
	//	log.Fatalf("database err : %v", err)
	//}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database err : %v", err)
	}

	//connection Pool setting
	sqlDB, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
}
