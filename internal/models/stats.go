package models

import(
	"time"
)

type UploadedStats struct {
	SentAt     				time.Time				 `json:"sent_at"`		// binding: "required" removed because 0 time is an error
	UploadTime 				int64    				 `json:"upload_time" binding:"required"`
}

type DeviceStats struct {
	AvgUploadTime 		string 					 `json:"avg_upload_time"`
	Uptime        		float64					 `json:"uptime"`
}
