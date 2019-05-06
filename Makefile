all: build docker docker-push
dep:
	go get github.com/appleboy/drone-facebook/template
	go get github.com/joho/godotenv
	go get github.com/urfave/cli
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o rocket
docker:
	docker build -t mike1pol/drone-rocket:latest .
docker-push:
	docker push mike1pol/drone-rocket:latest
