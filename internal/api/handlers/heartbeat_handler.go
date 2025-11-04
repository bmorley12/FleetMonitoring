package handlers


import(
	"net/http"
	"github.com/gin-gonic/gin"

	"example/FleetMonitoring/internal/models"
	"example/FleetMonitoring/internal/services"
)


func PostHeartBeat(c *gin.Context) {
	deviceID := c.Param("device_id")


	// Checks request for required parameters 
	var req models.Heartbeat
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)									// return 400
		return
	}


	// Attempts to record heartbeat, handles potential errors
	if err := services.RecordHeartBeat(deviceID, req); err != nil {
		if err == services.ErrDeviceNotFound {
			c.Status(http.StatusNotFound)									// return 404
			return
		}
		c.Status(http.StatusInternalServerError)				// return 500
		return
	}

	c.JSON(204, gin.H{"description": "the request was completed successfully"})			// return 204
}