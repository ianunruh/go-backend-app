.PHONY: generate
generate:
	go generate ./...

.PHONY: prettier
prettier:
	prettier --write .

.PHONY: server
server:
	OTEL_SERVICE_NAME=go-backend-app-server \
	APP_METRICS_LISTEN_ADDR=localhost:9081 \
	APP_DEBUG_LISTEN_ADDR=localhost:9082 \
	go run . serve
