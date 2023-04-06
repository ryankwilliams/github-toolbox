CONTAINER_ENGINE ?= podman
OUTPUT_DIR = out

build:
	mkdir -p "$(OUTPUT_DIR)"
	go build -o "$(OUTPUT_DIR)/github-toolbox"

build-image:
	$(CONTAINER_ENGINE) build -t ocm-toolbox:main .

format:
	gofmt -s -w .
