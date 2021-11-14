package main

import (
	"fmt"

	"github.com/prometheus/common/log"

	"github.com/orensimple/trade-core-app/internal/app/adapter"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("bitbank_host", "https://public.bitbank.cc")
	viper.SetDefault("billing_host", "http://localhost:8081")
	viper.SetDefault("order_host", "http://localhost:8082")
	// for local development
	viper.SetDefault("mysql_host", "0.0.0.0")
	viper.SetDefault("mysql_replica_enabled", "false")
	viper.SetDefault("mysql_replica_count", "0")
	viper.SetDefault("mysql_port", "3306")
	viper.SetDefault("mysql_user", "root")
	viper.SetDefault("mysql_password", "my-secret-pw")
	viper.SetDefault("mysql_database", "trade")
	viper.SetDefault("app_port", "80")
	viper.SetDefault("app_domain", "localhost")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error(err)
	}

	err = viper.BindEnv("mysql_user", "MYSQL_USER")
	if err != nil {
		log.Error(err)
	}
	err = viper.BindEnv("mysql_password", "MYSQL_PASSWORD")
	if err != nil {
		log.Error(err)
	}

	r := adapter.Router()
	port := viper.Get("app_port")
	err = r.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Error(err)
	}
}
