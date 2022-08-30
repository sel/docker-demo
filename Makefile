TAG := v3

.PHONY: push
push:
	docker build -t docker.io/slarkin/docker-demo:$(TAG) .
	docker push docker.io/slarkin/docker-demo:$(TAG)
