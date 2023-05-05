postgres:
	docker run -d -p 5432:5432 --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:14.2-alpine

dbshell:
	docker exec -it postgres psql -U root -d newDB

createdb:
	docker exec -it postgres createdb --username=root --owner=root newDB

dropdb:
	docker exec -it postgres dropdb --username=root newDB
	
migrateup:
	go run main.go -migrateup

migratedown:
	go run main.go -migratedown
run:
	go run main.go -run
dbseed:
	go run main.go -dbseed

build:
	go build -o main main.go

.PHONY:
	postgres
	dbshell
	createdb
	dropdb
	migrateup
	migratedown
	run
	build