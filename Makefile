-include .env

.PHONY:run
run:
	go build .
	go run .