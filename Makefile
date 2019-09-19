.PHONY: build     # Builds docker image for deployment
build:
	go build -mod=vendor -o ./bin/cron_parser cmd/cron_parser.go

.PHONY: run        # creates and runs the container from built image
run:
	./bin/cron_parser --current-time=16:10