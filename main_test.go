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

func TestCreateQR(t *testing.T) {
	location1 = app.Location{
		Path:        uuid.NewString(),
		Address:     "Av Calle 2",
		Description: "En el 2do piso del edificio",
	}
	location1 = app.Location{
		Path:        uuid.NewString(),
		Address:     "Av Calle 3",
		Description: "En el baño del bar chelero",
	}
	qrCode = &app.QRCode{
		URL:       "https://google.com",
		CreatedAt: time.Now(),
		Locations: []app.Location{location1, location2},
	}
	var err = app.CreateQR(qrCode)
	if err != nil {
		t.Errorf("error creating qr: %s", err.Error())
	}
	if qrCode.URL != "https://google.com" {
		t.Errorf("error creating qr: verify URL")
	} else {
		fmt.Print("[OK] QR Code created successfully!\n")
	}
}

func TestIncrementViews(t *testing.T) {
	var expected = 1
	location1.IncrementView()
	if location1.Analytics.Views != expected {
		t.Errorf("[ERROR] expected %d got %d", expected, location1.Analytics.Views)
	} else {
		fmt.Print("[OK] incremented +1 View\n")
	}
}

func TestAddLocation(t *testing.T) {
	var err error
	var newLocation = app.Location{
		Address:     "Av Siempre Viva 1234",
		Description: "El bar de la esquina, frente a la entrada principal",
		Path:        uuid.NewString(),
	}
	var expected = 3
	err = qrCode.AddLocation(newLocation)
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
