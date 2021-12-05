SHELL := /bin/bash
.DEFAULT_GOAL := build

.EXPORT_ALL_VARIABLES:
SERVER_PORT ?= :8000
DB_USER ?= mohammad
DB_PASSWORD ?= 173946285
DB_HOST ?= localhost
DB_PORT ?= 3306
DB_NAME ?= mainapi
DB_ENGINE ?= SQL
SECRET_KEY ?= super secret key
PAGE_SIZE ?= 20

build:
	cd src; go build -o main.exe main-api-store-management
	./src/main.exe


