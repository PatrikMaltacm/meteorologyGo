package model

import "time"

type WheatherRequest struct {
	Pressure uint32 `json:"pressure" binding:"required"`
	Humidity uint16 `json:"humidity" binding:"required"`
	Temp     int16  `json:"temp" binding:"required"`
}

type WheaterResponse struct {
	ID        string    `json:"id"`
	Pressure  uint32    `json:"pressure"`
	Humidity  uint16    `json:"humidity"`
	Temp      int16     `json:"temp"`
	CreatedAt time.Time `json:"created_at"`
}
