docker-build-buyer:
	docker build -t $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_BUYER):$(DOCKER_TAG) -f ./buyer/Dockerfile ./buyer

