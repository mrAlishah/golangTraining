get:
	cd server
	go get -t -u ./...

buildserver:
	echo "Build API"
	CGO_ENABLED=0 GOOS=linux go build -o ./server/kittenserver ./server/main.go

runserver: buildserver
	echo "Run API"
	docker-compose -p kittenserver -f ./docker-compose.yml up --build --remove-orphans #&& docker-compose logs -f
  
