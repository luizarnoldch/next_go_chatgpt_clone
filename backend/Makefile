PSQL-LOCAL := templates/docker/psql.yml

init:
	go mod init main
update:
	go mod tidy
dev:
	go run main.go
test:
	go clean -testcache
	go test ./... -v
psql-cmd:
	docker exec -it postgres-local psql -U postgres -d chatgpt_clone
psql-up:
	docker-compose -f $(PSQL-LOCAL) up -d
psql-stop:
	docker-compose -f $(PSQL-LOCAL) stop
psql-start:
	docker-compose -f $(PSQL-LOCAL) start
psql-destroy:
	docker-compose -f $(PSQL-LOCAL) down -v
	# docker rmi $(shell docker images amazon/dynamodb-local -q)