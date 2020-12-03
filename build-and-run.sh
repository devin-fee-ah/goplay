$(go env GOPATH)/bin/swag init
go generate ./ent
go build
./api
