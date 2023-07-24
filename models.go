package main

import (
	"fmt"
	"reflect"
	"time"
)

type QRCode struct {
	URL       string     `json:"url"` // podría ser map[string]QRCode??
	CreatedAt time.Time  `json:"created_at"`
	Location  []Location `json:"location"`
}

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

var qrCodes = make(map[string]QRCode)

func RetrieveAllQrCodes() map[string]QRCode {
	return qrCodes
}

type Location struct {
	Path        string    `json:"path"` // uuid v4
	Address     string    `json:"address"`
	Description string    `json:"description"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Analytics   Analytics `json:"analytics"`
	// mutex       *sync.Mutex
}

func (l *Location) IncrementView() {
	l.Analytics.Views++
}

type Analytics struct {
	Views     int     `json:"views"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
