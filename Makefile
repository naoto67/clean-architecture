SHELL=/bin/zsh

run:
	 ENV=development go run -tags local ./cmd/server/main.go

.PHONY: test
test:
	ENV=test ROOT_PATH=$(ROOT_PATH) go test -v  -tags integration ./src/...
