package main

import (
	"encoding/csv"
	// "fmt"
	"log"
	"os"
	"time"
	"sync"
	"net/http"


	"github.com/gin-gonic/gin"
)


// ----------------- Data Structures -------------------------
type HeartbeatRequest struct {
	SentAt time.Time `json:"sent_at" binding:"required"`
}

type UploadStatsRequest struct {
	SentAt     time.Time `json:"sent_at"`		// binding: "required" removed because 0 time is an error
	UploadTime int64     `json:"upload_time" binding:"required"`
}

type GetDeviceStatsResponse struct {
	AvgUploadTime string  `json:"avg_upload_time"`
	Uptime        float64 `json:"uptime"`
}


type DeviceData struct {
	FirstHeartbeat time.Time
	LastHeartbeat time.Time
	SumHeartbeats int64
	TotalUploadTime int64
	UploadTimes   []int64
}


// ----------------- Global Variables -------------------------
var (
	validDevices = make(map[string]bool)
	deviceStore = make(map[string]*DeviceData)
	storeLock   sync.Mutex
)


// ----------------- Main -------------------------
func main(){
	readCSV("devices.csv")		// read list of valid devices

	router := gin.Default()		// initialize router

	api := router.Group("/api/v1")	// api grouping
	{
		api.POST("/devices/:device_id/heartbeat", postHeartbeat)
		api.POST("/devices/:device_id/stats", postStats)
		api.GET("/devices/:device_id/stats", getStats)
	}

	router.Run(":6733")
}


// ----------------- End points -------------------------
func postHeartbeat(c *gin.Context) {
	deviceID := c.Param("device_id")

	// Check for device in list of accepted devices
	if !ensureDeviceExists(deviceID) {
		c.Status(http.StatusNotFound)		// return 404
		return
	}

	// Checks request for required parameters 
	var req HeartbeatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)		// return 400
		return
	}

	storeLock.Lock()			// Apply mutex lock
	defer storeLock.Unlock()

	device, exists := deviceStore[deviceID]
	if !exists {		// creates device info if it hasn't been initialized
		device = &DeviceData{}
		deviceStore[deviceID] = device
	}

	if device.FirstHeartbeat.IsZero() {			// record the first heartbeat
		device.FirstHeartbeat = req.SentAt
	} 
	
	device.LastHeartbeat = req.SentAt		// store heartbeat
	device.SumHeartbeats += 1

	c.JSON(204, gin.H{"description": "the request was completed successfully"})
}

func postStats(c *gin.Context) {
	deviceID := c.Param("device_id")

	// Check for device in list of accepted devices
	if !ensureDeviceExists(deviceID) {
		c.Status(http.StatusNotFound)		// return 404
		return
	}

	// Checks request for required parameters 
	var req UploadStatsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)		// return 400
		return
	}

	storeLock.Lock()
	defer storeLock.Unlock()

	device, exists := deviceStore[deviceID]
	if !exists {	// creates device info if it hasn't been initialized
		device = &DeviceData{}
		deviceStore[deviceID] = device
	}

	device.UploadTimes = append(device.UploadTimes, req.UploadTime)
	device.TotalUploadTime += req.UploadTime
	c.Status(http.StatusNoContent)
}



func getStats(c *gin.Context) {
	deviceID := c.Param("device_id")

	// Check for device in list of accepted devices
	if !ensureDeviceExists(deviceID) {
		c.Status(http.StatusNotFound)		// return 404
		return
	}

	// Get Device information
	storeLock.Lock()
	device, exists := deviceStore[deviceID]
	storeLock.Unlock()

	// If the device is valid, but no information has been received
	if !exists || len(device.UploadTimes) == 0 {
		c.Status(http.StatusNoContent)		// return 204
		return
	}

	// Calculate avg upload time
	avg := time.Duration(device.TotalUploadTime / int64(len(device.UploadTimes)))


	// Calculate uptime percentage
	elapsedTime := device.LastHeartbeat.Sub(device.FirstHeartbeat)
	uptime := float64(device.SumHeartbeats) / float64(elapsedTime.Minutes()) * 100

	// Return results
	c.JSON(http.StatusOK, GetDeviceStatsResponse{
		AvgUploadTime: avg.String(),
		Uptime:        uptime,
	})
}



// ----------------- Helper Functions -------------------------
func check(err error, message string){
	if err != nil{
		log.Fatalf("%v: %v", message, err)
	}
}


func readCSV(fileName string) {
	// Open file
	file, err := os.Open(fileName)
	check(err, "Failed to open file")
	defer file.Close()
	reader := csv.NewReader(file)

	// Check header
	header , err := reader.Read()
	check(err, "Fialed to read header")
	if header[0] != "device_id"{
		log.Fatal("Wrong CSV header. Please double check file")
	}

	// Read file
	records, err := reader.ReadAll()
	check(err, "Fialed to read data")

	// Create list of valid devices
	for _, record := range records {
		validDevices[record[0]] = true
	}

}

// Quick function to check if device exists
func ensureDeviceExists(deviceID string) bool {
	return validDevices[deviceID]
}
