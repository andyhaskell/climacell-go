package main

import (
	"log"
	"os"

	"github.com/andyhaskell/climacell-go"
)

func main() {
	var c *climacell.Client
	c = climacell.New(os.Getenv("CLIMACELL_API_KEY"))

	realTime, err := c.RealTime(climacell.ForecastArgs{
		Location:   &climacell.LatLon{Lat: 42.3826, Lon: -71.1460},
		UnitSystem: "si",
		Fields:     []string{"temp,no2,road_risk,fire_index"},
	})
	if err != nil {
		log.Fatalf("error getting forecast data: %v", err)
	}

	temp, ok := realTime.Temp.GetValue()
	if ok {
		log.Printf("The temperature at %s is %f degrees %s",
			realTime.ObservationTime.Value, temp, realTime.Temp.Units)
	} else {
		log.Printf("The temperature at %s is unavailable", realTime.ObservationTime.Value)
	}

	no2, ok := realTime.NO2.GetValue()
	if ok {
		log.Printf("The no2 at %s is %f  %s",
			realTime.ObservationTime.Value, no2, realTime.NO2.Units)
	} else {
		log.Printf("The no2 at %s is unavailable", realTime.ObservationTime.Value)
	}

	roadRisk, ok := realTime.RoadRisk.GetValue()
	if ok {
		log.Printf("The road risk at %s is %s",
			realTime.ObservationTime.Value, roadRisk)
	} else {
		log.Printf("The road risk at %s is unavailable", realTime.ObservationTime.Value)
	}

	fireIndex, ok := realTime.FireIndex.GetValue()
	if ok {
		log.Printf("The fire index at %s is %f  %s",
			realTime.ObservationTime.Value, fireIndex, realTime.FireIndex.Units)
	} else {
		log.Printf("The fire index at %s is unavailable", realTime.ObservationTime.Value)
	}
}
