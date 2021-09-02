package models

type Asset struct {
	Ticker          string  `json:"Ticker"`
	Name            string  `json:"Name"`
	Type            string  `json:"Type"`
	Sector          string  `json:"Sector"`
	Price           float64 `json:"Price"`
	Dy              string  `json:"DY"`
	Max             float64 `json:"Max"`
	Min             float64 `json:"Min"`
	DPA             float64 `json:"DPA"`
	FairPrice       float64 `json:"FairPrice"`
	WithinFairPrice bool    `json:"WithinFairPrice"`
}
