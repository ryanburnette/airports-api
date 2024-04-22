# Airports API

An API endpoint for getting basic information about airports. Works great in
aviation applications as a central data source with an API wrapper.

## Usage

Find an airport by ICAO identifier.

Make a request to the API. Let's say we want some information about [KATL][1].

```sh
curl https://ryanburnette.github.io/airports-api/icao/katl.json
```

Get a response.

```json
{
    "_csv_row": 9,
    "airport_name": "Hartsfield Jackson Atlanta Intl",
    "city": "Atlanta",
    "country": "United States",
    "iata": "ATL",
    "icao": "KATL",
    "latitude": 33.636719,
    "longitude": -84.428067,
    "elevation": 1026,
    "utc_offset": -5,
    "_class": "A",
    "timezone": "America/New_York"
}
```

## Static Generation

The JSON files can be generated from the CSV with the included go script:

```sh
go run bin/airports-csv-to-json.go --csv ./airport-data/airports.csv
```

```text
reading ./airport-data/airports.csv
   wrote icao/.json
   ...
wrote errors.log for 1000 airports without an ICAO designation
wrote icao.json
```

## Contribute

Please fork and make a pull request to contribute.

## License

Apache2

[1]: https://ryanburnette.github.io/airports-api/icao/katl.json
