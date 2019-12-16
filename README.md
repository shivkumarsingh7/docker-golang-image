# docker-golang-image

## Build docker image
docker build -t go-app .

## Run Docker in detached mode
docker run -d -p 8080:8080 go-app

## Run Docker
docker run -p 8080:8080 go-app