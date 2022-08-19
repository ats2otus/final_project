APP_NAME?=api
DATE?=`date "+%FT%H:%M:%S"`
IMAGE_REGISTRY?=github.com/ats2otus/final_project
IMAGE_VERSION?=$(shell git rev-parse --short=8 HEAD)
IMAGE_TAG?=${IMAGE_REGISTRY}:${IMAGE_VERSION}

docs:
	swag init --ot json -g ./cmd/api/routes.go

build:
	CGO_ENABLED=0 go build \
		-ldflags="-s -w" \
		-trimpath \
		-o ./bin/${APP_NAME} \
		./cmd/${APP_NAME}

image:
	docker build -t ${IMAGE_TAG} .

run:image
	IMAGE_TAG=${IMAGE_TAG} docker-compose up --build -d

stop:
	IMAGE_TAG=${IMAGE_TAG} docker-compose down

.PHONY: stop run image build docs
