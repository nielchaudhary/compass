.PHONY: compass-core compass-monitor core-dev monitor-dev build-core build-monitor

compass-core:
	go run ./cmd/core/main.go

compass-monitor:
	go run ./cmd/monitor/main.go

core-dev:
	@echo "Running compass-core in dev mode"
	air -c .air-core.toml

monitor-dev:
	@echo "Running compass-monitor in dev mode"
	air -c .air-monitor.toml


build-core:
	@echo "Building Compass: Core SERVER"
	cd cmd/core && go build -o ../../bin/core


build-monitor:
	@echo "Building Compass: Monitor SERVER"
	cd cmd/monitor && go build -o ../../bin/monitor

