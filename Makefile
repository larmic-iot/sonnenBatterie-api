CONTAINER_NAME=larmic-sonnen-batterie-api
IMAGE_NAME=larmic/sonnen-batterie-api
IMAGE_TAG=latest

SONNENBATTERIE_IP=10.0.40.100
SONNENBATTERIE_USER_PASSWORD=my-password

help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## —— Build 🏗️———————————————————————————————————————————————————————————————————————————————————————————————————
docker-build: ## Builds docker image including automated tests
	@echo "Remove docker image if already exists"
	docker rmi -f ${IMAGE_NAME}:${IMAGE_TAG}
	@echo "Build go docker image"
	DOCKER_BUILDKIT=1 docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .
	@echo "Prune intermediate images"
	docker image prune --filter label=stage=intermediate -f

## —— Run application 🏃🏽————————————————————————————————————————————————————————————————————————————————————————
docker-run: ## Runs docker container. Set parameter SONNENBATTERIE_IP and SONNENBATTERIE_USER_PASSWORD
	@echo "Run docker container"
	docker run -d -p 8080:8080 --rm -e SONNENBATTERIE_IP='${SONNENBATTERIE_IP}' -e SONNENBATTERIE_USER_NAME='User' -e SONNENBATTERIE_USER_PASSWORD='${SONNENBATTERIE_USER_PASSWORD}' --name ${CONTAINER_NAME} ${IMAGE_NAME}

docker-logs: ## Prints logs of running container
	@echo "Logging..."
	docker logs -f ${CONTAINER_NAME}

docker-stop: ## Stops running docker container
	@echo "Stop docker container"
	docker stop ${CONTAINER_NAME}
