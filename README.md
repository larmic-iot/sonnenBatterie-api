# sonnenBatterie-api

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Go build](https://github.com/larmic/sonnenBatterie-api/workflows/Go%20build/badge.svg)
![Docker build and push](https://github.com/larmic/sonnenBatterie-api/workflows/Docker%20build%20and%20push/badge.svg)
[![Docker hub image](https://img.shields.io/docker/image-size/larmic/sonnen-batterie-api?label=dockerhub)](https://hub.docker.com/repository/docker/larmic/sonnen-batterie-api)
![Docker Image Version (latest by date)](https://img.shields.io/docker/v/larmic/sonnen-batterie-api)

A REST api client (adapter) for the [SonnenBatterie](https://sonnen.de/stromspeicher/sonnenbatterie-10/). The default
Sonnen-Battery-API (with token) provides too little information. This application does not need the token, but use the 
normal access to the SonnenBatterie. REST endpoints documented in [open api 3.1](open-api-3.yaml).

This project inspired by [tp-link-hs110-api written in go](https://github.com/larmic/tp-link-hs110-api) and
improves my Go knowledge.

## Versioning

[Semantic Versioning 2.x](https://semver.org/) is used. Version number **MAJOR.MINOR.PATCH** with

* **MAJOR** version increase on incompatible API changes
* **MINOR** version increase on adding new functionality in a backwards compatible manner
* **PATCH** version increase on backwards compatible bug fixes or documentation

## Usage

The easiest way is to use the docker image. Otherwise, the artifact will have to be built by yourself.

```sh 
$ docker pull larmic/sonnen-batterie-api
$ docker run -d -p 8080:8080 --rm \
 -e SONNENBATTERIE_IP='<my-battery-ip>' \
 -e SONNENBATTERIE_USER_NAME='User' \
 -e SONNENBATTERIE_USER_PASSWORD='<my-password>' \
 --name larmic-sonnen-batterie-api larmic/sonnen-batterie-api
```

## Example requests

See [open api 3 specification](open-api-3.yaml) for further information.

```sh 
$ curl http://localhost:8080/sonnen-battery-api                    # Open Api 3.1 specification
$ curl http://localhost:8080/sonnen-battery-api/api/consumption    # Energy consumption
$ curl http://localhost:8080/sonnen-battery-api/api/status         # Battery status (incl. greed feed in, production and charge level)
```

## Build application by yourself

### Requirements

* Docker
* Go 1.21.x (if you want to build it without using docker builder)

### Build and run it

See [Makefile](Makefile)!

```sh 
$ make              # prints available make goals
```