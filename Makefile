run:
	 ENV=development go run main.go

test:
	 ENV=test go test -v ./src/...
