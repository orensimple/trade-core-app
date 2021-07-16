package mysql

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connection gets connection of mysql database
func Connection() (db *gorm.DB) {
	host := viper.Get("MYSQLHOST")
	user := viper.Get("MYSQLUSER")
	pass := viper.Get("MYSQLPASSWORD")
	dataBase := viper.Get("MYSQLDATABASE")
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v", user, pass, host, dataBase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
