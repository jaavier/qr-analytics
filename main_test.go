package main

import "testing"

func TestCreateQR(t *testing.T) {
	var err = CreateQR(QRCode{
		URL: "https://google.com",
	})
	if err != nil {
		t.Errorf("error creating qr: %s", err.Error())
	}
}
