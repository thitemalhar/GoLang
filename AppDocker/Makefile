#.PHONY: docker-build docker-run
ID=
docker-build:
	docker build -t my-go-app .

docker-run:
	docker run -p 8080:8081 -d my-go-app

docker-ps:
	ID=`docker ps --format "{{.ID}}"`
	docker stop $(ID)
