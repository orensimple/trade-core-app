package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/orensimple/trade-core-app/internal/app/adapter/mysql"
	"github.com/orensimple/trade-core-app/internal/app/adapter/repository"
	"github.com/orensimple/trade-core-app/internal/app/adapter/service"
	"github.com/penglongli/gin-metrics/ginmetrics"
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

	// init prometheus metrics
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(r)

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
	r.POST("/register", ctrl.register)
	r.GET("/logout", authMiddleware.LogoutHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.GET("/users/search", ctrl.userSearchHTML)
	}

	api := r.Group("/api")
	// api.Use(authMiddleware.MiddlewareFunc())
	{
		api.POST("/user", ctrl.register)
		api.GET("/user/:id", ctrl.userGet)
		api.PUT("/user/:id", ctrl.userUpdate)
		api.DELETE("/user/:id", ctrl.userDelete)
		api.GET("/users/search/", ctrl.userSearch)
		api.GET("/users/mock/", ctrl.userMock)

		api.GET("/ticker", ctrl.ticker)
		api.GET("/candlestick", ctrl.candlestick)

		api.GET("/parameter", ctrl.parameter)
		api.GET("/order", ctrl.order)
	}

	return r
}
