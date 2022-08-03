# Money Count Project

[![Go Report Card](https://goreportcard.com/badge/github.com/i0tool5/money-count)](https://goreportcard.com/report/github.com/i0tool5/money-count)

Simple RESTFull application, created to track expenses and provide information on them

It hasn't done yet, but it will eventually

## DB

Currently used database is Postgresql

[migrate](https://github.com/golang-migrate/migrate) tool is used to apply migrations to the database

## ORM

gorm is used with postgres driver

## Quick Start 

```shell
make docker.run 

# Flow:
#     - Build and run postgres Docker container
#     - Create all databases
#     - Apply migrations (github.com/golang-migrate/migrate)

```
