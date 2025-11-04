package models

import(
	"time"
)


type Heartbeat struct {
	SentAt 		time.Time  		   `json:"sent_at" binding:"required"`
}