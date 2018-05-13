.PHONY: test build dist clean

default: test build dist

build:
	@mkdir -p build/
	GOOS=linux go build -o build/main

test:
	go test

dist:
	zip -j deployment.zip build/main

clean:
	-rm -rf build/
	-rm deployment.zip
