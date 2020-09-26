#!make
export GOPRIVATE=github.com/maxidelgado

init: deps
	@echo "setting up git..."
	@git init
	@git add .
	@git commit -m "First commit"
	@git remote add origin https://github.com/maxidelgado/skeleton-go.git
	@git push -u origin master

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/api main.go

deploy: clean build
	sls deploy --verbose

deps:
	@echo "vendoring dependencies..."
	@go mod vendor
