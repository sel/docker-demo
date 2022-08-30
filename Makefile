PROJECT := docker-demo
REPO := docker.io/slarkin/$(PROJECT)
TAG := v4
LISTEN_PORT := 8000

.DEFAULT_GOAL := build

.PHONY: push
push: build
	docker push $(REPO):$(TAG)

.PHONY: build
build:
	docker build --tag $(REPO):$(TAG) .

.PHONY: run
run: build stop
	docker run \
	    --name $(PROJECT) --detach --rm \
	    --env LISTEN_PORT=$(LISTEN_PORT) --publish $(LISTEN_PORT):$(LISTEN_PORT) \
	    $(REPO):$(TAG)
	docker ps

.PHONY: stop
stop:
	docker stop $(PROJECT) 2>/dev/null || true
