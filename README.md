# JULO Mini Wallet

## Prerequisites

The following must be installed on your local in order to run the project properly

- Go programming language runtime (minimum version `1.20`)
- Docker (minimum version `24.0.2`)
- Docker Compose (minimum version `2.18.1`)
- Make, automated execution using Makefile (minimum version `3.81`)

## Instructions

The following are the instructions to run the app in your local

### Build the app

First of all, we need to build our app into an executable/compiled file

```bash
make build
```

### Run the database

The app has database dependency to store the data, spin up the database instance by executing the following command

```bash
docker-compose up -d db
```

### Migrate the database

We then need to migrate the database tables, but we have to ensure that the database is ready to accept connections by using the following command

```bash
docker-compose ps
```

Make sure the database is in `healthy` status. Once this done, do the migrations using the below command

```bash
# Do migration up (prepare the tables)
make migrate-up

# If something unwanted happens or you just want to clean up the database, then just migrate it down
make migrate-down
```

### Run the app

Finally, you are good to run the app

```bash
docker-compose up -d app
```

You can import and use the following Postman collection to use the app

<https://documenter.getpostman.com/view/8411283/SVfMSqA3?version=latest#a1d69cb0-1de4-4cba-af84-388748c69431>
