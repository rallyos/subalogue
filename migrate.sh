#!/bin/bash

curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xvz
./migrate.linux-amd64 -source file://db/migrations -database $DATABASE_URL up

