package climacell

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func hourlyForecastHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		m, _ := url.ParseQuery(r.URL.RawQuery)
		lat, err := strconv.ParseFloat(m["lat"][0], 64)
		lon, err := strconv.ParseFloat(m["lon"][0], 64)
		var temp = 15.10
		data := []HourlyForecast{
			{
				BaseResponseType: BaseResponseType{
					LatLon: LatLon{
						Lat: lat,
						Lon: lon,
					},
					ObservationTime: DateValue{
						Value: time.Now(),
					},
				},
				WeatherType:    WeatherType{
					Temp: &FloatValue{
						Value: &temp,
					},
				},
				AirQualityType: AirQualityType{},
				RoadRiskType:   RoadRiskType{},
				FireIndexType:  FireIndexType{}},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Printf("error")
		}
	}
	return http.HandlerFunc(fn)
}

func serverMock() *httptest.Server {
	mux := http.NewServeMux()
	const addr = "http://localhost:12345"
	log.Printf("Now listening on %s...\n", addr)
	mux.Handle("/weather/forecast/hourly", hourlyForecastHandler())
	server := httptest.NewServer(mux)
	return server
}

func TestHourlyForecastEndpoint(t *testing.T) {
	server := serverMock()
	defer server.Close()

	client := New("test_api_key")
	client.baseURL = server.URL

	forecast, _ := client.HourlyForecast(ForecastArgs{
		Location: LatLon{
			Lat: 11.3,
			Lon: 52.4,
		},
		Start:      time.Time{},
		End:        time.Time{},
		Timestep:   0,
		UnitSystem: "",
		Fields:     nil,
	})

	value, _ := forecast[0].Temp.GetValue()
	expectedTemp := 15.10
	if expectedTemp != value {
		t.Errorf("Did not get expected result. Wanted %f, got: %f\n", expectedTemp, value)
	}
}
