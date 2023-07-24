package handlers

import (
	"fmt"
	"net/http"
	"qranalytics/app"
	"reflect"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func AddQrCode(c echo.Context) error {
	var qrCode app.QRCode
	var newQrCode app.QRCode
	var err error

	c.Bind(&qrCode)
	if reflect.DeepEqual(app.QRCode{}, qrCode) {
		return c.String(http.StatusBadRequest, "QR Code cannot be empty")
	}

	newQrCode = qrCode
	newQrCode.Id = uuid.NewString()

	err = app.CreateQR(&newQrCode)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Error creating QR Code")
	}

	return c.JSON(http.StatusOK, qrCode)
}

func RetrieveAllQrCodes(c echo.Context) error {
	return c.JSON(http.StatusOK, app.RetrieveAllQrCodes())
}
