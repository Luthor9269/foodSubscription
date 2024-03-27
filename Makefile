.PHONY: start
start: stop server

.PHONY: stop
stop:
	-docker stop foodsubscription-db-1
	-docker rm -v foodsubscription-db-1

.PHONY: server
server: docker
	go run main.go

.PHONY: docker
docker:
	-docker compose up -d