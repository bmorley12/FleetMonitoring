package models

import(
	"time"
)


type DeviceData struct {
	FirstHeartbeat 				time.Time
	LastHeartbeat 				time.Time
	SumHeartbeats 				int64
	TotalUploadTime 			int64
	UploadTimes  				  []int64
}