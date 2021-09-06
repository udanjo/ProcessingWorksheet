package models

import (
	"time"
)

type Asset struct {
	Ticker          string    `gorm:"size:8;not null;primary_key" json:"Ticker"`
	Name            string    `gorm:"size:250" json:"Name"`
	Type            string    `json:"Type"`
	Sector          string    `json:"Sector"`
	Price           float64   `json:"Price"`
	Dy              string    `json:"DY"`
	Max             float64   `json:"Max"`
	Min             float64   `json:"Min"`
	DPA             float64   `json:"DPA"`
	FairPrice       float64   `json:"FairPrice"`
	WithinFairPrice bool      `json:"WithinFairPrice"`
	CreatedAt       time.Time `json:"created"`
}
