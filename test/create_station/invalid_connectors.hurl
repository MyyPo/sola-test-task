# Invalid 0 connectors
POST http://localhost:12499/stations
{
  "uid": "{{uuid3}}",
  "public": true,
  "title": "Station 1",
  "description": "Description of station 1",
  "address": "123 Main St",
  "coordinates": [1.234, 2.345],
  "connectors": []
}
HTTP 400

# Invalid 9 connectors
POST http://localhost:12499/stations
{
  "uid": "{{uuid3}}",
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
      "max_power_kw": 50.0
    },
    {
      "id": "connector_3",
      "type": "Type1",
      "max_power_kw": 50.0
    },
    {
      "id": "connector_4",
      "type": "Type2",
      "max_power_kw": 50.0
    },
    {
      "id": "connector_5",
      "type": "CCS",
      "max_power_kw": 50.0
    },
    {
      "id": "connector_6",
      "type": "CHAdeMO",
      "max_power_kw": 50.0
    },
    {
      "id": "connector_7",
      "type": "Type1",
      "max_power_kw": 50.0
    },
    {
      "id": "connector_8",
      "type": "Type2",
      "max_power_kw": 50.0
    },
    {
      "id": "connector_9",
      "type": "CCS",
      "max_power_kw": 50.0
    }
  ]
}
HTTP 400

# Invalid too much connector power
POST http://localhost:12499/stations
{
  "uid": "{{uuid3}}",
  "public": true,
  "title": "Station 1",
  "description": "Description of station 1",
  "address": "123 Main St",
  "coordinates": [1.234, 2.345],
  "connectors": [
    {
      "id": "connector_1",
      "type": "CCS",
      "max_power_kw": 9999999
    }
  ]
}
HTTP 400

# Invalid too little connector power
POST http://localhost:12499/stations
{
  "uid": "{{uuid3}}",
  "public": true,
  "title": "Station 1",
  "description": "Description of station 1",
  "address": "123 Main St",
  "coordinates": [1.234, 2.345],
  "connectors": [
    {
      "id": "connector_1",
      "type": "CCS",
      "max_power_kw": 0
    }
  ]
}
HTTP 400

# Invalid no connector power
POST http://localhost:12499/stations
{
  "uid": "{{uuid3}}",
  "public": true,
  "title": "Station 1",
  "description": "Description of station 1",
  "address": "123 Main St",
  "coordinates": [1.234, 2.345],
  "connectors": [
    {
      "id": "connector_1",
      "type": "CCS"
    },
    {
      "id": "connector_1",
      "type": "CCS",
      "max_power_kw": 25.11
    }
  ]
}
HTTP 400
