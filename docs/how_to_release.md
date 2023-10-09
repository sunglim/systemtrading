# How to release

To release the app, there're 4 steps - a) VERSION udpate, b) git tag, c) docker release, d) go package release

## VERSION update

The docker release stage is depends on VERSION. The version in VERSION is used to docker version. See Makefile for details.

## git tag

It is intended to use same version in VERSION file.

## Docker release

* `docker login`
* Update version in [VERSION](https://github.com/sunglim/systemtrading/blob/main/VERSION)
* run `make push` to build a docker image and push to docker hub

## Go package release

is still an optional stage.
