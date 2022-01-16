SHELL := /bin/bash
.DEFAULT_GOAL := buildb

.EXPORT_ALL_VARIABLES:
SERVER_PORT ?= 8000
DB_USER ?= fibbinbbrzdnns
DB_PASSWORD ?= f59c8fd499c54e2edab8ba8191e99344f282a8dced74621d97e11bd0a6168458
DB_HOST ?= ec2-34-242-89-204.eu-west-1.compute.amazonaws.com
DB_PORT ?= 5432
DB_NAME ?= d8th8fb7pjerr2
DB_ENGINE ?= PostgreSQL
SECRET_KEY ?= super secret key
PAGE_SIZE ?= 20

build:
	go build -o main.exe main-api-store-management
	./main.exe


