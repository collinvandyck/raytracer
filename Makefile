.PHONY: default
default: test lint

.PHONY: test
test: 
	go test .

.PHONY: lint
lint:
	golangci-lint run

build:
	mkdir -p bin
	go build -o bin/cannon ./cmd/cannon
	go build -o bin/sphere ./cmd/sphere

.PHONY: bench
bench:
	go test -run donotrun -bench . -memprofile mem.out -cpuprofile cpu.out

.PHONY: fill
fill:
	go run ./cmd/fill
	open fill.ppm

.PHONY: cannon
cannon:
	go run ./cmd/cannon
	open cannon.ppm

.PHONY: clockface
clockface:
	go run ./cmd/clockface
	open clockface.ppm

.PHONY: sphere
sphere:
	go run ./cmd/sphere
	open sphere.ppm

.PHONY: shade
shade:
	go run ./cmd/shade
	open shade.png

.PHONY: table
table:
	go run ./cmd/table

