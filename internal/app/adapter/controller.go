package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/orensimple/trade-core-app/internal/app/adapter/mysql"
	"github.com/orensimple/trade-core-app/internal/app/adapter/repository"
	"github.com/orensimple/trade-core-app/internal/app/adapter/service"
)

// Controller is a controller
type Controller struct {
	BitbankService      service.Bitbank
	ParameterRepository repository.Parameter
	OrderRepository     repository.Order
	UserRepository      repository.User
}

const identityKey = "id"

// Router is routing settings
func Router() *gin.Engine {
	r := gin.Default()
	db := mysql.Connection()

	userRepository := repository.NewUserRepo(db)
	orderRepository := repository.NewOrderRepo(db)
	parameterRepository := repository.NewParameterRepo(db)

	ctrl := Controller{
		UserRepository:      userRepository,
		OrderRepository:     orderRepository,
		ParameterRepository: parameterRepository,
	}

	// init the jwt middleware
	authMiddleware := ctrl.auth()

	r.LoadHTMLGlob("internal/app/adapter/view/*")
	r.GET("/health", ctrl.health)
	r.GET("/", ctrl.index)
	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/register", ctrl.register)
	r.GET("/logout", authMiddleware.LogoutHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.GET("/user/search", ctrl.userSearchHTML)
	}

	api := r.Group("/api")
	api.Use(authMiddleware.MiddlewareFunc())
	{
		api.GET("/user/mock", ctrl.userMock)
		api.GET("/user/search", ctrl.userSearch)

		api.GET("/ticker", ctrl.ticker)
		api.GET("/candlestick", ctrl.candlestick)

		api.GET("/parameter", ctrl.parameter)
		api.GET("/order", ctrl.order)
	}

	return r
}
