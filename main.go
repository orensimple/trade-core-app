package main

import (
	"github.com/orensimple/trade-core-app/internal/app/adapter"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("BITBANK_HOST", "https://public.bitbank.cc")
	// for local development
	viper.SetDefault("MYSQLHOST", "0.0.0.0:3306")
	viper.SetDefault("MYSQLUSER", "root")
	viper.SetDefault("MYSQLPASSWORD", "my-secret-pw")

	viper.BindEnv("MYSQLHOST", "MYSQLHOST")
	viper.BindEnv("MYSQLUSER", "MYSQLUSER")
	viper.BindEnv("MYSQLPASSWORD", "MYSQLPASSWORD")

	r := adapter.Router()
	err := r.Run(":8082")
	if err != nil {
		panic(err)
	}
}
