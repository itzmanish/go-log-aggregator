test:
	CI=true go test -v ./...
build:
	go build -o log-aggregator .
