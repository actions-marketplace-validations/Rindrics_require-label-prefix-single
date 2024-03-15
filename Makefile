LOG_LEVEL ?= info

.PHONY: test
# set log level to see debug logs
test:
	@LOG_LEVEL=${LOG_LEVEL} go test -v ./...
