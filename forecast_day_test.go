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
var dailyForecastAllFields = []byte(`{
  "temp": [
    {
      "observation_time": "2020-05-01T00:00:00Z",
      "min": {"value": 11.23, "units": "C"}
    },
    {
      "observation_time": "2020-05-01T01:00:00Z",
      "max": {"value": 23.58, "units": "C"}
    }
  ],
  "precipitation_accumulation": {
    "value": 0.1123,
    "units": "in"
  },
  "precipitation": [
    {
      "observation_time": "2020-05-01T00:00:00Z",
      "min": {"value": 0.128, "units": "mm/hr"}
    },
    {
      "observation_time": "2020-05-01T01:00:00Z",
      "max": {"value": 0.250, "units": "mm/hr"}
    }
  ],
  "precipitation_probability": {
    "value": 10,
    "units": "%"
  },
  "feels_like": [
    {
      "observation_time": "2020-05-01T00:00:00Z",
      "min": {"value": 11.23, "units": "C"}
    },
    {
      "observation_time": "2020-05-01T01:00:00Z",
      "max": {"value": 23.58, "units": "C"}
    }
  ],
  "humidity": [
    {
      "observation_time": "2020-05-01T00:00:00Z",
      "min": {"value": 11, "units": "%"}
    },
    {
      "observation_time": "2020-05-01T01:00:00Z",
      "max": {"value": 23, "units": "%"}
    }
  ],
  "baro_pressure": [
    {
      "observation_time": "2020-05-01T00:00:00Z",
      "min": {"value": 70000.58, "units": "Pa"}
    },
    {
      "observation_time": "2020-05-01T01:00:00Z",
      "max": {"value": 75000.13, "units": "Pa"}
    }
  ],
  "wind_speed": [
    {
      "observation_time": "2020-05-01T00:00:00Z",
      "min": {"value": 21, "units": "beaufort"}
    },
    {
      "observation_time": "2020-05-01T01:00:00Z",
      "max": {"value": 34, "units": "beaufort"}
    }
  ],
  "wind_direction": [
    {
      "observation_time": "2020-05-01T00:00:00Z",
      "min": {"value": 55.89, "units": "degrees"}
    },
    {
      "observation_time": "2020-05-01T01:00:00Z",
      "max": {"value": 144.00, "units": "degrees"}
    }
  ],
  "visibility": [
    {
      "observation_time": "2020-05-01T00:00:00Z",
      "min": {"value": 11.23, "units": "mi"}
    },
    {
      "observation_time": "2020-05-01T01:00:00Z",
      "max": {"value": 58, "units": "mi"}
    }
  ],
  "sunrise":          {"value": "2020-05-01T11:23:58.123Z"},
  "sunset":           {"value": "2020-05-01T12:34:56.789Z"},
  "moon_phase":       {"value": "first_quarter"},
  "weather_code":     {"value": "mostly_clear"},
  "observation_time": {"value": "2020-05-01"},
  "lat":              91,
  "lon":              -181
}`)

// TestDeserializeAllFieldsPresent tests that if all fields are present on a
// ForecastDay's JSON, all of those fields can be retrieved.
func TestDeserializeForecastAllFieldsPresent(t *testing.T) {
	var f ForecastDay
	require.NoError(t, json.Unmarshal(dailyForecastAllFields, &f))

	assertMinMax(t, 11.23, 23.58, f.Temp)
	assertMinMax(t, 0.128, 0.250, f.Precipitation)
	assertMinMax(t, 11.23, 23.58, f.FeelsLike)
	assertMinMax(t, 11, 23, f.Humidity)
	assertMinMax(t, 70000.58, 75000.13, f.BaroPressure)
	assertMinMax(t, 21, 34, f.WindSpeed)
	assertMinMax(t, 55.89, 144.00, f.WindDirection)
	assertMinMax(t, 11.23, 58, f.Visibility)

	if precipitationAcc, ok := f.PrecipitationAccumulation.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 0.1123, precipitationAcc)
	}
	if precipitationProb, ok := f.PrecipitationProbability.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 10, precipitationProb)
	}

	expSunrise, err := time.Parse(time.RFC3339, "2020-05-01T11:23:58.123Z")
	require.NoError(t, err)
	if tm, ok := f.Sunrise.GetValue(); assert.True(t, ok) {
		assert.WithinDuration(t, expSunrise, tm, time.Second)
	}
	expSunset, err := time.Parse(time.RFC3339, "2020-05-01T12:34:56.789Z")
	require.NoError(t, err)
	if tm, ok := f.Sunset.GetValue(); assert.True(t, ok) {
		assert.WithinDuration(t, expSunset, tm, time.Second)
	}

	if moonPhase, ok := f.MoonPhase.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, "first_quarter", moonPhase)
	}
	if weatherCode, ok := f.WeatherCode.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, "mostly_clear", weatherCode)
	}

	expObservationTime, err := time.Parse("2006-01-02", "2020-05-01")
	require.NoError(t, err)
	assert.EqualValues(t, expObservationTime, f.ObservationTime.Value)
	assert.EqualValues(t, 91, f.Lat)
	assert.EqualValues(t, -181, f.Lon)
}

func assertMinMax(t *testing.T, expectedMin, expectedMax interface{}, got *ForecastMinAndMax) {
	min, minOK := got.Min().GetValue()
	max, maxOK := got.Max().GetValue()

	if expectedMin == nil {
		assert.False(t, minOK, "should not have min, got %f", min)
	} else if assert.True(t, minOK) {
		assert.EqualValues(t, expectedMin, min)
	}

	if expectedMax == nil {
		assert.False(t, maxOK, "should not have max, got %f", max)
	} else if assert.True(t, maxOK) {
		assert.EqualValues(t, expectedMax, max)
	}
}
