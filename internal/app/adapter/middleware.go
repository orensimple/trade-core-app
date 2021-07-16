package adapter

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/orensimple/trade-core-app/internal/app/application/usecase"
	"github.com/orensimple/trade-core-app/internal/app/domain"
	"github.com/prometheus/common/log"
	"golang.org/x/crypto/bcrypt"
)

func (ctrl Controller) auth() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*domain.User); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}

			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &domain.User{
				Email: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var req domain.LoginRequest
			if err := c.ShouldBind(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			user, err := usecase.GetUser(ctrl.UserRepository, &domain.User{Email: req.Username})
			if err != nil {
				return "", errors.New("user not found")
			}

			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*domain.User); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			// TODO deprecated
			location := url.URL{Path: "/"}
			c.Redirect(http.StatusFound, location.RequestURI())
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		// TODO deprecated
		LoginResponse: func(c *gin.Context, i int, s string, t time.Time) {
			location := url.URL{Path: "/auth/user/search"}
			c.Redirect(http.StatusFound, location.RequestURI())
		},

		// Read cookie
		TokenLookup:   "cookie:token",
		TokenHeadName: "Bearer",

		TimeFunc: time.Now,

		// Send cookie
		SendCookie:     true,
		SecureCookie:   false, // non HTTPS dev environments
		CookieHTTPOnly: true,  // JS can't modify
		CookieDomain:   "localhost",
		CookieName:     "token", // default jwt
		CookieSameSite: http.SameSiteDefaultMode,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware
}
