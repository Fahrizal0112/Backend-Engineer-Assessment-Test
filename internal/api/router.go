package api

import (
	"banking-service/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter(accountService *service.AccountService) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	handler := NewHandler(accountService)

	e.POST("/daftar", handler.HandleDaftar)
	e.POST("/tabung", handler.HandleTabung)
	e.POST("/tarik", handler.HandleTarik)
	e.GET("/saldo/:no_rekening", handler.HandleSaldo)

	return e
}
