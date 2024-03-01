POST http://localhost:12499/stations
{
  "uid": "{{uuid4}}",
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

# Duplicate uid
POST http://localhost:12499/stations
{
  "uid": "{{uuid4}}",
  "public": false,
  "connectors": [
    {
      "id": "connector_1",
      "type": "CCS",
      "max_power_kw": 50.0
    }
  ]
}
HTTP 409
[Asserts]
jsonpath "$.error" == "Please, try providing alternative input: the provided station data is using duplicate value for uid"
