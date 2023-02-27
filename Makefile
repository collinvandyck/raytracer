.PHONY: test bench fill cannon lint

test: lint
	go test .

bench: lint
	go test -run notests -bench . -memprofile mem.out -cpuprofile cpu.out

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
