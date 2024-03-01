# Invalid coordinates latitude
POST http://localhost:12499/stations
{
  "uid": "{{uuid3}}",
  "public": true,
  "title": "Station 1",
  "description": "Description of station 1",
  "address": "123 Main St",
  "coordinates": [90.111, 2.345],
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
HTTP 400

# Invalid coordinates longitude
POST http://localhost:12499/stations
{
  "uid": "{{uuid3}}",
  "public": true,
  "title": "Station 1",
  "description": "Description of station 1",
  "address": "123 Main St",
  "coordinates": [1.234, -181.333],
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
HTTP 400

# Invalid empty object
POST http://localhost:12499/stations
{
}
HTTP 400
