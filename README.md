<p align="center">
  <a href="https://subalogue.shifting-photons.dev"><img src="logo.png" width="400" height="100"/></a>
</p>

---  

<p align="center">
  A subscription manager aiming to offer a better visibility of your personal subscriptions.</br></br>
  <img src="https://img.shields.io/github/workflow/status/shiftingphotons/subalogue/Test"/></br>
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

#### Prerequisites
Subalogue uses [auth0](https://auth0.com/) for authentication, so at this point that is still a requirement to run the app locally. In the future this will be omitted to make things easy.

#### Config
Bring your own .env.development file to the project root folder with the following environment variables set:  
```
DATABASE_URL=URL_TO_YOUR_DB
SESSION_KEY=RANDOM_STRING
AUTH0_CLIENT_ID=YOUR_AUTH0_CLIENT_ID
AUTH0_DOMAIN=YOUR_AUTH0_DOMAIN
AUTH0_CLIENT_SECRET=YOUR_AUTH0_CLIENT_SECRET
AUTH0_CALLBACK_URL=YOUR_AUTH0_CALLBACK_URL
REDIRECT_APP_URL=URL_TO_THE_CLIENT_APP
```

#### Run the server
Build and run the docker containers, the most friendly way is through:
```
docker-compose up
```  
The web container runs [Reflex](https://github.com/cespare/reflex) on startup, so the binary will be rebuilt and started again on every file change.
#### Apply the migrations
A better way for this is still in TODO.  
```
docker-compose exec api bash -c "migrate -database ${DATABASE_URL} -path db/migrations up
```

#### Check out the client
The frontend is built with Vue. [Check out the app](https://github.com/shifting-photons/subalogue_client).

## How To Contribute
If you are interested, thank you.  
It's important to note that I'm not a Go expert and this is a pet project on which I still learn.  
With that said - if for some reason you've found this repo, this is still very much in active development, which I prefer to do alone until all intended features are done and polished.
