# Demo

## Stack
Golang version 1.19
Framework go gin
Database MySQL 8.0

## Getting started

### Environment Variables

Create a `.env` file in the root Peppermint directory and copy the default values from the `.env.template` file

```shell script
cp .env.template .env
```
Update the env value

### Init the database
Start MySQL by docker-compose
```shell script
docker-compose up mysql -d
```

Create database, tables
```shell script
bash ./init-mysql.sh
```

### Local Development Environment
### Start with code
```shell script
go run cmd/main.go api
```

#### Start with docker-compose

The dependencies for this project have all been provisioned for you through docker-compose. In order to run the application locally, run the following from the root demo directory

```shell script
docker-compose up  # start the application
```

```shell script
docker-compose down # stop the application
```

```shell script
docker-compose down --remove-orphans --volumes; docker-compose build; docker-compose up  # rebuild the application and rerun
```

### Start.sh
```shell script
     bash ./start.sh
```