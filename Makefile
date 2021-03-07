CONTAINER_NAME=sonnen-batterie-api
IMAGE_NAME=larmic/sonnen-batterie-api
IMAGE_TAG=latest

docker-all: docker-build

docker-build:
	@echo "Remove docker image if already exists"
	docker rmi -f ${IMAGE_NAME}:${IMAGE_TAG}
	@echo "Build go docker image"
	DOCKER_BUILDKIT=1 docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .
	@echo "Prune intermediate images"
	docker image prune --filter label=stage=intermediate -f

docker-remove-dangling:
	@echo "Remove dangling images"
	docker rmi -f $$(docker images --filter "dangling=true" -q --no-trunc)
