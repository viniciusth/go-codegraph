
artifacts:
	@echo "Creating artifacts directory..."
	mkdir -p artifacts

build-graph: artifacts
	@echo "Building graph..."
	cd api && go run cmd/main.go

build-go: artifacts
	@echo "Building api..."
	cd api && go build -o ../artifacts/main-go cmd/main.go

.PHONY: dev-gui
dev-gui:
	@echo "Running GUI..."
	cd gui && npm run dev -- --open
