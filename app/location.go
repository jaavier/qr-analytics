package app

type Location struct {
	Path        string    `json:"path"` // uuid v4
	Address     string    `json:"address"`
	Description string    `json:"description"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Analytics   Analytics `json:"analytics"`
}

func (l *Location) IncrementView() {
	l.Analytics.Views++
}
