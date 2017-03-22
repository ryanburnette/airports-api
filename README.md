# Airports API

An API endpoint for getting basic information about airports. Works great in
aviation applications as a central data source with an API wrapper.

## Usage

Find an airport by ICAO identifier.

Make a request.

```sh
curl https://airports-api.s3-us-west-2.amazonaws.com//icao/katl.json
```

Get a response.

```json
{  
   "airport_name":"Hartsfield Jackson Atlanta Intl",
   "city":"Atlanta",
   "country":"United States",
   "iata":"ATL",
   "icao":"KATL",
   "latitude":33.636719,
   "longitude":-84.428067,
   "elevation":1026,
   "utc_offset":-5,
   "_class":"A",
   "timezone":"America/New_York"
}
```

## Middleman

The API is a static resource built on [Middleman](http://middlemanapp.com).

## Contribute

Please fork and make a pull request to contribute.

## License

Apache2
