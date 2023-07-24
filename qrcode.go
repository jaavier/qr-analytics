package main

import (
	"fmt"
	"reflect"
)

func CreateQR(qr *QRCode) error {
	if reflect.DeepEqual(QRCode{}, qr) {
		return fmt.Errorf("error: qr incorrect")
	} else {
		qrCodes[(*qr).URL] = *qr
	}
	return nil
}

func NewQRCode() *QRCode {
	return &QRCode{}
}
