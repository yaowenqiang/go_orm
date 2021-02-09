tidy:
	go mod tidy
startmq:
	docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
co:
	go run distributed/coordinator/exec/main.go
sensor:
	go run distributed/sensors/sensor.go
data:
	go run distributed/datamanager/exec/main.go

