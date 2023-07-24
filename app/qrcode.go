package app

import (
	"fmt"
	"reflect"
	"time"
)

type QRCode struct {
	URL       string     `json:"url"`
	CreatedAt time.Time  `json:"created_at"`
	Location  []Location `json:"location"`
}

var QrCodes = make(map[string]QRCode)

func (qc *QRCode) RetrieveLocations() []Location {
	return qc.Location
}

func (qc *QRCode) AddLocation(location Location) error {
	if reflect.DeepEqual(Location{}, location) {
		return fmt.Errorf("you cannot add location empty")
	}
	qc.Location = append(qc.Location, location)
	return nil
}

func CreateQR(qr *QRCode) error {
	if reflect.DeepEqual(QRCode{}, qr) {
		return fmt.Errorf("error: qr incorrect")
	} else {
		QrCodes[(*qr).URL] = *qr
	}
	return nil
}

func NewQRCode() *QRCode {
	return &QRCode{}
}

func RetrieveAllQrCodes() map[string]QRCode {
	return QrCodes
}
