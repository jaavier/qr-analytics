package app

type Analytics struct {
	Views     int     `json:"views"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
