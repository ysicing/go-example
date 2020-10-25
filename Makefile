BUILD_VERSION   ?= $(shell cat version.txt || echo "0.0.1")
BUILD_DATE      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD || echo "0.0.0")

help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

fmt:

	@echo gofmt -l
	@OUTPUT=`gofmt -l . 2>&1`; \
	if [ "$$OUTPUT" ]; then \
		echo "gofmt must be run on the following files:"; \
        echo "$$OUTPUT"; \
        exit 1; \
    fi

#lint:
#
#	@echo golint ./...
#	@OUTPUT=`command -v golint >/dev/null 2>&1 && golint ./... 2>&1`; \
#	if [ "$$OUTPUT" ]; then \
#		echo "golint errors:"; \
#		echo "$$OUTPUT"; \
#		exit 1; \
#	fi

default: fmt ## fmt code

build: clean ## 构建二进制
	@echo "build bin ${BUILD_VERSION} ${BUILD_DATE} ${COMMIT_SHA1}"
	#@bash hack/docker/build.sh ${version} ${tagversion} ${commit_sha1}
	# go get github.com/mitchellh/gox
	@CGO_ENABLED=1 GOARCH=amd64 go build -o dist/go-example \
    	-ldflags   "-X 'app/cmd/command.commit=${COMMIT_SHA1}' \
                    -X 'app/cmd/command.date=${BUILD_DATE}' \
                    -X 'app/cmd/command.release=${BUILD_VERSION}'"

docker: ## 构建镜像
	docker build -t goexample .

clean: ## clean
	rm -rf dist/*

install: clean ## install
	go install \
		-ldflags "-X 'app/cmd/command.commit=${COMMIT_SHA1}' \
                                       -X 'app/cmd/command.date=${BUILD_DATE}' \
                                       -X 'app/cmd/command.release=${BUILD_VERSION}'"


.PHONY : build release clean install

.EXPORT_ALL_VARIABLES:

GO111MODULE = on
GOPROXY = https://goproxy.cn
GOSUMDB = sum.golang.google.cn