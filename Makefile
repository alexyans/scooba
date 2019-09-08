all: install build
.PHONY: all

install:
	@echo "Building Docker image..."
	@docker build --rm -t scooba:latest .
	@echo "Done."
.PHONY: install

build:
	@echo "Building binary..."
	@docker run --rm -it -v `pwd`:/go/src/github.com/alexyans/scooba scooba:latest /bin/bash scripts/build.sh
	@echo "Done."
.PHONY: build

shell:
	@echo "Starting shell session..."
	@docker run --rm -it -v `pwd`:/go/src/github.com/alexyans/scooba/ scooba:latest bash
.PHONY: shell

clean:
	@echo "Deleting binary and Docker image..."
	@rm -rf scooba lib
	@docker image rm scooba:latest
	@echo "Done."
.PHONY: clean
