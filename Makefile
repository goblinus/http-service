.PHONY: clean
clean:
	rm -rf bin/*

.PHONY: build
build:
	go build -o bin/app cmd/*

.PHONY: start
start:
	bin/app

.PHONY: run
run: clean build start