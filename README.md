# sonnenBatterie-api

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Go build](https://github.com/larmic/sonnenBatterie-api/workflows/Go%20build/badge.svg)
![Docker build and push](https://github.com/larmic/sonnenBatterie-api/workflows/Docker%20build%20and%20push/badge.svg)
[![Docker hub image](https://img.shields.io/docker/image-size/larmic/sonnen-batterie-api?label=dockerhub)](https://hub.docker.com/repository/docker/larmic/sonnen-batterie-api)
![Docker Image Version (latest by date)](https://img.shields.io/docker/v/larmic/sonnen-batterie-api)

TODO: 
* implement this tool
* write some intro
* battery charge status?

## Versioning

[Semantic Versioning 2.x](https://semver.org/) is used. Version number **MAJOR.MINOR.PATCH** with

* **MAJOR** version increase on incompatible API changes
* **MINOR** version increase on adding new functionality in a backwards compatible manner
* **PATCH** version increase on backwards compatible bug fixes or documentation

## Usage

The easiest way is to use the docker image. Otherwise, the artifact will have to be built by yourself.

```sh 
$ docker pull larmic/sonnenBatterie-api
$ docker run -d -p 8080:8080 --rm --name larmic-sonnenBatterie-api larmic/sonnenBatterie-api
```

## Example requests

```sh 
$ curl http://localhost:8080                        # Open Api 3.1 specification
$ curl http://localhost:8080/10.0.0.1               # General energy plug information
$ curl http://localhost:8080/10.0.0.1/consumption   # Energy consumption
```

## Build application by yourself

### Requirements

* Docker 
* Go 1.15.x (if you want to build it without using docker builder)

### Build it

```sh 
$ make docker-build                             # build local docker image
$ make docker-push                              # push local docker image to hub.docker.com
$ make docker-all                               # build and push docker image to hub.docker.com
$ make IMAGE_TAG="0.0.1" docker-all             # build and push docker image with specific version
```

### Run it native

```sh 
$ make run                                      # start native app 
$ curl http://localhost:8080/api/10.0.0.210     # call rest service
$ ctrl+c                                        # stop native app
```

### Run it using docker

```sh 
$ make docker-run                               # start docker image 
$ curl http://localhost:8080/api/10.0.0.210     # call rest service
$ make docker-stop                              # stop and remove docker app
```