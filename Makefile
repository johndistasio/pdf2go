.PHONY: test build dist clean

default: test build dist

build:
	GOOS=linux go build -o main

test:
	go test

dist:
	zip deployment.zip main

clean:
	rm deployment.zip
