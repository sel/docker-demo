PROJECT := docker-demo
REPO := docker.io/slarkin/$(PROJECT)
GO_VERSION := 1.20
LISTEN_PORT := 8000

GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_UNTRACKED_CHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(GIT_UNTRACKED_CHANGES),)
    GIT_COMMIT := $(GIT_COMMIT)-dirty
endif

BUILD_DATE_UTC := $(shell date -u +"%y%m%d%H%M")
PUBLISHED_TAG := $(GIT_COMMIT)-go$(GO_VERSION)-$(BUILD_DATE_UTC)

.DEFAULT_GOAL := push

.PHONY: push
push: build
	docker tag $(REPO):$(GIT_COMMIT) $(REPO):$(PUBLISHED_TAG)
	docker push $(REPO):$(PUBLISHED_TAG)

.PHONY: build
build:
	docker build . \
	    --squash \
	    --build-arg GO_VERSION=$(GO_VERSION) \
	    --build-arg VERSION=$(GIT_COMMIT) \
	    --tag $(REPO):$(GIT_COMMIT)

.PHONY: run
run: build stop
	docker run \
	    --name $(PROJECT) --detach --rm \
	    --env LISTEN_PORT=$(LISTEN_PORT) --publish $(LISTEN_PORT):$(LISTEN_PORT) \
	    $(REPO):$(GIT_COMMIT)
	docker ps

.PHONY: stop
stop:
	docker stop $(PROJECT) 2>/dev/null || true
