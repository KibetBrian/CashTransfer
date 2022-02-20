#Create database container
cc:
	docker run --name fisa-database -p 5432:5432 -e POSTGRES_USER=briankibet -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres:14-alpine
createdb:
	docker exec -it fisa-database createdb --username=briankibet --owner=briankibet fisa
dropdb:
	docker exec -it fisa-database dropdb --username=briankibet fisa
#Migrate database up 
migrateup:
	migrate -path ./database/migrations  -database "postgres://briankibet:${POSTGRES_PASSWORD}@localhost:5432/fisa?sslmode=disable" -verbose up
#Migrate database down
migratedown:
	migrate -path ./database/migrations  -database "postgres://briankibet:${POSTGRES_PASSWORD}@localhost:5432/fisa?sslmode=disable" -verbose down
#List running containers
rc:
	docker ps 
#List all containers
lc:
	docker ps -a
#Delete container
dc:
	docker rm fisa-database 
#Stop container
sc:
	docker stop fisa-database
#PSQL shell
ps:
	docker exec -it fisa-database psql -U briankibet