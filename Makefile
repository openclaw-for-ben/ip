REGISTRY ?= docker.io
IMAGE ?= bborbe/ip
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)

default: precommit

include Makefile.precommit

.PHONY: build
build:
	DOCKER_BUILDKIT=1 \
	docker build \
	--rm=true \
	--platform=linux/amd64 \
	-t $(REGISTRY)/$(IMAGE):$(BRANCH) \
	-f Dockerfile .

.PHONY: upload
upload:
	docker push $(REGISTRY)/$(IMAGE):$(BRANCH)

.PHONY: clean
clean:
	docker rmi $(REGISTRY)/$(IMAGE):$(BRANCH) || true

.PHONY: buca
buca: build upload clean

deps:
	go mod tidy
	go mod verify

run:
	go run . -listen=:8080 -logtostderr -v=2

.PHONY: default build upload clean buca deps run
