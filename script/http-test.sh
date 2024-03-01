#!/usr/bin/env bash

cd "${0%/*}"

CMD=$1

function main() {
	if [ "$CMD" == "local" ]; then
		hurl --test \
			--variable uuid1="$(uuidgen)" \
			--variable uuid2="$(uuidgen)" \
			--variable uuid3="$(uuidgen)" \
			--variable uuid4="$(uuidgen)" \
			../test/create_station/valid.hurl \
			../test/create_station/invalid_absent.hurl \
			../test/create_station/invalid_coordinates.hurl \
			../test/create_station/invalid_connectors.hurl \
			../test/create_station/invalid_duplicate.hurl \
			../test/create_station/language.hurl
	else
		echo "Error: unknown http-test command specified: $CMD"
		exit 1
	fi
}

main
