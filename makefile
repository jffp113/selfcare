#SHELL := /bin/bash

#DOCKER_IMAGE_NAME=selfcare
#VERSION=0.0.1
#DOCKER_FULL_IMAGE_NAME=jffp113/${DOCKER_IMAGE_NAME}

# ==============================================================================
# Building containers


# ==============================================================================
# Building go files
build: clean
	mkdir -p target
	GOOS=linux GOARCH=amd64 go build -o target/selfcare-linux-amd64 app/main.go
	GOOS=linux GOARCH=arm64 go build -o target/selfcare-linux-arm64 app/main.go
	GOOS=windows GOARCH=amd64 go build -o target/selfcare-windows-amd64.exe app/main.go
	GOOS=darwin GOARCH=amd64 go build -o target/selfcare-dawrin-amd64 app/main.go
	GOOS=darwin GOARCH=arm64 go build -o target/selfcare-dawrin-arm64 app/main.go

clean:
	rm -rf target