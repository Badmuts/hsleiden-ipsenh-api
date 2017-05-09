REPO=badmuts

# Name of the image
IMAGE=hsleiden-ipsenh-api

# Current branch-commit (example: master-ab01c1z)
CURRENT=`echo $$TRAVIS_BRANCH | cut -d'/' -f 2-`-$$(git rev-parse HEAD | cut -c1-7)

# Colors
GREEN=\033[0;32m
NC=\033[0m

.PHONY: coverage

run: start
suite: install lint test coverage
package: install compile build

# Install packages
install:
	glide install

# Run linters, simple code quality check
lint:
	golint $$(go list ./... | grep -v /vendor/)

# go tool vet $$(go list ./... | grep -v /vendor/)

# Run tests
# Coverage is disabled because of this: https://lk4d4.darth.io/posts/multicover/
# mkdir -p coverage
# go test -v -coverprofile=coverage/c.out $$(go list ./... | grep -v /vendor/)
test:
	go test -v $$(go list ./... | grep -v /vendor/)

# Create coverage report
coverage:
	go tool cover -html=coverage/c.out -o coverage/coverage.html

# Jenkins step to run complete pipeline
ci-tests:
	echo "$(GREEN)--- BUILDING TEST IMAGE ---$(NC)"
	docker build -t $(IMAGE):test -f operations/docker/Dockerfile.test .
	echo "$(GREEN)--- RUNNING TEST SUITE ---$(NC)"
	docker run -v "$$(pwd):/go/src/github.com/$(REPO)/$(IMAGE)" --rm $(IMAGE):test bash -c 'make suite'
	echo "$(GREEN)--- COMPILE BINARY ---$(NC)"
	docker run -v "$$(pwd):/go/src/github.com/$(REPO)/$(IMAGE)" --rm $(IMAGE):test bash -c 'make compile'

# Ci step to run complete pipeline
ci: ci-tests build push cleanup

# Create binary
compile:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# Create docker image with tag badmuts/$(IMAGE):branch-sha
build:
	echo "$(GREEN)--- BUILDING DOCKER IMAGE ---$(NC)"
	docker build -t $(REPO)/$(IMAGE):$(CURRENT) -f operations/docker/Dockerfile .

# Push image to the hub, this also build the image
push: build
	echo "$(GREEN)--- PUSHING IMAGE TO HUB ---$(NC)"
	docker push $(REPO)/$(IMAGE):$(CURRENT)

# Cleanup step to remove test image and build image
cleanup:
	docker rmi $(IMAGE):test
	docker rmi $(REPO)/$(IMAGE):$(CURRENT)

# Run development via docker-compose. This autoreloads/compiles on change etc.
start:
	docker-compose up