package climacell

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// not real weather data (which is why the sample is from the geographically
// impossible coordinates 91 degrees north and 181 degrees west), but the same
// format as a real API response
var everyWeatherField = []byte(`
{
  "lat":                         91.128,
  "lon":                         -181.250,
  "temp":                        {"value": 10,       "units": "C"},
  "feels_like":                  {"value": 10,       "units": "C"},
  "dewpoint":                    {"value": -0.50,    "units": "C"},
  "wind_speed":                  {"value": 5,        "units": "beaufort"},
  "wind_gust":                   {"value": 10,       "units": "beaufort"},
  "baro_pressure":               {"value": 20.71,    "units": "inHg"},
  "visibility":                  {"value": 6.21371,  "units": "mi"},
  "humidity":                    {"value": 20.71,    "units": "%"},
  "wind_direction":              {"value": 250.25,   "units": "degrees"},
  "precipitation":               {"value": 10,       "units": "mm/hr"},
  "precipitation_type":          {"value": "rain"},
  "cloud_cover":                 {"value": 12.8,     "units": "%"},
  "cloud_ceiling":               {"value": 6000,     "units": "ft"},
  "cloud_base":                  {"value": 5000,     "units": "ft"},
  "surface_shortwave_radiation": {"value": 250.1123, "units": "w/sqm"},
  "fire_index":                  {"value": 3.6},
  "sunrise":                     {"value": "2020-04-12T12:34:56.789Z"},
  "sunset":                      {"value": "2020-04-12T23:45:56.789Z"},
  "moon_phase":                  {"value": "waning_gibbous"},
  "weather_code":                {"value": "mostly_clear"},
  "epa_aqi":                     {"value": 25},
  "epa_primary_pollutant":       {"value": "pm25"},
  "china_aqi":                   {"value": 12},
  "china_primary_pollutant":     {"value": "pm25"},
  "pm25":                        {"value": 10,       "units": "µg/m3"},
  "pm10":                        {"value": 15,       "units": "µg/m3"},
  "o3":                          {"value": 8,        "units": "ppb"},
  "no2":                         {"value": 40,       "units": "ppb"},
  "co":                          {"value": 3,        "units": "ppm"},
  "so2":                         {"value": 1,        "units": "ppb"},
  "epa_health_concern":          {"value": "Good"},
  "china_health_concern":        {"value": "Good"},
  "observation_time":            {"value": "2020-04-12T12:00:00.000Z"}
}`)

// TestDeserializeWeatherWithAllFields validates that if we try to deserialize
// a Weather sample with all fields non-null, the deserialization succeeds.
func TestDeserializeWeatherWithAllFields(t *testing.T) {
	var w Weather
	require.NoError(t, json.Unmarshal(everyWeatherField, &w))

	assert.Equal(t, 91.128, w.Lat)
	assert.Equal(t, -181.250, w.Lon)
	expObservationTime, err := time.Parse(time.RFC3339, "2020-04-12T12:00:00.000Z")
	require.NoError(t, err)
	if assert.NotNil(t, w.ObservationTime) {
		assert.WithinDuration(t, expObservationTime, w.ObservationTime.Value, time.Second)
	}

	if assert.NotNil(t, w.Temp) {
		assert.EqualValues(t, w.Temp.Value, 10)
	}
	if assert.NotNil(t, w.FeelsLike) {
		assert.EqualValues(t, w.FeelsLike.Value, 10)
	}
	if assert.NotNil(t, w.DewPoint) {
		assert.EqualValues(t, w.DewPoint.Value, -0.50)
	}
	if assert.NotNil(t, w.WindSpeed) {
		assert.EqualValues(t, w.WindSpeed.Value, 5)
	}
	if assert.NotNil(t, w.WindGust) {
		assert.EqualValues(t, w.WindGust.Value, 10)
	}
	if assert.NotNil(t, w.BaroPressure) {
		assert.EqualValues(t, w.BaroPressure.Value, 20.71)
	}
	if assert.NotNil(t, w.Visibility) {
		assert.EqualValues(t, w.Visibility.Value, 6.21371)
	}
	if assert.NotNil(t, w.Humidity) {
		assert.EqualValues(t, w.Humidity.Value, 20.71)
	}
	if assert.NotNil(t, w.WindDirection) {
		assert.EqualValues(t, w.WindDirection.Value, 250.25)
	}
	if assert.NotNil(t, w.Precipitation) {
		assert.EqualValues(t, w.Precipitation.Value, 10)
	}
	if assert.NotNil(t, w.PrecipitationType) {
		assert.EqualValues(t, w.PrecipitationType.Value, "rain")
	}
	if assert.NotNil(t, w.CloudCover) {
		assert.EqualValues(t, w.CloudCover.Value, 12.8)
	}
	if assert.NotNil(t, w.CloudCeiling) {
		assert.EqualValues(t, w.CloudCeiling.Value, 6000)
	}
	if assert.NotNil(t, w.CloudBase) {
		assert.EqualValues(t, w.CloudBase.Value, 5000)
	}
	if assert.NotNil(t, w.SurfaceShortwaveRadiation) {
		assert.EqualValues(t, w.SurfaceShortwaveRadiation.Value, 250.1123)
	}
	if assert.NotNil(t, w.FireIndex) {
		assert.EqualValues(t, w.FireIndex.Value, 3.6)
	}

	expSunrise, err := time.Parse(time.RFC3339, "2020-04-12T12:34:56.789Z")
	require.NoError(t, err)
	if assert.NotNil(t, w.Sunrise) {
		assert.WithinDuration(t, expSunrise, w.Sunrise.Value, time.Second)
	}
	expSunset, err := time.Parse(time.RFC3339, "2020-04-12T23:45:56.789Z")
	require.NoError(t, err)
	if assert.NotNil(t, w.Sunset) {
		assert.WithinDuration(t, expSunset, w.Sunset.Value, time.Second)
	}

	if assert.NotNil(t, w.MoonPhase) {
		assert.EqualValues(t, w.MoonPhase.Value, "waning_gibbous")
	}
	if assert.NotNil(t, w.WeatherCode) {
		assert.EqualValues(t, w.WeatherCode.Value, "mostly_clear")
	}
	if assert.NotNil(t, w.EpaAQI) {
		assert.EqualValues(t, w.EpaAQI.Value, 25)
	}
	if assert.NotNil(t, w.EPAPrimaryPollutant) {
		assert.EqualValues(t, w.EPAPrimaryPollutant.Value, "pm25")
	}
	if assert.NotNil(t, w.EPAHealthConcern) {
		assert.EqualValues(t, w.EPAHealthConcern.Value, "Good")
	}
	if assert.NotNil(t, w.ChinaAQI) {
		assert.EqualValues(t, w.ChinaAQI.Value, 12)
	}
	if assert.NotNil(t, w.ChinaPrimaryPollutant) {
		assert.EqualValues(t, w.ChinaPrimaryPollutant.Value, "pm25")
	}
	if assert.NotNil(t, w.ChinaHealthConcern) {
		assert.EqualValues(t, w.ChinaHealthConcern.Value, "Good")
	}
	if assert.NotNil(t, w.PMTwoPointFive) {
		assert.EqualValues(t, w.PMTwoPointFive.Value, 10)
	}
	if assert.NotNil(t, w.PMTen) {
		assert.EqualValues(t, w.PMTen.Value, 15)
	}
	if assert.NotNil(t, w.O3) {
		assert.EqualValues(t, w.O3.Value, 8)
	}
	if assert.NotNil(t, w.NO2) {
		assert.EqualValues(t, w.NO2.Value, 40)
	}
	if assert.NotNil(t, w.CO) {
		assert.EqualValues(t, w.CO.Value, 3)
	}
	if assert.NotNil(t, w.SO2) {
		assert.EqualValues(t, w.SO2.Value, 1)
	}
}

// not real weather data (which is why the sample is from the geographically
// impossible coordinates 91 degrees north and 181 degrees west), but the same
// format as a real API response
var minimalWeatherData = []byte(`
{
  "lat":              91.128,
  "lon":              -181.250,
  "temp":             {"value": 10, "units": "C"},
  "observation_time": {"value": "2020-04-12T12:00:00.000Z"}
}`)

// TestDeserializeWeatherWithAllFields validates that if we try to deserialize
// a Weather sample with almost all fields non-null, the deserialization
// succeeds, with nil values for all fields that were absent.
func TestDeserializeWeatherWithAlmostAllFieldsAbsent(t *testing.T) {
	var w Weather
	require.NoError(t, json.Unmarshal(minimalWeatherData, &w))

	assert.Equal(t, 91.128, w.Lat)
	assert.Equal(t, -181.250, w.Lon)
	expObservationTime, err := time.Parse(time.RFC3339, "2020-04-12T12:00:00.000Z")
	require.NoError(t, err)
	if assert.NotNil(t, w.ObservationTime) {
		assert.WithinDuration(t, expObservationTime, w.ObservationTime.Value, time.Second)
	}

	if assert.NotNil(t, w.Temp) {
		assert.EqualValues(t, w.Temp.Value, 10)
	}

	assert.Nil(t, w.FeelsLike)
	assert.Nil(t, w.DewPoint)
	assert.Nil(t, w.WindSpeed)
	assert.Nil(t, w.WindGust)
	assert.Nil(t, w.BaroPressure)
	assert.Nil(t, w.Visibility)
	assert.Nil(t, w.Humidity)
	assert.Nil(t, w.WindDirection)
	assert.Nil(t, w.Precipitation)
	assert.Nil(t, w.PrecipitationType)
	assert.Nil(t, w.CloudCover)
	assert.Nil(t, w.CloudCeiling)
	assert.Nil(t, w.CloudBase)
	assert.Nil(t, w.SurfaceShortwaveRadiation)
	assert.Nil(t, w.FireIndex)
	assert.Nil(t, w.Sunrise)
	assert.Nil(t, w.Sunset)
	assert.Nil(t, w.MoonPhase)
	assert.Nil(t, w.WeatherCode)
	assert.Nil(t, w.EpaAQI)
	assert.Nil(t, w.EPAPrimaryPollutant)
	assert.Nil(t, w.EPAHealthConcern)
	assert.Nil(t, w.ChinaAQI)
	assert.Nil(t, w.ChinaPrimaryPollutant)
	assert.Nil(t, w.ChinaHealthConcern)
	assert.Nil(t, w.PMTwoPointFive)
	assert.Nil(t, w.PMTen)
	assert.Nil(t, w.O3)
	assert.Nil(t, w.NO2)
	assert.Nil(t, w.CO)
	assert.Nil(t, w.SO2)
}
