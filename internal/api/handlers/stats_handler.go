package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"

	"example/FleetMonitoring/internal/models"
	"example/FleetMonitoring/internal/services"
)


// POST: api/v1/:device_id/stats
func PostStats(c *gin.Context) {
	deviceID := c.Param("device_id")

	// Checks request for required parameters 
	var req models.UploadedStats
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)							// return 400
		return
	}

	// Attempts to record stats, handles errors for failures
	if err := services.RecordStats(deviceID, req); err != nil {
		if err == services.ErrDeviceNotFound {
			c.Status(http.StatusNotFound)							// return 404
			return
		}
		c.Status(http.StatusInternalServerError)		// return 500
		return
	}

	c.Status(http.StatusNoContent)								// return 204
}



func GetStats(c *gin.Context) {
	deviceID := c.Param("device_id")
	
	// Calculate device statistics
	response, err := services.CalculateStats(deviceID)

	// Handle possible errors
	if err != nil {
		if err == services.ErrNoDeviceStats {
			c.Status(http.StatusNoContent)						// return 204
			return
		} else if err == services.ErrDeviceNotFound {
			c.Status(http.StatusNotFound)							// return 404
			return
		}

		c.Status(http.StatusInternalServerError)		// return 500
		return
	}


	c.JSON(http.StatusOK, response)								// return 200
}