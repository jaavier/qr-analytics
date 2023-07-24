package main

import "time"

type QRCode struct {
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	Analytics Analytics `json:"analytics"`
}

type Analytics struct {
	Views     int     `json:"views"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
