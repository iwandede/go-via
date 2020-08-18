# Go Native Framework
> A starter project with Golang, GORM and Postgresql


Golang framework boilerplate with Postgresql, GORM resource. Supports multiple configuration environments.


This project use a [GORM](http://gorm.io/).

Setup Database in file : 
```
config/production.yaml for environtment production
config/development.yaml for environtment development
```


### Application structure

```
.
├── Makefile
├── Procfile
├── README.md
├── main.go
├── config
│   ├── config.go
│   ├── development.yaml
│   ├── production.yaml
├── controllers
│   └── user.go
├── database
│   └── database.go
├── middleware
│   └── handlers.go
├── models
│   └── user.go
└── server
    ├── router.go
    └── server.go
```

## Installation

```sh
go get
```

## Run Application

#### Development Environment

```
go run main.go -config config/development.yaml  -env development
```

```
go build ./go-via && -config config/development.yaml  -env development
```

#### Production Environment

```
go run main.go
```

```
go build ./go-via
```

## Usage example

`curl http://localhost:8888/ping`


## Release History

* 0.0.1
    * Configuration by environment, Auth and Log middlewares, User entity.

