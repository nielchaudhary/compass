.PHONY: run-core run-monitor core-dev monitor-dev build-core build-monitor clean-artifacts

run-core:
	@echo "Running core server"
	go run ./cmd/core/main.go

run-monitor:
	@echo "Running monitor server"
	go run ./cmd/monitor/main.go

core-dev:
	@echo "Running compass-core in dev mode with hot reload"
	air -c .air-core.toml

monitor-dev:
	@echo "Running compass-monitor in dev mode with hot reload"
	air -c .air-monitor.toml


build-core:
	@echo "Building Compass: Core SERVER"
	cd cmd/core && go build -o ../../bin/core


build-monitor:
	@echo "Building Compass: Monitor SERVER"
	cd cmd/monitor && go build -o ../../bin/monitor

clean-artifacts:
	@echo "Cleaning Build Artifacts"
	rm -rf bin
	rm -rf tmp
	@echo "Build Artifacts Cleaned up"
