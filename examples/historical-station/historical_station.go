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

	weatherSamples, err := c.HistoricalStation(climacell.ForecastArgs{
		Location:   &climacell.LatLon{Lat: 41.9742, Lon: -87.9073},
		UnitSystem: "us",
		Fields:     []string{"temp"},
		Start:      time.Now().Add(-24 * time.Hour),
		End:        time.Now(),
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
