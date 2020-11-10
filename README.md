<p align="center">
  <img src="logo.png" width="400" height="100"/>
</p>

---  

<p align="center">
  <strong>Manage your subscriptions with ease.</strong></br>
  The API server side and main repo behind https://subalogue.shifting-photons.dev</br>
  <img src="https://img.shields.io/badge/version-0.1.0-brightgreen" align="center"/></br>
</p>

## Features
- Keep a list your subscriptions
- ... more to come

## Roadmap
- Set the period on which the subscription is paid
- Set subscription category
- Set subscription tags
- Filter by payment period, category or tags
- Filter by keyword
  
  
## How To Use
### Run It Locally
If you are interested in using Subalogue on your machine instead of the [hosted version](https://subalogue.shifting-photons.dev), there is a straighforward, although not polished way.  
Subalogue uses [auth0](https://auth0.com/) for authentication, so at this point that is still a requirement to run the app locally. At future point this will be omitted to make things easy.

#### Config
Bring your own .env.development file to the project root folder with the following environment variables set:  
```
DATABASE_URL
SESSION_KEY
AUTH0_CLIENT_ID
AUTH0_DOMAIN
AUTH0_CLIENT_SECRET
AUTH0_CALLBACK_URL
```

#### Run the server
Build and run the docker containers, the most friendly way is through `docker-compose up`  
The web container runs [Reflex](https://github.com/cespare/reflex) on startup, so the binary will be rebuilt and started again on every file change.

#### Check out the client
The frontend is built with Vue, [check it out here](https://github.com/shifting-photons/subalogue_client).

### Development
#### Database changes
[Migrate](https://github.com/golang-migrate/migrate) Is used to create migrations for the DB.  

#### Create a migration
To create new migration, run the migrate container with entrypoint override:  
`docker-compose run --entrypoint="migrate create -ext sql -dir /migrations -seq {name}" migrate`

#### Apply the migrations
Most of the times, a single `docker-compose run migrate` is sufficient.  
There are some specifics that are best explained in their [README](https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md)

## How To Contribute
If you are interested, thank you.  
It's important to note that I'm not a Go expert and this is a pet project on which I still learn.  
With that said - if for some reason you've found this repo, this is still very much in active development, which I prefer to do alone until all intended features are done and polished.
