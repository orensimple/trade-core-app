package adapter

import (
	"net/http"
	"strconv"

	"github.com/prometheus/common/log"

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
	var req domain.RegisterRequest
	err := c.ShouldBind(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong request params"})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, domain.SimpleResponse{Status: "failed generate hash password"})

		return
	}

	male, err := strconv.ParseBool(req.Male)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong request params, gender"})

		return
	}

	user, err := usecase.GetUser(ctrl.UserRepository, &domain.User{Email: req.Email})
	if err != nil && err.Error() != "user not found" {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "failed get user"})

		return
	}
	if user != nil {
		c.JSON(http.StatusConflict, domain.SimpleResponse{Status: "user with such mail exists"})

		return
	}

	newUser := domain.User{
		ID: uuid.New(),
		// TODO key by email
		Email:     req.Email,
		Password:  string(hash),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Passport:  req.Passport,
		Address:   req.Address,
		About:     req.About,
		Male:      male,
	}

	usecase.CreateUser(ctrl.UserRepository, &newUser)

	// location := url.URL{Path: "/"}
	// c.Redirect(http.StatusFound, location.RequestURI())
	c.JSON(http.StatusOK, newUser)
}

func (ctrl Controller) userMock(c *gin.Context) {
	err := usecase.CreateUsersMock(ctrl.UserRepository, 1000000)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, domain.SimpleResponse{Status: "failed generate hash password"})

		return
	}

	c.JSON(http.StatusOK, domain.SimpleResponse{Status: "OK"})
}

func (ctrl Controller) userSearch(c *gin.Context) {
	var req domain.UserSearchRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong request params"})

		return
	}

	res, _ := usecase.SearchUsers(ctrl.UserRepository, &domain.User{LastName: req.LastName, FirstName: req.FirstName})
	c.JSON(http.StatusOK, res)
}

func (ctrl Controller) userGet(c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong id"})

		return
	}

	res, err := usecase.GetUser(ctrl.UserRepository, &domain.User{ID: uuid})
	if err != nil && err.Error() != "user not found" {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "failed get user"})

		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, domain.SimpleResponse{Status: "user not found"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (ctrl Controller) userUpdate(c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong id"})

		return
	}

	user, err := usecase.GetUser(ctrl.UserRepository, &domain.User{ID: uuid})
	if err != nil && err.Error() != "user not found" {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "failed get user"})

		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, domain.SimpleResponse{Status: "user not found"})

		return
	}

	var req domain.RegisterRequest
	err = c.ShouldBind(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong request params"})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, domain.SimpleResponse{Status: "failed generate hash password"})

		return
	}

	male, err := strconv.ParseBool(req.Male)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong request params, gender"})

		return
	}

	user.Email = req.Email
	user.Password = string(hash)
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Passport = req.Passport
	user.Address = req.Address
	user.About = req.About
	user.Male = male

	err = usecase.UpdateUser(ctrl.UserRepository, user)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, domain.SimpleResponse{Status: "failed delete user"})

		return
	}

	c.JSON(http.StatusOK, user)
}

func (ctrl Controller) userDelete(c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong id"})

		return
	}

	err = usecase.DeleteUser(ctrl.UserRepository, &domain.User{ID: uuid})
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, domain.SimpleResponse{Status: "failed delete user"})

		return
	}

	c.JSON(http.StatusOK, domain.SimpleResponse{Status: "OK"})
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
