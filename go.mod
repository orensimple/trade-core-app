module github.com/orensimple/trade-core-app

// +heroku goVersion go1.15
go 1.15

require (
	github.com/appleboy/gin-jwt/v2 v2.6.4
	github.com/brianvoe/gofakeit v3.18.0+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.1.2
	github.com/penglongli/gin-metrics v0.1.4 // indirect
	github.com/prometheus/common v0.10.0
	github.com/spf13/viper v1.7.1
	golang.org/x/crypto v0.0.0-20190605123033-f99c8df09eb5
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.9
)
