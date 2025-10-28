# FleetMonitoring
SafelyYou coding challenge


## Running the app
Can be run through a docker container:
 
`$> docker build -t fleet-monitoring-app`

`$> docker run -it --rm -p 6733:6733 --name fleet-monitoring-dock fleet-monitoring-app`


or with

`$> go run main.go`


## Your Questions

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
