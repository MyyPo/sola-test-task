#!/usr/bin/env bash

cd "${0%/*}"

CMD=$1

function main() {
	if [ "$CMD" == "local" ]; then
		docker-compose -f ../local.docker-compose.yaml build
		docker-compose -f ../local.docker-compose.yaml up -d

		until docker-compose -f ../local.docker-compose.yaml exec -T sola-database pg_isready; do
			sleep 1
		done
	else
		echo "Error: unknown run command specified: $CMD"
		exit 1
	fi
}

main
