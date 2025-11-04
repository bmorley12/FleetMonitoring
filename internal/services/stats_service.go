package services

import (
	"time"
	"example/FleetMonitoring/internal/models"
	"example/FleetMonitoring/internal/repository"
)


// Record device statistics to stored device data
func RecordStats(deviceID string, req models.UploadedStats) error{
	if !repository.EnsureDeviceExists(deviceID) {			// Verify device is valid
		return ErrDeviceNotFound
	}

	device := repository.GetOrCreateDevice(deviceID)									// Get device data
	device.UploadTimes = append(device.UploadTimes, req.UploadTime)		// add new upload times
	device.TotalUploadTime += req.UploadTime													// add to total upload time

	return nil
}


// Calculate the device statistics and return results or error
func CalculateStats(deviceID string) (models.DeviceStats, error){
	if !repository.EnsureDeviceExists(deviceID) {			// verify device is valid
		return models.DeviceStats{}, ErrDeviceNotFound
	}

	device := repository.GetOrCreateDevice(deviceID)		// Get device data

	// If device exists with no data, return no data and the corresponding error
	if len(device.UploadTimes) == 0{										
		return models.DeviceStats{}, ErrNoDeviceStats
	}

	// Calculate avg upload time
	avg := time.Duration(device.TotalUploadTime / int64(len(device.UploadTimes)))


	// Calculate uptime percentage
	elapsedTime := device.LastHeartbeat.Sub(device.FirstHeartbeat)
	uptime := float64(device.SumHeartbeats) / float64(elapsedTime.Minutes()) * 100

	// Build response
	response := models.DeviceStats{
		AvgUploadTime: avg.String(),
		Uptime:        uptime,
	}

	return response, nil			// return successful results
}