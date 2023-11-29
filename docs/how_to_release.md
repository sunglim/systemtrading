# How to release

To release the app, there're 4 steps - a) [VERSION](https://github.com/sunglim/systemtrading/blob/main/VERSION) update, b) git tag, c) docker release, d) go package release.

Currently the step b,c,d is done by github action.

## VERSION update

[VERSION] keeps the version number to be released. The version in VERSION is used to docker version. See Makefile for details.

## git tag

Current policy is that using the same version in VERSION file. For example, if the version in VERSION file is '0.1.6', the tag is 'v0.1.6'.

## Docker release

* `docker login`
* Update version in [VERSION](https://github.com/sunglim/systemtrading/blob/main/VERSION)
* run `make push` to build a docker image and push to docker hub

## Go package release

is still an optional stage.
