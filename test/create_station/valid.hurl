# Create a new valid public station
POST http://localhost:12499/stations
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
HTTP 201

# Create a new valid private station
POST http://localhost:12499/stations
{
  "uid": "{{uuid2}}",
  "public": false,
  "connectors": [
    {
      "id": "connector_1",
      "type": "CCS",
      "max_power_kw": 50.0
    }
  ]
}
HTTP 201
