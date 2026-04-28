package request

type WheatherRequest struct {
	Pressure uint32 `json:"pressure"`
	Humidity uint16 `json:"humidity"`
	Temp     int16  `json:"temp"`
}
