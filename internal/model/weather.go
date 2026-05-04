package model

import "time"

type StationInfo struct {
	ID          string `json:"id"`
	StationName string `json:"stationName"`
}

type WeatherRequest struct {
	StationID string  `json:"station_id"`
	Pressure  uint32  `json:"pressure" binding:"required"`
	Humidity  uint16  `json:"humidity" binding:"required"`
	Temp      int16   `json:"temp" binding:"required"`
	Lat       float64 `json:"lat" binding:"required"`
	Long      float64 `json:"long" binding:"required"`
}

type WeatherResponse struct {
	ID        string    `json:"id"`
	Pressure  uint32    `json:"pressure"`
	Humidity  uint16    `json:"humidity"`
	Temp      int16     `json:"temp"`
	StationID string    `json:"station_id"`
	Lat       float64   `json:"lat"`
	Long      float64   `json:"long"`
	CreatedAt time.Time `json:"created_at"`
}
