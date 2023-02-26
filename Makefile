.PHONY: run fill cannon lint

run: lint
	go test .

fill:
	go run ./cmd/fill
	open fill.ppm

cannon:
	go run ./cmd/cannon
	open cannon.ppm

lint:
	golangci-lint run
