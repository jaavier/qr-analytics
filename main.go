package main

import (
	"fmt"
	"qranalytics/app/handlers"

	"github.com/labstack/echo/v4"
)

const (
	LOW      = 0
	MEDIUM   = 1
	HIGH     = 2
	CRITICAL = 3
)

func main() {
	var err error
	var e *echo.Echo
	var handleError = func() {
		fmt.Println("CUSTOM CALL")
	}

	err = LoadEnv()
	logEvent(LogEvent{level: LOW, err: err, failed: func() {
		err = LoadEnv("./.envs")
		logEvent(LogEvent{level: LOW, err: err, success: func() { fmt.Println("LOADED ./.envs") }})
	}})

	err = Connect()
	logEvent(LogEvent{level: LOW, err: err, success: func() {
		fmt.Println("CONNECTED TO MONGODB")
	}})

	e = echo.New()

	e.GET("/api/qr", handlers.RetrieveAllQrCodes)
	e.POST("/api/qr", handlers.AddQrCode)
	e.POST("/api/qr/:qrId/locations", handlers.AddLocation)

	err = e.Start(":3000")

	logEvent(LogEvent{level: CRITICAL, err: err, failed: handleError})
}
