## soccer-gateway

### Authors 
- Mohammad Wildan Kosasih <mwkosasih@gmail.com>

### About
Service edge (gateway) is a single mediation that handles the requests for client. Use restful API. 

### How
How to run this project :
- clone this project into go path
- run in terminal `go run main.go`

### Structure
```
├── domain
├── proto
├── route
├── util             
├── main.go
```
- Domain: package for handling request, hit service use grpc, build response
- Proto: folder to collect protobuf file
- Route: configuration of route
- Util: common app process / helper