go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate -source file://db/migrations -database ${DATABASE_URL} up
go build .
