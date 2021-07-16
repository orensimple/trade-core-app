package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/orensimple/trade-core-app/internal/app/application/usecase"
	"github.com/orensimple/trade-core-app/internal/app/domain"
	"github.com/orensimple/trade-core-app/internal/app/domain/valueobject"
	"golang.org/x/crypto/bcrypt"
)

func (ctrl Controller) health(c *gin.Context) {
	c.JSON(http.StatusOK, domain.SimpleResponse{Status: "OK"})
}

func (ctrl Controller) index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello in Trade-app",
	})
}

func (ctrl Controller) userSearchHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "user_search.tmpl", gin.H{
		"title": "Search other users",
	})
}

func (ctrl Controller) register(c *gin.Context) {
	hash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.MinCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.SimpleResponse{Status: "failed generate hash password"})

		return
	}

	newUser := domain.User{
		ID:        uuid.New(),
		FirstName: "NewCrate",
		Email:     "new@mail.ru",
		Password:  string(hash),
	}

	user, _ := usecase.CreateUser(ctrl.UserRepository, &newUser)
	c.JSON(http.StatusOK, user)
}

func (ctrl Controller) userMock(c *gin.Context) {
	err := usecase.CreateUsersMock(ctrl.UserRepository, 1000000)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.SimpleResponse{Status: "failed generate hash password"})

		return
	}

	c.JSON(http.StatusOK, domain.SimpleResponse{Status: "OK"})
}

func (ctrl Controller) userSearch(c *gin.Context) {
	var req domain.UserSearchRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong request params"})

		return
	}

	res, _ := usecase.SearchUsers(ctrl.UserRepository, &domain.User{LastName: req.LastName, FirstName: req.FirstName})
	c.JSON(http.StatusOK, res)
}

func (ctrl Controller) ticker(c *gin.Context) {
	pair := valueobject.BtcJpy
	ticker := usecase.Ticker(ctrl.BitbankService, pair)
	c.JSON(http.StatusOK, ticker)
}

func (ctrl Controller) candlestick(c *gin.Context) {
	args := usecase.OhlcArgs{
		E: ctrl.BitbankService,
		P: valueobject.BtcJpy,
		T: valueobject.OneMin,
	}
	candlestick := usecase.Ohlc(args)
	c.JSON(http.StatusOK, candlestick)
}

func (ctrl Controller) parameter(c *gin.Context) {
	parameter := usecase.Parameter(ctrl.ParameterRepository)
	c.JSON(http.StatusOK, parameter)
}

func (ctrl Controller) order(c *gin.Context) {
	order := usecase.AddNewCardAndEatCheese(ctrl.OrderRepository)
	c.JSON(http.StatusOK, order)
}