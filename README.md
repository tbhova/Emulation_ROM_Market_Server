# Emulation_ROM_Market_Server
Server Program For Emulation Market

# Tools Needed

[GoLang](https://golang.org/)

[GogLand](https://www.jetbrains.com/go/)

Also get the GogLand Protoc plugin (install within the IDE)

[Flyway Command Line](https://flywaydb.org)

[Postgres](https://www.postgresql.org/download/)

[Lastest protoc](https://github.com/google/protobuf/releases)

Protoc needs to be placed somewhere on the path

# Go Packages

Run the following commands:

```
go get github.com/lib/pq
go get google.golang.org/grpc
go get -u github.com/golang/lint/golint
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/proto
```

So We can use protoc (protobuf compiler):

```
export PATH=$PATH:$GOPATH/bin
````
