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
  "road_risk":                   {"value": "low_risk"},
  "road_risk_score":             {"value": "Low Risk"},
  "road_risk_confidence":        {"value": 100 },
  "road_risk_conditions":        {"value": "Low visibility"},
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
	assert.EqualValues(t, expObservationTime, w.ObservationTime.Value)

	if temp, ok := w.Temp.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 10, temp)
	}
	if feelsLike, ok := w.FeelsLike.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 10, feelsLike)
	}
	if dewPoint, ok := w.DewPoint.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, -0.50, dewPoint)
	}
	if windSpeed, ok := w.WindSpeed.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 5, windSpeed)
	}
	if windGust, ok := w.WindGust.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 10, windGust)
	}
	if baroPressure, ok := w.BaroPressure.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 20.71, baroPressure)
	}
	if visibility, ok := w.Visibility.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 6.21371, visibility)
	}
	if humidity, ok := w.Humidity.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 20.71, humidity)
	}
	if windDirection, ok := w.WindDirection.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 250.25, windDirection)
	}
	if precipitation, ok := w.Precipitation.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 10, precipitation)
	}
	if precipitationType, ok := w.PrecipitationType.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, "rain", precipitationType)
	}
	if cloudCover, ok := w.CloudCover.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 12.8, cloudCover)
	}
	if cloudCeiling, ok := w.CloudCeiling.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 6000, cloudCeiling)
	}
	if cloudBase, ok := w.CloudBase.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 5000, cloudBase)
	}
	if ssw, ok := w.SurfaceShortwaveRadiation.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 250.1123, ssw)
	}
	if fireIndex, ok := w.FireIndex.GetValue(); ok {
		assert.EqualValues(t, 3.6, fireIndex)
	}

	expSunrise, err := time.Parse(time.RFC3339, "2020-04-12T12:34:56.789Z")
	require.NoError(t, err)
	if sunrise, ok := w.Sunrise.GetValue(); assert.True(t, ok) {
		assert.WithinDuration(t, expSunrise, sunrise, time.Second)
	}
	expSunset, err := time.Parse(time.RFC3339, "2020-04-12T23:45:56.789Z")
	require.NoError(t, err)
	if sunset, ok := w.Sunset.GetValue(); assert.True(t, ok) {
		assert.WithinDuration(t, expSunset, sunset, time.Second)
	}

	if moonPhase, ok := w.MoonPhase.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, "waning_gibbous", moonPhase)
	}
	if weatherCode, ok := w.WeatherCode.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, "mostly_clear", weatherCode)
	}

	if roadRisk, ok := w.RoadRisk.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, "low_risk", roadRisk)
	}
	if roadRiskScore, ok := w.RoadRiskScore.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, "Low Risk", roadRiskScore)
	}
	if roadRiskConfidence, ok := w.RoadRiskConfidence.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 100, roadRiskConfidence)
	}
	if roadRiskConditions, ok := w.RoadRiskConditions.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, "Low visibility", roadRiskConditions)
	}

	if epaAQI, ok := w.EpaAQI.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 25, epaAQI)
	}
	if epaPrimaryPollutant, ok := w.EPAPrimaryPollutant.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, "pm25", epaPrimaryPollutant)
	}
	if epaHealthConcern, ok := w.EPAHealthConcern.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, "Good", epaHealthConcern)
	}
	if chinaAQI, ok := w.ChinaAQI.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 12, chinaAQI)
	}
	if chinaPrimaryPollutant, ok := w.ChinaPrimaryPollutant.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, "pm25", chinaPrimaryPollutant)
	}
	if chinaHealthConcern, ok := w.ChinaHealthConcern.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, "Good", chinaHealthConcern)
	}
	if pmTwoPointFive, ok := w.PMTwoPointFive.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 10, pmTwoPointFive)
	}
	if pmTen, ok := w.PMTen.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 15, pmTen)
	}
	if o3, ok := w.O3.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 8, o3)
	}
	if no2, ok := w.NO2.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 40, no2)
	}
	if co, ok := w.CO.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 3, co)
	}
	if so2, ok := w.SO2.GetValue(); assert.True(t, ok) {
		assert.EqualValues(t, 1, so2)
	}
}

// not real weather data (which is why the sample is from the geographically
// impossible coordinates 91 degrees north and 181 degrees west), but the same
// format as a real API response
var minimalWeatherData = []byte(`
{
  "lat":              91.128,
  "lon":              -181.250,
  "observation_time": {"value": "2020-04-12T12:00:00.000Z"}
}`)

// TestDeserializeWeatherWithAllFieldsAbsent validates that if we try to
// deserialize a Weather sample with all nullable fields absent, the
// deserialization succeeds, with nil values for any absent fields.
func TestDeserializeWeatherWithAllFieldsAbsent(t *testing.T) {
	var w Weather
	require.NoError(t, json.Unmarshal(minimalWeatherData, &w))

	assert.Equal(t, 91.128, w.Lat)
	assert.Equal(t, -181.250, w.Lon)
	expObservationTime, err := time.Parse(time.RFC3339, "2020-04-12T12:00:00.000Z")
	require.NoError(t, err)
	assert.EqualValues(t, expObservationTime, w.ObservationTime.Value)

	assert.Nil(t, w.Temp)
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
	assert.Nil(t, w.RoadRisk)
	assert.Nil(t, w.RoadRiskScore)
	assert.Nil(t, w.RoadRiskConfidence)
	assert.Nil(t, w.RoadRiskConditions)
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

// not real weather data (which is why the sample is from the geographically
// impossible coordinates 91 degrees north and 181 degrees west), but the same
// format as a real API response
var everyWeatherFieldNull = []byte(`
{
  "lat":                         91.128,
  "lon":                         -181.250,
  "temp":                        {"value": null, "units": "C"},
  "feels_like":                  {"value": null, "units": "C"},
  "dewpoint":                    {"value": null, "units": "C"},
  "wind_speed":                  {"value": null, "units": "beaufort"},
  "wind_gust":                   {"value": null, "units": "beaufort"},
  "baro_pressure":               {"value": null, "units": "inHg"},
  "visibility":                  {"value": null, "units": "mi"},
  "humidity":                    {"value": null, "units": "%"},
  "wind_direction":              {"value": null, "units": "degrees"},
  "precipitation":               {"value": null, "units": "mm/hr"},
  "precipitation_type":          {"value": null},
  "cloud_cover":                 {"value": null, "units": "%"},
  "cloud_ceiling":               {"value": null, "units": "ft"},
  "cloud_base":                  {"value": null, "units": "ft"},
  "surface_shortwave_radiation": {"value": null, "units": "w/sqm"},
  "fire_index":                  {"value": null},
  "sunrise":                     {"value": null},
  "sunset":                      {"value": null},
  "moon_phase":                  {"value": null},
  "weather_code":                {"value": null},
  "road_risk":                   {"value": null},
  "road_risk_score":             {"value": null},
  "road_risk_confidence":        {"value": null},
  "road_risk_conditions":        {"value": null},
  "epa_aqi":                     {"value": null},
  "epa_primary_pollutant":       {"value": null},
  "china_aqi":                   {"value": null},
  "china_primary_pollutant":     {"value": null},
  "pm25":                        {"value": null, "units": "µg/m3"},
  "pm10":                        {"value": null, "units": "µg/m3"},
  "o3":                          {"value": null, "units": "ppb"},
  "no2":                         {"value": null, "units": "ppb"},
  "co":                          {"value": null, "units": "ppm"},
  "so2":                         {"value": null, "units": "ppb"},
  "epa_health_concern":          {"value": null},
  "china_health_concern":        {"value": null},
  "observation_time":            {"value": "2020-04-12T12:00:00.000Z"}
}`)

// TestDeserializeWeatherWithAllFieldsNull validates that if we try to
// deserialize a Weather sample with almost all nullable fields present but
// with null values, the deserialization succeeds, and for each field's
// GetValue method, a false ok value is returned.
func TestDeserializeWeatherWithAllFieldsNull(t *testing.T) {
	var w Weather
	require.NoError(t, json.Unmarshal(minimalWeatherData, &w))

	assert.Equal(t, 91.128, w.Lat)
	assert.Equal(t, -181.250, w.Lon)
	expObservationTime, err := time.Parse(time.RFC3339, "2020-04-12T12:00:00.000Z")
	require.NoError(t, err)
	assert.EqualValues(t, expObservationTime, w.ObservationTime.Value)

	_, ok := w.Temp.GetValue()
	assert.False(t, ok, "Temp was present")
	_, ok = w.FeelsLike.GetValue()
	assert.False(t, ok, "FeelsLike was present")
	_, ok = w.DewPoint.GetValue()
	assert.False(t, ok, "DewPoint was present")
	_, ok = w.WindSpeed.GetValue()
	assert.False(t, ok, "WindSpeed was present")
	_, ok = w.WindGust.GetValue()
	assert.False(t, ok, "WindGust was present")
	_, ok = w.BaroPressure.GetValue()
	assert.False(t, ok, "BaroPressure was present")
	_, ok = w.Visibility.GetValue()
	assert.False(t, ok, "Visibility was present")
	_, ok = w.Humidity.GetValue()
	assert.False(t, ok, "Humidity was present")
	_, ok = w.WindDirection.GetValue()
	assert.False(t, ok, "WindDirection was present")
	_, ok = w.Precipitation.GetValue()
	assert.False(t, ok, "Precipitation was present")
	_, ok = w.PrecipitationType.GetValue()
	assert.False(t, ok, "PrecipitationType was present")
	_, ok = w.CloudCover.GetValue()
	assert.False(t, ok, "CloudCoder was present")
	_, ok = w.CloudCeiling.GetValue()
	assert.False(t, ok, "CloudCeiling was present")
	_, ok = w.CloudBase.GetValue()
	assert.False(t, ok, "CloudBase was present")
	_, ok = w.SurfaceShortwaveRadiation.GetValue()
	assert.False(t, ok, "SurfaceShortwaveRadiation was present")
	_, ok = w.FireIndex.GetValue()
	assert.False(t, ok, "FireIndex was present")
	_, ok = w.Sunrise.GetValue()
	assert.False(t, ok, "Sumrise was present")
	_, ok = w.Sunset.GetValue()
	assert.False(t, ok, "Sunset was present")
	_, ok = w.MoonPhase.GetValue()
	assert.False(t, ok, "MoonPhase was present")
	_, ok = w.WeatherCode.GetValue()
	assert.False(t, ok, "WeatherCode was present")
	_, ok = w.RoadRisk.GetValue()
	assert.False(t, ok, "road risk was present")
	_, ok = w.RoadRiskScore.GetValue()
	assert.False(t, ok, "road risk score was present")
	_, ok = w.RoadRiskConfidence.GetValue()
	assert.False(t, ok, "road risk confidence was present")
	_, ok = w.RoadRiskConditions.GetValue()
	assert.False(t, ok, "road risk conditions was present")
	_, ok = w.EpaAQI.GetValue()
	assert.False(t, ok, "EpaAQI was present")
	_, ok = w.EPAPrimaryPollutant.GetValue()
	assert.False(t, ok, "EPAPrimaryPollutant was present")
	_, ok = w.EPAHealthConcern.GetValue()
	assert.False(t, ok, "EPAHealthConcern was present")
	_, ok = w.ChinaAQI.GetValue()
	assert.False(t, ok, "ChinaAQI was present")
	_, ok = w.ChinaPrimaryPollutant.GetValue()
	assert.False(t, ok, "ChinaPrimaryPollutant was present")
	_, ok = w.ChinaHealthConcern.GetValue()
	assert.False(t, ok, "ChinaHealthConcern was present")
	_, ok = w.PMTwoPointFive.GetValue()
	assert.False(t, ok, "PMTwoPointFive was present")
	_, ok = w.PMTen.GetValue()
	assert.False(t, ok, "PMTen was present")
	_, ok = w.O3.GetValue()
	assert.False(t, ok, "O3 was present")
	_, ok = w.NO2.GetValue()
	assert.False(t, ok, "NO2 was present")
	_, ok = w.CO.GetValue()
	assert.False(t, ok, "CO was present")
	_, ok = w.SO2.GetValue()
	assert.False(t, ok, "SO2 was present")
}
