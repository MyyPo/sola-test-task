# Invalid absent uid
POST http://localhost:12499/stations
{
  "uid": "",
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
HTTP 400

# Invalid absent public
POST http://localhost:12499/stations
{
  "uid": "{{uuid3}}",
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
HTTP 400
