package acquisition

import "time"

// Datum TODO
type Datum struct {
	Time     time.Time `json:"time"`
	SensorID uint64    `json:"sensor_id"`
	X        int32     `json:"x"`
	Y        int32     `json:"y"`
	Z        int32     `json:"z"`
	T        int16     `json:"t"`
}
