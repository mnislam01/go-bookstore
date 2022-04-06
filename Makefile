startDB:
	docker-compose -f db.yml up
stopDB:
	docker-compose -f db.yml stop
build:
	go build -o NewApp src/main/main.go
start:
	./NewApp
clear:
	rm NewApp
