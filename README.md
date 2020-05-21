# journal-app

[![Build Status](https://travis-ci.org/hafizbadrie/journal-app.svg?branch=master)](https://travis-ci.org/hafizbadrie/journal-app)

My training ground with Golang.

# Notes

## Initial setup:
1. Create the code base under `$GOPATH`.
2. Need to ensure I have `dep` tool.
3. Init go mod, by running `go mod init`.
4. Then, run `go mod vendor` to enable vendoring in golang.
5. Run `dep init` and `dep ensure`, but I need to remove my ssh passphrase first.
6. Build the code `go build main.go`.
7. Run it `./main`.

## For database:
1. Database creation should be handled by ansible script. Database name will exist in ansible variable. That variable is used to create database and an environment variable for the app to connect to DB.
2. I use github.com/golang-migrate/migrate. Install it in mac with `brew install golang-migrate`.
3. Run `create -ext sql -dir db/migrations -seq create_journals_table` to create table.
4. Run `migrate -database postgres://hafizbadrielubis@localhost:5432/journal_app_development?sslmode=disable -path db/migrations up` to apply the schema.
