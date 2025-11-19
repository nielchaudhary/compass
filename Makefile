.PHONY: compass-core compass-monitor

compass-core:
	go run ./cmd/core/main.go

compass-monitor:
	go run ./cmd/monitor/main.go