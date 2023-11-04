#/bin/sh

echo "Building the executable for linux"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o executable

echo "Building the docker image"
docker build . -t cchampou/go:latest

echo "Pushing the docker image"
docker push cchampou/go:latest
