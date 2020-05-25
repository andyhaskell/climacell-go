# ClimaCell API Go client

This is the beginning of a Go client for the [ClimaCell API](https://www.climacell.co/weather-api/). It is currently in very early development stage. Docs coming soon!

### 📢 CONTRIBUTORS AND MAINTAINERS WANTED!

I would love your contributions to this ClimaCell Go client, including:

💻 Pull requests
🔭 Opening new issues and feature requests
📝 Documentation and tutorials

#### Disclaimer: This is not an official ClimaCell project; original author of this project is not affiliated with ClimaCell.

### Installing
Use `go get` to retrieve the library and add it to the your `GOPATH` workspace, or project's Go module dependencies.   
```
go get github.com/andyhaskell/climacell-go
```

### Usage example
```
package main

import (
	"github.com/andyhaskell/climacell-go"
	"log"
	"os"
	"time"
)

func main() {
	var c *climacell.Client
	c = climacell.New(os.Getenv("CLIMACELL_API_KEY"))

	weatherSamples, err := c.HourlyForecast(climacell.ForecastArgs{
		Location:     &climacell.LatLon{Lat: 42.3826, Lon: -71.146},
		UnitSystem: "us",
		Fields:     []string{"temp"},
		Start:  time.Now(),
		End:    time.Now().Add(24*time.Hour),
	})

	if err != nil {
		log.Fatalf("error getting forecast data: %v", err)
	}

	for _, w := range weatherSamples {
		log.Printf("The temperature at %s is %f degrees %s\n",
			w.ObservationTime.Value, *w.Temp.Value, w.Temp.Units)
	}
}
```
