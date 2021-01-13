APP?=wedding-invitation-go

RELEASE?=$(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null || echo "latest")
COMMIT?=$(shell git rev-parse --short HEAD)
BUILT?=$(shell date +"%Y-%m-%d_%H:%M:%S_%Z")

REGISTRY_PREFIX?=dystargate
CONTAINER_IMAGE?=${REGISTRY_PREFIX}/${APP}:${RELEASE}

GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

.PHONY: vendor
vendor:
	go mod vendor

fmt:
	@gofmt -w ${GOFILES}

image:
	docker build -t ${CONTAINER_IMAGE} .

push: image
	docker push ${CONTAINER_IMAGE}