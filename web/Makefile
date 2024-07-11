DOCKER_TAG ?= latest

# docker-build:
# 	@make --no-print-directory codegen
# 	docker build -t pineapple217/Sortify:$(DOCKER_TAG) --build-arg GIT_COMMIT=$(shell git log -1 --format=%h) . 

# docker-push:
# 	docker push pineapple217/mb:$(DOCKER_TAG)

# docker-update:
# 	@make --no-print-directory docker-build
# 	@make --no-print-directory docker-push

codegen:
	go generate ./ent
	templ generate

build:
	@make --no-print-directory codegen
	go build -o ./tmp/main.exe ./cmd/main

start:
	@./tmp/main.exe

make run:
	@make --no-print-directory build
	@make --no-print-directory start