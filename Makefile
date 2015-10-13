.PHONY: build

NAME=ianneub/http-test-server
OUTPUT=httpecho

build:
	docker run -v "$(PWD):/src" centurylink/golang-builder
	docker build -t $(NAME) .
	rm $(OUTPUT)

run:
	docker run --rm -it -p 8080:8080 $(NAME)

push:
	docker push $(NAME)

