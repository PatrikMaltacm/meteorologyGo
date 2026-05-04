package model

import "time"

type CreateStationResponse struct {
	Lat  float64 `json:"lat" binding:"required"`
	Long float64 `json:"long" binding:"required"`
}

type StationResponse struct {
	ID        string    `json:"id" `
	Lat       float64   `json:"lat"`
	Long      float64   `json:"long"`
	CreatedAt time.Time `json:"created_at"`
}
