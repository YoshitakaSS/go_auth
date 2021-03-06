package web

import (
	"net/http"

	"github.com/YoshitakaSS/go_auth/presenter/controller"
	"github.com/YoshitakaSS/go_auth/presenter/request"
	"github.com/labstack/echo"
)

// NewServer is return echo mapping
func NewServer() *echo.Echo {
	e := echo.New()
	e.Validator = request.NewValidator()

	// TOP
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "TOP")
	})

	// API Group
	api := e.Group("/api")

	// api.Use(middleware.BodyDump(bodyDumpHandler))

	api.GET("/users/:", func(c echo.Context) (err error) {
		return controller.FindUserByName(c)
	})

	// ユーザー登録
	api.POST("/users/regsiter", func(c echo.Context) (err error) {
		return controller.RegisterUser(c)
	})

	// ユーザーログイン
	api.POST("/users/login", func(c echo.Context) error {
		return c.String(http.StatusOK, "ユーザーログイン用Route")
	})

	return e
}

// func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
// 	fmt.Printf("Request Body: %v\n", string(reqBody))
// 	fmt.Printf("Response Body: %v\n", string(resBody))
// }
