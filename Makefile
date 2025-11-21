.PHONY: build run push

REGISTRY=ghcr.io
REPO=$(shell git config --get remote.origin.url | sed 's/.*github.com[:/]\(.*\)\.git/\1/' | tr '[:upper:]' '[:lower:]')
IMAGE=$(REGISTRY)/$(REPO)

build:
	docker build -t $(IMAGE):latest .

run:
	docker run --rm -it -p 8080:8080 $(IMAGE):latest

push:
	docker push $(IMAGE):latest
