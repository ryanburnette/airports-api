# Far Almanac Airports API

The goal of this API is for some basic data about any airport to be available
anytime.

The initial goal for this project was to provide a quick way to determine
the time zone and UTC offset for most airports.

## Usage

Find an airport by ICAO identifier.

Make a request.

```sh
curl http://airports.api.faralmanac.com/icao/katl.json
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
