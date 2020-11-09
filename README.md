<p align="center">
  <img src="logo.png" width="400" height="100"/>
</p>

---  

<p align="center">
  <img src="https://img.shields.io/badge/version-0.1.0-brightgreen" align="center"/></br>
  The API server side of https://subalogue.shifting-photons.dev</br>
  <strong>A nice welcoming first sentence hoes here.</strong>
</p>

## Features

## How To Use

## How To Contribute


### Config
You'll need an .env.{test,development} files with the following environment variables set:  
```
DATABASE_URL
SESSION_KEY
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

