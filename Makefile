all: build docker docker-push
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o rocket
docker:
	docker build -t mike1pol/drone-rocket:latest .
docker-push:
	docker push mike1pol/drone-rocket:latest
