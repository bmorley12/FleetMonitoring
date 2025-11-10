# FleetMonitoring

## Running the app
Can be run through a docker container:
 
```bash
docker compose build

docker compose up
```


or with
```bash
go run ./cmd/FleetMonitoring/main.go
```

## Testing

Testing can be done by running any of the executables in `./device_simulators`. Please choose the correct one for your operating system.

Standalone tests can be achieved with the following example curl commands:

```bash
curl \
  -X POST \
  -H 'content-type: application/json' \
  -d '{"sent_at": "2024-04-02T16:58:00Z"}' \
  'http://127.0.0.1:6733/api/v1/devices/60-6b-44-84-dc-64/heartbeat'

curl \
  -X POST \
  -H 'content-type: application/json' \
  -d '{"sent_at": "0001-01-01T00:00:00Z", "upload_time": 223543506424}' \
  'http://127.0.0.1:6733/api/v1/devices/60-6b-44-84-dc-64/stats'

curl \
  -H 'content-type: application/json' \
  'http://127.0.0.1:6733/api/v1/devices/60-6b-44-84-dc-64/stats'

```


## FAQ

**How long did you spend working on the problem? What did you find to be the most difficult part?**

It took about 3 hours. Most of the appp was working quite quickly, however I had an issue with the POST/stats endpoint returning 400 when the device simulator was sending requests, but 204 when I was sending requests from postman. I eventually found that the device simulator sends its "sent_at" field with the placeholder info "0001-01-01T00:00:00Z" which is considered 0 in the Time go library, and doesn't pass the required binding placed on that field. Troubleshooting that took most of my time.


**How would you modify your data model or code to account for more kinds of metrics**

I would most likely increase the number of fields within POST/stats, and increase the data stored in my DeviceData struct. 


**Discuss your solution's runtime complexity**

Reading from CSV - O(n)

POST/heartbeat - O(1)

POST/stats - O(1)

GET/stats - O(1)

I tried to stay away from O(n) (or worse) runtimes as much as possible, considering this project is likely being scaled to include thousands of devices and tonnes of data. 
