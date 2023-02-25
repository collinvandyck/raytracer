.PHONY: run

run:
	go test .

cannon:
	go run ./cmd/exp/cannon
	open cannon.ppm
