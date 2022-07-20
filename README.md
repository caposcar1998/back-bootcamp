# Bootcamp exercise - Backend application
This repository stores backend application source code for the bootcamp final exercise.

## database
Package for database connection implementations.

## migrations
Database migrations folder.
Project use golang-migrate library.

To run migrations:
> make migrateup dbuser= dbpass= dbserver= dbport= dbname=
> make migratedown dbuser= dbpass= dbserver= dbport= dbname=

To create a new migration:
> migrate create -ext sql -dir migrations -seq {identifier}

## models
Package for business models.

## repositories
Package for repository implementations.

## services
Package for service implementations, depends on repositories.
