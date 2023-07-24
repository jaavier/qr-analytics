package handlers

import (
	"net/http"
	"qranalytics/app"

	"github.com/labstack/echo/v4"
)

func AddLocation(c echo.Context) error {
	qrId := c.Param("qrId")
	allQrCodes := app.RetrieveAllQrCodes()
	if value, ok := allQrCodes[qrId]; !ok {
		return c.String(http.StatusBadRequest, "Not found")
	} else {
		var bodyLocations []app.Location
		c.Bind(&bodyLocations)
		for _, location := range bodyLocations {
			value.AddLocation(location)
		}
		return c.JSON(http.StatusOK, value)
	}
}
