package mysql

import (
	"fmt"
	"time"

	"github.com/prometheus/common/log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// Connection gets connection of mysql database
func Connection() (db *gorm.DB) {
	var hostSuffix string
	replicaEnabled := viper.GetBool("mysql_replica_enabled")
	host := viper.Get("mysql_host")
	port := viper.Get("mysql_port")
	user := viper.Get("mysql_user")
	pass := viper.Get("mysql_password")
	dataBase := viper.Get("mysql_database")

	if replicaEnabled {
		hostSuffix = "-primary"
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v%v:%v)/%v", user, pass, host, hostSuffix, port, dataBase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}

	if replicaEnabled {
		replicaCount := viper.GetInt("mysql_replica_count")
		hostSuffix = "-secondary"
		for i := 1; i <= replicaCount; i++ {
			dsn := fmt.Sprintf("%v:%v@tcp(%v%v:%v)/%v", user, pass, host, hostSuffix, port, dataBase)

			db.Use(dbresolver.Register(dbresolver.Config{
				Replicas: []gorm.Dialector{mysql.Open(dsn)},
			}))
		}
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
