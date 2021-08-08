package main

import (
	"fmt"

	"github.com/prometheus/common/log"

	"github.com/orensimple/trade-core-app/internal/app/adapter"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("bitbank_host", "https://public.bitbank.cc")
	// for local development
	viper.SetDefault("mysql_host", "0.0.0.0")
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

	viper.BindEnv("mysql_user", "MYSQL_USER")
	viper.BindEnv("mysql_password", "MYSQL_PASSWORD")

	r := adapter.Router()
	port := viper.Get("app_port")
	err = r.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Error(err)
	}
}
