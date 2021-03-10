# Postgres Bookmark Service

## Postgres Docker

`docker run --name bookmark-service-db -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres`

SQL migrations are located under `db/migrations` 

## Run service with docker-compose

Start `docker-compose up -d`

Rebuild `docker-compose build`

Stop `docker-compose down -v`

## Manual test

GET bookmarks `curl localhost:8080/bookmark`

Create bookmark 

```
curl --location --request POST 'localhost:8080/bookmark' \
--header 'Content-Type: application/json' \
--data-raw '{
"category":"general",
"name":"YouTube",
"uri":"https://youtube.com"
}'
```

## Check docker-compose logs
All `docker-compose logs`

Individual service


`docker-compose logs bookmark-service`

`docker-compose logs postgres`
