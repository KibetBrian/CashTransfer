#Create database container
cpc:
	docker run --name fisa-database --network fisa-infra -p 5432:5432 -e POSTGRES_USER=briankibet -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres:14-alpine
#Create working database
createdb:
	docker exec -it fisa-database createdb --username=briankibet --owner=briankibet fisa

#Build app image
bai:
	docker build -t fisa-app:1.0 .

#Create app container
rac:
	docker run --name fisa-app --network fisa-infra -p 8080:8080 

dropdb:
	docker exec -it fisa-database dropdb --username=briankibet fisa

#Migrate database up 
migrateup:
	migrate -path ./database/migrations  -database "postgres://briankibet:${POSTGRES_PASSWORD}@localhost:5432/fisa?sslmode=disable" -verbose up

#Migrate database down
migratedown:
	migrate -path ./database/migrations  -database "postgres://briankibet:${POSTGRES_PASSWORD}@localhost:5432/fisa?sslmode=disable" -verbose down

#Run container
rc:
	docker run fisa-database

#List running containers
lrc:
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

#Test 
test:
	go test -v -cover ./... -count=1
#Run 
run:
	go run main.go 
#Spin docker containers

dcu:
	docker compose up

#Stop running containers
dcd:
	docker compose down

#Run redis
drr:
	docker run --name redis -p 6379:6379 -d redis