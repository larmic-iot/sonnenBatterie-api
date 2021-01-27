CONTAINER_NAME=sonnen-batterie-api
IMAGE_NAME=larmic/sonnen-batterie-api
IMAGE_TAG=latest

run:
	go run main.go

docker-all: docker-build

docker-build:
	@echo "Remove docker image if already exists"
	docker rmi -f ${IMAGE_NAME}:${IMAGE_TAG}
	@echo "Build go docker image"
	DOCKER_BUILDKIT=1 docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .
	@echo "Prune intermediate images"
	docker image prune --filter label=stage=intermediate -f

docker-run:
	docker run -d -p 8080:8080 --rm --name ${CONTAINER_NAME} ${IMAGE_NAME}

docker-stop:
	docker stop ${CONTAINER_NAME}

docker-remove-dangling:
	@echo "Remove dangling images"
	docker rmi -f $$(docker images --filter "dangling=true" -q --no-trunc)
