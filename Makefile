# vi: ft=make
.PHONY: run proto test benchmark
run:
	go run main.go

proto:
	protoc -I email_service/ email_service/email_service.proto --go_out=plugins=grpc:email_service

test:
	go test -v ./...

benchmark:
	go test -bench=./... -benchmem -benchtime 10s

docker:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o build/email_service .
	docker build . -t lileio/email_service:latest
