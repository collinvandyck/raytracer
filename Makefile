.PHONY: run fill cannon lint

run: lint
	go test .

fill:
	go run ./cmd/exp/fill
	open fill.ppm

cannon:
	go run ./cmd/exp/cannon
	open cannon.ppm

lint:
	golangci-lint run
