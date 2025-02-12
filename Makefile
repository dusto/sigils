GIT_COMMIT_SHA:=$(shell git rev-parse --short HEAD)
GIT_BUILD:=$(shell git describe --dirty --always)
GIT_REF=$(shell git rev-parse --abbrev-ref HEAD)


build:
	CGO_ENABLED=1 GOOS=linux go build \
   -ldflags " \
   -X 'github.com/dusto/sigils/internal/version.Version=$(shell cat version)' \
   -X 'github.com/dusto/sigils/internal/version.GIT_COMMIT=${GIT_COMMIT_SHA}' \
   -X 'github.com/dusto/sigils/internal/version.GIT_BRANCH=${GIT_REF}' \
   -X 'github.com/dusto/sigils/internal/version.GIT_BUILD=${GIT_BUILD}'\
   " 

build-docker:
	docker build \
		--build-arg GIT_COMMIT_SHA=$(GIT_COMMIT_SHA) \
		--build-arg GIT_BUILD=$(GIT_BUILD) \
		--build-arg GIT_REF=$(GIT_REF) \
		--build-arg VERSION=$(shell cat version) \
		-t dusto/sigilslocal .
