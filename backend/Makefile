BUILD_DIR=build
APP_NAME=server
IMAGE_MAME=decodethtx
APP_VERSION=v1
APP_PORT=3000

BUILD_FLAGS := -ldflags '-w -s -extldflags "-static"' -a

.PHONY: clean
clean:
	if [ -d $(BUILD_DIR) ] ; then rm -rf $(BUILD_DIR) ; fi


.PHONY: build
build: clean
	mkdir -p build/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    $(BUILD_FLAGS) -o $(BUILD_DIR)/ ./...


.PHONY: docker-build
docker-build:
	DOCKER_BUILDKIT=1 docker build -t $(IMAGE_MAME):$(APP_VERSION) --build-arg APP_NAME=$(APP_NAME) .


.PHONY: docker-run
docker-run:
	docker run -dp $(APP_PORT):$(APP_PORT) --name $(IMAGE_MAME)-$(APP_NAME) -e APP_PORT=$(APP_PORT) --restart unless-stopped $(IMAGE_MAME):$(APP_VERSION)

.PHONY: docker-stop
docker-stop:
	docker stop $(IMAGE_MAME)-$(APP_NAME) && docker rm $(IMAGE_MAME)-$(APP_NAME)

.PHONY: docker-clean
docker-clean:docker-stop
	docker rmi $(IMAGE_MAME):$(APP_VERSION)
