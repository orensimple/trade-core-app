package main

import (
	"fmt"

	"github.com/prometheus/common/log"

	"github.com/orensimple/trade-core-app/internal/app/adapter"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("BITBANK_HOST", "https://public.bitbank.cc")
	// for local development
	viper.SetDefault("MYSQLHOST", "0.0.0.0:3306")
	viper.SetDefault("MYSQLUSER", "root")
	viper.SetDefault("MYSQLPASSWORD", "my-secret-pw")
	viper.SetDefault("MYSQLDATABASE", "trade")
	viper.SetDefault("PORT", "8083")
	viper.SetDefault("DOMAIN", "localhost")

	viper.BindEnv("MYSQLHOST", "MYSQLHOST")
	viper.BindEnv("MYSQLUSER", "MYSQLUSER")
	viper.BindEnv("MYSQLPASSWORD", "MYSQLPASSWORD")
	viper.BindEnv("MYSQLDATABASE", "MYSQLDATABASE")
	viper.BindEnv("PORT", "PORT")
	viper.BindEnv("DOMAIN", "DOMAIN")

	r := adapter.Router()
	port := viper.Get("PORT")
	err := r.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Error(err)
	}
}
