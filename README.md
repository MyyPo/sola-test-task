# Charging stations data validation web-service

A simple **Go** web service built with **Gin** and **Jet** as a test task for Sola.

### Features

- Implements the listed validation of incoming JSON input for stations and attached connectors.
- Provides internationalization of errors based on incoming `Accept-Language` http headers using gin's built-in go-validator package. For simplicity, currently implemented for Spanish alongside English.

### Requirements and usage

To run locally with default configuration using **Docker** with **docker-compose**,
have to run either `just run local` or `docker-compose -f ../local.docker-compose.yaml up -d` from the root directory of the project.

It will deploy server on port **12499** and a postgres database on port **5432**.
Migrtions will be applied automatically.

The only endpoint is **POST** /stations

Example payload:

```
{
  "uid": "{{uuid1}}",
  "public": true,
  "title": "Station 1",
  "description": "Description of station 1",
  "address": "123 Main St",
  "coordinates": [1.234, 2.345],
  "connectors": [
    {
      "id": "connector_1",
      "type": "CCS",
      "max_power_kw": 50.0
    },
    {
      "id": "connector_2",
      "type": "CHAdeMO",
      "max_power_kw": 30.0
    }
  ]
}
```
