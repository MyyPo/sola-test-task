# Get error output in Spanish for too short title
POST http://localhost:12499/stations
Accept-Language: fr-CH, fr;q=0.9, es;q=0.8, de;q=0.7, *;q=0.5
{
  "uid": "{{uuid3}}",
  "public": true,
  "title": "x",
  "description": "Description of station 1",
  "address": "123 Main St",
  "coordinates": [1.234, 2.345],
  "connectors": [
    {
      "id": "connector_1",
      "type": "CCS",
      "max_power_kw": 50.0
    }
  ]
}
HTTP 400
[Asserts]
jsonpath "$.error" == "Entrada de solicitud incorrecta. Por favor, proporcione datos válidos: Title debe tener al menos 3 caracteres de longitud"

# Get error output in English for invalid latitude
POST http://localhost:12499/stations
Accept-Language: en-US,en;q=0.5
{
  "uid": "{{uuid3}}",
  "public": true,
  "title": "Some title",
  "description": "Description of station 1",
  "address": "123 Main St",
  "coordinates": [-122.111, 2.345],
  "connectors": [
    {
      "id": "connector_1",
      "type": "CCS",
      "max_power_kw": 50.0
    }
  ]
}
HTTP 400
[Asserts]
jsonpath "$.error" == "Bad request input. Please, provide valid data: coordinates must contain valid latitude coordinates"
