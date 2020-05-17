# journal-app
My training ground with Golang.

# Activities
What I went through:

1. Create the code base under `$GOPATH`
2. Need to ensure I have `dep` tool
3. Init go mod, by running `go mod init`
4. Then, run `go mod vendor` to enable vendoring in golang
5. Run `dep init` and `dep ensure`, but I need to remove my ssh passphrase first
6. Build the code `go build main.go`
7. Run it `./main` 

For database
1. I use github.com/golang-migrate/migrate. Install it in mac with `brew install golang-migrate`
2. Run `create -ext sql -dir db/migrations -seq create_journals_table` to create table
3. Run `migrate -database postgres://hafizbadrielubis@localhost:5432/journal_app_development?sslmode=disable -path db/migrations up` to apply the schema
