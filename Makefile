.PHONY: test bench fill cannon lint

test: lint
	go test .

bench: lint
	go test . -bench .

fill:
	go run ./cmd/fill
	open fill.ppm

cannon:
	go run ./cmd/cannon
	open cannon.ppm

table:
	go run ./cmd/table

lint:
	golangci-lint run
