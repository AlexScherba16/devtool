# devtool
USB devices network scanner

# How to run application
- clone project https://github.com/AlexScherba16/devtool
- cd devtool
- protoc --go_out=. --proto_path=./proto -I ./proto ./proto/*.proto
- go mod init devtool
- go mod tidy
- go run cmd/devtool/main.go

If build process failed
- CGO_CFLAGS="-Wno-nullability-completeness" go run cmd/devtool/main.go

# How to get connected devices
- curl http://localhost:8080/api/v1/devices 
