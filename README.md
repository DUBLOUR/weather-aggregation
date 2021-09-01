# weather-aggregation

The application aggregates different weather API sources (by default is [OpenWeather](https://openweathermap.org/), [Stormglass.io](https://stormglass.io/), [Weather API](https://www.weatherapi.com/)) and return local temperature, humidity and speed of wind for the requested city.

##### Setup: 

`git clone https://github.com/DUBLOUR/weather-aggregation.git && cd weather-aggregation`

`go build -o bin/cli cmd/cli/main.go`

##### Usage:

`go run cmd/cli/main.go Berlin`

or

`./bin/cli Paris`

**Note:** for easy launch all API keys are distributed with code. If you will launch the app more than 2-3 times please create your keys and change `pkg/*/config.go` files.

The possible output of the program:
```
[./bin/cli Kyiv]
2021/09/01 13:42:55 Set `data/log` as LogFile
2021/09/01 13:42:55 Set `data/city.json` as MetricDbFile
2021/09/01 13:42:56 [400:] https://api.weatherapi.com/v1/current.json?aqi=no&key=271c665a32d242b7a65113844210708&q=Kyiv
2021/09/01 13:43:02 [200:] https://geocode.xyz?auth=866303333362274101661x21265&json=1&locate=Kyiv
2021/09/01 13:43:04 [200:] https://api.stormglass.io/v2/weather/point?end=1630492982&lat=50.44676&lng=30.51313&params=airTemperature%2CwindSpeed%2Chumidity&source=sg&start=1630492982
{"Temp":20.45,"Hum":53,"Wind":3.8}
```

The trace-log describing the startup and tracking HTTP requests is printed to stdout.

The permanent log is writing to `data/log`.

Analytics (frequency of city requests) storage in `data/city.json`.


# Design Review

### pkg/weatherMaster

Module for distributing requests to third-party APIs. Accepts a class corresponding to the `ISource` interface (examples of adapters in `adapters.go`).
Uses and specifies the list of services used in `internal/server/main.go`.


### pkg/generalApiReader

Unified http request method with error handling and logging. Accepts `http.Request` and returns processed data in the passed format.
It is used to request `pkg/weatherapi`, `pkg/openweathermap`, `pkg/stormglass` and `pkg/geocode` (the last is used to determine the coordinates of a given city).

### internal/server

Initializing `weatherMaster`, specifying the list of third-party APIs used, configuring logging and analytics.

