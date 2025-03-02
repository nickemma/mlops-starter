# Build automation for the project
.PHONY: all train build-go deploy clean

all: train build-go

train:
	@echo "Training model..."
	cd python/training && python train.py

build-go:
	@echo "Building Go binary..."
	cd go/api && go build -o ../../bin/mlops-api

docker-build:
	docker build -t your-registry/mlops-go-api -f Dockerfile.go .
	docker build -t your-registry/mlops-py -f Dockerfile.python .

deploy:
	kubectl apply -f deployments/

clean:
	rm -rf bin/

enable-ci:
    @mkdir -p .github/workflows
    @cp .github/workflow-templates/* .github/workflows/
    @echo "CI/CD workflows enabled"

disable-ci:
    @rm -rf .github/workflows
    @echo "CI/CD workflows disabled"