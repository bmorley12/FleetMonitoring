package services

import (
	"example/FleetMonitoring/internal/models"
	"example/FleetMonitoring/internal/repository"
)


// Records heartbeat to stored device data
func RecordHeartBeat(deviceID string, req models.Heartbeat) error {
	if !repository.EnsureDeviceExists(deviceID) {
		return ErrDeviceNotFound
	}

	device := repository.GetOrCreateDevice(deviceID)		// retrieve device

	if device.FirstHeartbeat.IsZero() {									// record the first heartbeat
		device.FirstHeartbeat = req.SentAt
	} 
	
	device.LastHeartbeat = req.SentAt										// store heartbeat
	device.SumHeartbeats += 1														// increment number of heartbeats


	return nil
}