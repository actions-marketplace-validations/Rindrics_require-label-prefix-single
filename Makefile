LOG_LEVEL ?= info

.PHONY: test
test:
	@LOG_LEVEL=${LOG_LEVEL} go test -v ./...
