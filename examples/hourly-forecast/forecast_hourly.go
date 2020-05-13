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
		temp, ok := w.Temp.GetValue()
		if ok {
			log.Printf("The temperature at %s is %f degrees %s",
				w.ObservationTime.Value, temp, w.Temp.Units)
		} else {
			log.Printf("The temperature at %s is unavailable", w.ObservationTime.Value)
		}
	}
}