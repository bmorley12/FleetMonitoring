package api


import (
	"github.com/gin-gonic/gin"
	"example/FleetMonitoring/internal/api/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	devices := r.Group("/api/v1/devices")	// api grouping
	{
		devices.POST("/:device_id/heartbeat", handlers.PostHeartBeat)
		devices.POST("/:device_id/stats", handlers.PostStats)
		devices.GET("/:device_id/stats", handlers.GetStats)
	}

	
}