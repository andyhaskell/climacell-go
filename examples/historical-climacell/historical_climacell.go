package main

import (
	"log"
	"os"
	"time"

	"github.com/andyhaskell/climacell-go"
)

func main() {
	var c *climacell.Client
	c = climacell.New(os.Getenv("CLIMACELL_API_KEY"))

	weatherSamples, err := c.HistoricalClimaCell(climacell.ForecastArgs{
		Location:   &climacell.LatLon{Lat: 42.3826, Lon: -71.1460},
		UnitSystem: "si",
		Fields:     []string{"temp,no2,road_risk,fire_index"},
		Start:      time.Now().Add(-7 * time.Hour),
		End:        time.Now().Add(-1 * time.Hour),
		Timestep:   5,
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

		no2, ok := w.NO2.GetValue()
		if ok {
			log.Printf("The no2 at %s is %f  %s",
				w.ObservationTime.Value, no2, w.NO2.Units)
		} else {
			log.Printf("The no2 at %s is unavailable", w.ObservationTime.Value)
		}

		roadRisk, ok := w.RoadRisk.GetValue()
		if ok {
			log.Printf("The road risk at %s is %s",
				w.ObservationTime.Value, roadRisk)
		} else {
			log.Printf("The road risk at %s is unavailable", w.ObservationTime.Value)
		}

		fireIndex, ok := w.FireIndex.GetValue()
		if ok {
			log.Printf("The fire index at %s is %f  %s",
				w.ObservationTime.Value, fireIndex, w.FireIndex.Units)
		} else {
			log.Printf("The fire index at %s is unavailable", w.ObservationTime.Value)
		}

	}
}
