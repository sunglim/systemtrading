REGISTRY ?= limasdf
TAG_PREFIX = v
VERSION = $(shell cat VERSION)
TAG ?= $(TAG_PREFIX)$(VERSION)
ARCH ?= $(shell go env GOARCH)
OS ?= $(shell uname -s | tr A-Z a-z)
ALL_ARCH = amd64 arm arm64
GO_VERSION = 1.21.1
IMAGE = $(REGISTRY)/systemtrading
DOCKER_CLI ?= docker

export DOCKER_CLI_EXPERIMENTAL=enabled

build-local:
	GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=0 go build -o systemtrading

container: container-$(ARCH)

container-%:
	${DOCKER_CLI} build --pull -t $(IMAGE)-$*:$(TAG) --build-arg GOARCH=$* --build-arg GOVERSION=$(GO_VERSION) .

sub-container-%:
	$(MAKE) --no-print-directory ARCH=$* container

all-container: $(addprefix sub-container-,$(ALL_ARCH))

push: $(addprefix sub-push-,$(ALL_ARCH)) push-multi-arch;

sub-push-%: container-% do-push-% ;

do-push-%:
	${DOCKER_CLI} push $(IMAGE)-$*:$(TAG)

push-multi-arch:
	${DOCKER_CLI} manifest create --amend $(IMAGE):$(TAG) $(shell echo $(ALL_ARCH) | sed -e "s~[^ ]*~$(IMAGE)\-&:$(TAG)~g")
	@for arch in $(ALL_ARCH); do ${DOCKER_CLI} manifest annotate --arch $${arch} $(IMAGE):$(TAG) $(IMAGE)-$${arch}:$(TAG); done
	${DOCKER_CLI} manifest push --purge $(IMAGE):$(TAG)
