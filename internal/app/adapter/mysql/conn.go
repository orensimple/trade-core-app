package mysql

import (
	"fmt"
	"time"

	"github.com/prometheus/common/log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connection gets connection of mysql database
func Connection() (db *gorm.DB) {
	host := viper.Get("mysql_host")
	port := viper.Get("mysql_port")
	user := viper.Get("mysql_user")
	pass := viper.Get("mysql_password")
	dataBase := viper.Get("mysql_database")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, dataBase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
