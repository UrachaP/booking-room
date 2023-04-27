package router

import (
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// JWTMiddleware Middleware
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := c.Request().Header.Get("Authorization")
		token := strings.Replace(accessToken, "Bearer ", "", -1)
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}
		//log.Println(claims)
		log.Println(claims["sub"])

		return next(c)
	}
}

func JwtBasicMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwt.StandardClaims{},
		SigningKey: []byte("secret"),
		SuccessHandler: func(c echo.Context) {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*jwt.StandardClaims)
			log.Println(claims.Subject)
		},
	})
}

func (r *router) revokedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwt.StandardClaims)

		revoked, err := r.handlers.GetRevokedTokenFromRedis(claims.Id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if revoked != "0" {
			return c.JSON(http.StatusUnauthorized, "revoked token")
		}
		return next(c)
	}
}
