.PHONY: build

NAME=ianneub/httpecho

build:
	docker build -t $(NAME) .

run:
	docker run --rm -it -p 8080:8080 $(NAME)

push:
	docker push $(NAME)
