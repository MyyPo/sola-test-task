#!/usr/bin/env bash

cd "${0%/*}"

CMD=$1

function main() {
	if [ "$CMD" == "local" ]; then
		rm -rf ../.gen/jet
		jet -dsn=postgresql://postgres:postgres-sola@localhost:5432/sola?sslmode=disable -path=../.gen/jet
	else
		echo "Error: unknown gen-models command specified: $CMD"
		exit 1
	fi
}

main
