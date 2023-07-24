package main

import (
	"qranalytics/app/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	LoadEnv()
	Connect()
	e := echo.New()

	e.GET("/api/qr", handlers.RetrieveAllQrCodes)
	e.POST("/api/qr", handlers.AddQrCode)
	e.POST("/api/qr/:qrId/locations", handlers.AddLocation)

	e.Start(":3000")
}
