IMAGE_TAG ?= suapapa/okt-api:latest
# BUILD_ARCHS ?= linux/amd64,linux/arm64
BUILD_ARCHS ?= linux/arm64
BUILD_FLAGS ?= --no-cache
DOCKERFILE ?= Dockerfile
CONTEXT ?= .

build_image:
	docker buildx build $(BUILD_FLAGS) --platform $(BUILD_ARCHS) -t $(IMAGE_TAG) -f $(DOCKERFILE) $(CONTEXT)

push_image: build_image
	docker push $(IMAGE_TAG)

.PHONY: build_image push_image
