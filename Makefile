export PROTO_DIR := $(PWD)/proto
IMG_TAG ?= local
ACR_ID ?= local.io
CLIENT_API_IMG_NAME ?= $(ACR_ID)/clientapi
PORT_DOMAIN_IMG_NAME ?= $(ACR_ID)/portdomainservice

generate:
	go generate ./...

lint:
	golangci-lint run --timeout=5m

docker/clientApi:
	docker build -t $(CLIENT_API_IMG_NAME):$(IMG_TAG) -t $(CLIENT_API_IMG_NAME):latest -f Dockerfile services/clientapi

docker/portDomain:
	docker build -t $(PORT_DOMAIN_IMG_NAME):$(IMG_TAG) -t $(PORT_DOMAIN_IMG_NAME):latest -f Dockerfile services/portDomainService
