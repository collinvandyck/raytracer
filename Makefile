.PHONY: run

run:
	go test .

fill:
	go run ./cmd/exp/fill
	open fill.ppm

cannon:
	go run ./cmd/exp/cannon
	open cannon.ppm
