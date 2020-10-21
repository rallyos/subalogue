## Getting started

### Config
You'll need an .env file with the following environment variables set:  
```
DATABASE_URL
AUTH0_CLIENT_ID
AUTH0_DOMAIN
AUTH0_CLIENT_SECRET
AUTH0_CALLBACK_URL
```

### Run the server
Build and run the docker containers, the most friendly way is through:  
`docker-compose up`

### Migrations

##### Create migration

Run the migrate container with entrypoint override:  
`docker-compose run --entrypoint="migrate create -ext sql -dir /migrations -seq {name}" migrate`

##### Apply the migrations
Most of the times, a single `docker-compose run migrate` is sufficient.

If a manual change is needed, you need to override the entrypoint command and provide a database url:  
`docker-compose run --entrypoint="migrate -database ${DATABASE_URL} -path /migrations up" migrate`

