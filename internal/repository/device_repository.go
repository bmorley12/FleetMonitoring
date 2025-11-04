package repository


import (
	"sync"
	"encoding/csv"
	"log"
	"os"
	
	"example/FleetMonitoring/internal/models"
)


var (
	deviceStore = make(map[string]*models.DeviceData)					// Stores devices with data
	ValidDevices = make(map[string]bool)											// Stores a list of valid devices
	storeLock   sync.Mutex																		// Mutex lock
)


// Quick function to check if device exists
func EnsureDeviceExists(deviceID string) bool {
	return ValidDevices[deviceID]
}

// This function reads a csv file at a given path and populates 
// a list of valid devices
func GetValidDevices(path string) {
	// Open file
	file, err := os.Open(path)
	Check(err, "Failed to open file")
	defer file.Close()
	reader := csv.NewReader(file)

	// Check header
	header , err := reader.Read()
	Check(err, "Fialed to read header")
	if header[0] != "device_id"{
		log.Fatal("Wrong CSV header. Please double check file")
	}

	// Read file
	records, err := reader.ReadAll()
	Check(err, "Fialed to read data")

	// Create list of valid devices
	for _, record := range records {
		ValidDevices[record[0]] = true
		// log.Println(record[0], ValidDevices[record[0]])
	}

}

// Quick error checking function
func Check(err error, message string){
	if err != nil{
		log.Fatalf("%v: %v", message, err)
	}
}

// If a device already exists, this function returns that device
// if a device does not exist, it is created and then returned
func GetOrCreateDevice(deviceID string) *models.DeviceData{
	storeLock.Lock()			// Apply mutex lock
	defer storeLock.Unlock()

	device, exists := deviceStore[deviceID]
	if !exists {		// creates device info if it hasn't been initialized
		device = &models.DeviceData{}
		deviceStore[deviceID] = device
	}

	return device
}


