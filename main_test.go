package main

import (
	"fmt"
	"qranalytics/app"
	"testing"
	"time"

	"github.com/google/uuid"
)

var qrCode = app.NewQRCode()
var location1 app.Location
var location2 app.Location
var location3 app.Location

func init() {
	location1 = app.Location{
		Path:        uuid.NewString(),
		Address:     "Av Calle 2",
		Description: "En el 2do piso del edificio",
	}
	location2 = app.Location{
		Path:        uuid.NewString(),
		Address:     "Av Calle 3",
		Description: "En el baño del bar chelero",
	}
	location3 = app.Location{
		Address:     "Av Las Perdices 3333",
		Description: "Cafetería estación de buses",
		Path:        uuid.NewString(),
	}
	qrCode = &app.QRCode{
		URL:       "https://google.com",
		CreatedAt: time.Now(),
		Locations: []app.Location{location1, location2},
	}
}

func TestCreateQR(t *testing.T) {
	var expected = "https://google.com"
	var got app.QRCode
	var err = app.CreateQR(qrCode)
	var ok bool

	if got, ok = app.QrCodes[qrCode.URL]; !ok {
		t.Errorf("error creating qr: %s", err.Error())
	} else {
		if got.URL != expected {
			t.Errorf("error creating qr: verify URL")
		} else {
			fmt.Print("[OK] QR Code created successfully!\n")
		}
	}
}

func TestIncrementViews(t *testing.T) {
	var expected = 1
	var got = 0

	location1.IncrementView()
	got = location1.Analytics.Views

	if got != expected {
		t.Errorf("[ERROR] expected %d got %d", expected, got)
	} else {
		fmt.Print("[OK] incremented +1 View\n")
	}
}

func TestAddLocation(t *testing.T) {
	var err error
	var expected = 3
	err = qrCode.AddLocation(location3)

	if err != nil {
		t.Errorf("error adding location: %s", err.Error())
	} else {
		got := len(qrCode.RetrieveLocations())
		if got != expected {
			t.Errorf("expected %d got %d", expected, got)
		} else {
			fmt.Print("[OK] Location added successfully\n")
		}
	}
}
