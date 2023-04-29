cucumber:
	docker-compose up -d
	#cd features && godog ./
	cd features && go test -v ./...
	docker-compose stop

run:
	docker-compose up -d
	go run main.go
	docker-compose stop