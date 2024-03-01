#!/usr/bin/env bash

cd "${0%/*}"

CMD=$1

cd ../migration

function main() {
	if [ "$CMD" == "local" ]; then
		export GOOSE_DRIVER=postgres
		export GOOSE_DBSTRING="postgresql://postgres:postgres-sola@localhost:5432/sola?sslmode=disable"
		goose up
	else
		echo "Error: unknown apply-migration command specified: $CMD"
		exit 1
	fi
}

main
