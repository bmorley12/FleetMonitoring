package main

import (
	"github.com/gin-gonic/gin"

	"example/FleetMonitoring/internal/api"
	"example/FleetMonitoring/internal/repository"
)


// ----------------- Main -------------------------
func main(){
	repository.GetValidDevices("./data/devices.csv")		// initialize list of valid devices
	
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()									// initialize router
	api.RegisterRoutes(router)
	router.Run(":6733")

}
