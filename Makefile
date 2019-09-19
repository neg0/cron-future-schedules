.PHONY: build     # Builds docker image for deployment
build:
	go build -mod=vendor -o ./bin/cron_parser cmd/cron_parser.go

.PHONY: test      # tests
test:
	go test ./...

.PHONY: run        # creates and runs the container from built image
run:
	./bin/cron_parser --current-time=16:10

.PHONY: lint       # Linter
lint:
	golangci-lint run \
        --enable=golint \
        --enable=gofmt \
        --enable=gosec  \
        --enable=misspell \
        --enable=maligned \
        --enable=interfacer \
        --enable=stylecheck  \
        --enable=unconvert \
        --enable=maligned \
        --enable=goconst \
        --enable=nakedret