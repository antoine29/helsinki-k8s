# ToDo DB migrations

This is a liquibase docker-based application meant to set up the database schema and tables for the ToDo project.

## Building the image 
You'll need a .env file to store the DB credentials with the following fields:

```shell
PG_HOST=172.17.0.1
PG_PORT=5432
PG_DBNAME=postgres
PG_USERNAME=postgres
PG_PASSWORD=postgres
CHANGE_LOG=/db1/db1.changelog-master.xml
```
Then build and run the image with:

```shell
$ docker build . -t be-migrations
$ docker run -it --rm --env-file ./.env be-migrations
```

