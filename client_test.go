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
				WeatherType: WeatherType{
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

func historicalStationHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var temp = 15.10
		data := []HistoricalStation{
			{
				WeatherType: WeatherType{
					Temp: &FloatValue{
						Value: &temp,
					}},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Printf("error")
		}
	}
	return http.HandlerFunc(fn)
}

func historicalClimaCellHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		m, _ := url.ParseQuery(r.URL.RawQuery)
		lat, err := strconv.ParseFloat(m["lat"][0], 64)
		lon, err := strconv.ParseFloat(m["lon"][0], 64)
		var temp = 15.10
		data := []HistoricalClimaCell{
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
				WeatherType: WeatherType{
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

func realTimeHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		m, _ := url.ParseQuery(r.URL.RawQuery)
		lat, err := strconv.ParseFloat(m["lat"][0], 64)
		lon, err := strconv.ParseFloat(m["lon"][0], 64)
		var temp = 15.10
		data := RealTime{

			BaseResponseType: BaseResponseType{
				LatLon: LatLon{
					Lat: lat,
					Lon: lon,
				},
				ObservationTime: DateValue{
					Value: time.Now(),
				},
			},
			WeatherType: WeatherType{
				Temp: &FloatValue{
					Value: &temp,
				},
			},
			AirQualityType: AirQualityType{},
			RoadRiskType:   RoadRiskType{},
			FireIndexType:  FireIndexType{},
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
	mux.Handle("/weather/forecast/hourly", hourlyForecastHandler())
	mux.Handle("/weather/historical/station", historicalStationHandler())
	mux.Handle("/weather/historical/climacell", historicalClimaCellHandler())
	mux.Handle("/weather/realtime", realTimeHandler())
	server := httptest.NewServer(mux)
	log.Printf("Test server listening on %s...\n", server.URL)
	return server
}

func TestHourlyForecastEndpoint(t *testing.T) {
	server := serverMock()
	defer server.Close()

	client := New("test_api_key")
	client.baseURL = server.URL

	forecast, err := client.HourlyForecast(ForecastArgs{
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

	if err != nil {
		t.Errorf("Hourly forecast returned an unexpected error: %v", err)
	}

	value, _ := forecast[0].Temp.GetValue()
	expectedTemp := 15.10
	if expectedTemp != value {
		t.Errorf("Did not get expected result. Wanted %f, got: %f\n", expectedTemp, value)
	}
}

func TestHistoricalStationEndpoint(t *testing.T) {
	server := serverMock()
	defer server.Close()

	client := New("test_api_key")
	client.baseURL = server.URL

	historical, err := client.HistoricalStation(ForecastArgs{
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

	if err != nil {
		t.Errorf("Hourly forecast returned an unexpected error: %v", err)
	}

	value, _ := historical[0].Temp.GetValue()
	expectedTemp := 15.10
	if expectedTemp != value {
		t.Errorf("Did not get expected result. Wanted %f, got: %f\n", expectedTemp, value)
	}
}

func TestHistoricalClimaCellEndpoint(t *testing.T) {
	server := serverMock()
	defer server.Close()

	client := New("test_api_key")
	client.baseURL = server.URL

	historical, err := client.HistoricalClimaCell(ForecastArgs{
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

	if err != nil {
		t.Errorf("Hourly forecast returned an unexpected error: %v", err)
	}

	value, _ := historical[0].Temp.GetValue()
	expectedTemp := 15.10
	if expectedTemp != value {
		t.Errorf("Did not get expected result. Wanted %f, got: %f\n", expectedTemp, value)
	}
}

func TestRealTimeEndpoint(t *testing.T) {
	server := serverMock()
	defer server.Close()

	client := New("test_api_key")
	client.baseURL = server.URL

	realTime, err := client.RealTime(ForecastArgs{
		Location: LatLon{
			Lat: 11.3,
			Lon: 52.4,
		},
		UnitSystem: "",
		Fields:     nil,
	})

	if err != nil {
		t.Errorf("Hourly forecast returned an unexpected error: %v", err)
	}

	value, _ := realTime.Temp.GetValue()
	expectedTemp := 15.10
	if expectedTemp != value {
		t.Errorf("Did not get expected result. Wanted %f, got: %f\n", expectedTemp, value)
	}
}
