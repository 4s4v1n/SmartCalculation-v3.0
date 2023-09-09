.PHONY: all docker-run clean

all: docker-run

docker-run:
	docker-compose up --build

clean:
	docker system prune
