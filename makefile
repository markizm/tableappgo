# This makefile provides some easy targets to use while developing Docker
# containers.

SUDO = sudo

REPO = sre
IMAGE = release 
VERSION = $(shell cat VERSION)
VERSION ?= latest
CONTAINER = release1
PORTS = -p 9011 

# Explicitly tell Make that these are not actual files to be built. This
# ensures that they will always run when called regardless of the filesystem.
#.PHONY: build tag run watch start stop restart rm reload enter logs push pushtag deploy

build:
	$(SUDO) docker build -t $(IMAGE) .
watch:
	$(SUDO) docker run --dt -p $(PORTS) --NET=host --name $(CONTAINER) $(IMAGE)
start:
	$(SUDO) docker start $(CONTAINER)
stop:
	$(SUDO) docker stop $(CONTAINER)
restart:
	$(SUDO) docker restart $(CONTAINER)
rm:
	$(SUDO) docker rm $(CONTAINER)
#pipe ensures these 4 targets happen in this order
reload: | build stop rm run
enter:
	$(SUDO) docker exec -it $(CONTAINER) /bin/bash
logs:
	$(SUDO) docker logs -f $(CONTAINER)
push:
	$(SUDO) docker push $(NS)/$(REPO)/$(IMAGE)
pushtag:
	$(SUDO) docker push $(NS)/$(REPO)/$(IMAGE):$(VERSION)

default: build
