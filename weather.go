package climacell

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type BaseResponseType struct {
	LatLon
	LocationId LocationID `json:"location_id"`
	// The time when this weather sample is from.
	ObservationTime DateValue `json:"observation_time"`
}

// Weather struct and related values
//

// Weather contains the data for a single weather sample in a location, and is
// returned from the ClimaCell API's /weather/* endpoints.
//
// Any pointer field on this struct will have a nil value if either:
// - The field was not requested in the API request that this sample was
//   retrieved through.
// - Or data was not available for this field's value for the requested time
//   and location, therefore returning a null value for the field in the API
//   response.
//
// For convenience, TimeValue, FloatValue, IntValue, and StringValue structs
// all have GetValue methods so that you can check for data without checking
// whether two pointer values are non-nil, like this:
//
// temp, ok := w.Temp.GetValue()
// if !ok {
// 	/* handle a temp value being absent */
// }
// /* work with the retrieved temp value */
type WeatherType struct {
	// The temperature for this weather sample.
	Temp *FloatValue `json:"temp,omitempty"`
	// The temperature it feels like for this weather sample, based on wind
	// chill and heat window.
	FeelsLike *FloatValue `json:"feels_like,omitempty"`
	// The temperature of the dew point for this weather sample.
	DewPoint *FloatValue `json:"dewpoint,omitempty"`
	// The percent relative humidity for this weather sample.
	Humidity *FloatValue `json:"humidity,omitempty"`
	// The wind speed for this weather sample.
	WindSpeed *FloatValue `json:"wind_speed,omitempty"`
	// The direction of the wind in degrees for this weather sample, where
	// 0 degrees means the wind is going exactly north.
	WindDirection *FloatValue `json:"wind_direction,omitempty"`
	// The wind gust speed for this weather sample.
	WindGust *FloatValue `json:"wind_gust,omitempty"`
	// The surface barometric pressure for this weather sample.
	BaroPressure *FloatValue `json:"baro_pressure,omitempty"`
	// The amount of precipitation for this weather sample.
	Precipitation *FloatValue `json:"precipitation,omitempty"`
	// The type of precipitation for this weather sample. Values include
	// "none", "rain", "snow", "ice pellets", and "freezing rain".
	PrecipitationType *StringValue `json:"precipitation_type,omitempty"`
	// When this weather sample is from a forecast, the percent probability
	// of precipitation.
	PrecipitationProbability *FloatValue `json:"precipitation_probability,omitempty"`
	// The sunrise time for this location.
	Sunrise *TimeValue `json:"sunrise"`
	// The sunset time for this location.
	Sunset *TimeValue `json:"sunset"`
	// The visibility distance for this weather sample.
	Visibility *FloatValue `json:"visibility,omitempty"`
	// The percent of the sky obscured by clouds for this weather sample.
	CloudCover *FloatValue `json:"cloud_cover"`
	// The lowest height at which there are clouds for this weather sample.
	CloudBase *FloatValue `json:"cloud_base"`
	// The highest height at which there are clouds for this weather
	// sample.
	CloudCeiling *FloatValue `json:"cloud_ceiling"`
	// The amount of solar radiation reaching the surface for this weather
	// sample.
	SurfaceShortwaveRadiation *FloatValue `json:"surface_shortwave_radiation"`
	// The phase of the moon. Values include "new_moon", "waxing_crescent",
	// "first_quarter", "waxing_gibbous", "full", "waning_gibbous",
	// "third_quarter", and "waning_crescent"
	MoonPhase *StringValue `json:"moon_phase"`
	// A text description of the weather. Possible values include
	// "freezing_rain_heavy", "freezing_rain", "freezing_rain_light",
	// "freezing_drizzle", "ice_pellets_heavy", "ice_pellets",
	// "ice_pellets_light", "snow_heavy", "snow", "snow_light", "flurries",
	// "tstorm", "rain_heavy", "rain", "rain_light", "drizzle",
	// "fog_light", "fog", "cloudy", "mostly_cloudy", "partly_cloudy",
	// "mostly_clear", and "clear".
	WeatherCode *StringValue `json:"weather_code"`
}

type AirQualityType struct {
	// Amount of particulate matter smaller than 2.5 micrometers for this
	// weather sample.
	PMTwoPointFive *FloatValue `json:"pm25"`
	// Amount of particulate matter smaller than 10 micrometers for this
	// weather sample.
	PMTen *FloatValue `json:"pm10"`
	// Amount of ozone in the air for this weather sample.
	O3 *FloatValue `json:"o3"`
	// Amount of nitrogen dioxide in the air for this weather sample.
	NO2 *FloatValue `json:"no2"`
	// Amount of carbon monoxide in the air for this weather sample.
	CO *FloatValue `json:"co"`
	// Amount of sulfur dioxide in the air for this weather sample.
	SO2 *FloatValue `json:"so2"`
	// Air quality index for this weather sample per United States
	// Environmental Protection Agency standard.
	EpaAQI *IntValue `json:"epa_aqi"`
	// Primary pollutant in the air for this weather sample per United
	// States Environmental Protection Agency standard.
	EPAPrimaryPollutant *StringValue `json:"epa_primary_pollutant"`
	// Health concern for this weather sample per United States
	// Environmental Protection Agency standard.
	EPAHealthConcern *StringValue `json:"epa_health_concern"`
	// Air quality index for this weather sample per China Ministry of
	// Ecology and Environment standard.
	ChinaAQI *IntValue `json:"china_aqi"`
	// Primary pollutant in the air for this weather sample per China
	// Ministry of Ecology and Environment standard.
	ChinaPrimaryPollutant *StringValue `json:"china_primary_pollutant"`
	// Health concern for this weather sample per China Ministry of Ecology
	// and Environment standard.
	ChinaHealthConcern *StringValue `json:"china_health_concern"`
}

type FireIndexType struct {
	// The level of risk of fires for this weather sample, from a scale of
	// 1-100, based on conditions that play a major role in fires.
	FireIndex *FloatValue `json:"fire_index"`
}

type RoadRiskType struct {
	// The road condition for this weather sample, only available for
	// weather samples in US locations. Possible values include
	// "low_risk", "moderate_risk", "mod_hi_risk", "high_risk", and
	// "extreme_risk".
	RoadRisk *StringValue `json:"road_risk"`
	// ClimaCell road risk (EU-only)
	RoadRiskScore *StringValue `json:"road_risk_score"`
	// An integer between 1 and 100 that is indicative of the level of confidence of road risk prediction (EU-only)
	RoadRiskConfidence *IntValue `json:"road_risk_confidence"`
	// Main weather conditions that are impacting the road risk score (EU-only)
	RoadRiskConditions *StringValue `json:"road_risk_conditions"`
}

type NowCastForecast struct {
	BaseResponseType
	WeatherType
	AirQualityType
	RoadRiskType
	FireIndexType
}

type HourlyForecast struct {
	BaseResponseType
	WeatherType
	AirQualityType
	RoadRiskType
	FireIndexType
}

type RealTime struct {
	BaseResponseType
	WeatherType
	AirQualityType
	RoadRiskType
	FireIndexType
}

type HistoricalClimaCell struct {
	BaseResponseType
	WeatherType
	AirQualityType
	RoadRiskType
	FireIndexType
}

type HistoricalStation struct {
	BaseResponseType
	WeatherType
}

// [TODO] If it can be determined that enum values like moon phase and
// precipitaiton type don't change their deserialization without the version
// number also being bumped up," it would be nice to have enums for these values
// instead of using StringValues.

// StringValue is a field on a Weather returned from the ClimaCell API that is
// of type string.
type StringValue struct {
	// Value indicates the string value for this field on a Weather.
	Value *string `json:"value"`
}

// GetValue returns this  struct's value and a true "ok" if present, or returns
// a blank string and false "ok" if either this StringValue is nil, or its
// Value is nil.
func (s *StringValue) GetValue() (val string, ok bool) {
	if s == nil || s.Value == nil {
		return "", false
	}
	return *s.Value, true
}

// FloatValue is a field on a Weather returned from the ClimaCell API that is a
// floating-point number.
type FloatValue struct {
	// Value indicates the float value for this field on a Weather.
	Value *float64 `json:"value"`
	// Units, if present, indicates the unit of measure for this value.
	Units string `json:"units,omitempty"`
}

// GetValue returns this struct's value and a true "ok" if present, or returns
// 0.0 and false "ok" if either this FloatValue is nil, or its Value is nil.
func (f *FloatValue) GetValue() (val float64, ok bool) {
	if f == nil || f.Value == nil {
		return 0.0, false
	}
	return *f.Value, true
}

// IntValue is a field on a Weather returned from the ClimaCell API that is an
// integer.
type IntValue struct {
	// Value indicates the integer value for this field on a Weather.
	Value *int `json:"value"`
	// Units, if present, indicates the unit of measure for this value.
	Units string `json:"units,omitempty"`
}

// GetValue returns this struct's value and a true "ok" if present, or returns
// 0 and false "ok" if either this IntValue is nil, or its Value is nil.
func (i *IntValue) GetValue() (val int, ok bool) {
	if i == nil || i.Value == nil {
		return 0, false
	}
	return *i.Value, true
}

// TimeValue is a field on a Weather returned from the ClimaCell API that is a
// timestamp.
type TimeValue struct {
	// Value indicates the timestamp value for this field on a Weather.
	Value *time.Time `json:"value"`
	// Units, if present, indicates the unit of measure for this value.
	Units string `json:"units,omitempty"`
}

// GetValue returns this struct's value and a true "ok" if present, or returns
// a blank string and false "ok" if either this TimeValue is nil, or its Value
// is nil.
func (t *TimeValue) GetValue() (val time.Time, ok bool) {
	if t == nil || t.Value == nil {
		return time.Time{}, false
	}
	return *t.Value, true
}

// DateValue is a timestsamp value that can be either in RFC3339 layout, or in
// YYYY-MM-DD layout. Unlike TimeValue, its value should always be non-nil and
// non-zero, as it is used as the timestamps for forecast data samples.
type DateValue struct{ Value time.Time }

// UnmarshalJSON deserializes a DateValue from JSON.
func (d *DateValue) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}

	var jsonDate jsonDateValue
	if err := json.Unmarshal(b, &jsonDate); err != nil {
		return err
	}
	*d = DateValue{Value: time.Time(jsonDate.Value)}
	return nil
}

type jsonDateValue struct {
	Value timeOrDate `json:"value"`
}

type timeOrDate time.Time

// UnmarshalJSON deserializes a timeOrDate from JSON.
func (t *timeOrDate) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}

	var timeStr string
	if err := json.Unmarshal(b, &timeStr); err != nil {
		return err
	}

	// Try parsing as an RFC3339 timestamp
	if tm, err := time.Parse(time.RFC3339, timeStr); err == nil {
		*t = timeOrDate(tm)
		return nil
	}

	// Try parsing as a date
	tm, err := time.Parse("2006-01-02", timeStr)
	if err == nil {
		*t = timeOrDate(tm)
	}
	return err
}

//
// Query parameters for weather data requests
//

// ForecastArgs is converted to query parameters for forecast endpoints.
type ForecastArgs struct {
	// Location, sets the location we are requesting weather data for,
	// which is either a location ID ("location_id" query parameter) or
	// latitude and longitude coordinates ("lat" and "lon" query
	// parameters).
	// A location is the one field that is required for forecast requests;
	// if it is absent and therefore no location query params are filled,
	// any request for forecast data will error with a 400.
	Location Location
	// Start, if nonzero, indicates the start of the time range we are
	// requesting weather data for, filling in the "start_time" query
	// parameter.
	Start time.Time
	// End, if nonzero, indicates the end of the time range we are
	// requesting we
	//ther data for, filling in the "end_time" query
	// parameter.
	End time.Time
	// Timestep, if nonzero, indicates the timestep in minutes for the
	// weather samples we are requesting by filling the "timestep" query
	// parameter. For example if timestep is 5 on the nowcast endpoint, we
	// are requesting nowcast data for every five minutes.
	// Only used on the /weather/historical/climacell and /weather/nowcast
	// endpoints; on other endpoints if this is used, the request will
	// error with a 400.
	Timestep int
	// UnitSystem indicates whether we are requesting weather data in SI or
	// US units of measure, filling in the "unit_system" query parameter.
	// The default is SI.
	UnitSystem string
	// Fields indicates which fields we want on the returned weather
	// sample, such as "temp", "humidity", etx.
	Fields []string
}

// QueryParams converts a ForecastArgs to query parameters to send on a request
// for weather data.
func (args ForecastArgs) QueryParams() url.Values {
	q := make(url.Values)
	if args.Location != nil {
		for k, v := range args.Location.LocationQueryParams() {
			q[k] = v
		}
	}

	if !args.Start.IsZero() {
		q.Add("start_time", args.Start.Format(time.RFC3339))
	}
	if !args.End.IsZero() {
		q.Add("end_time", args.End.Format(time.RFC3339))
	}
	if args.Timestep > 0 {
		q.Add("timestep", strconv.Itoa(args.Timestep))
	}
	if args.UnitSystem != "" {
		q.Add("unit_system", args.UnitSystem)
	}
	if len(args.Fields) > 0 {
		q.Add("fields", strings.Join(args.Fields, ","))
	}
	return q
}

// Location produces the query parameters needed for indicating which
// location to request weather data for.
type Location interface {
	// LocationQueryParams returns the query parameters to be added to a
	// request for the weather data for a location.
	LocationQueryParams() url.Values
}

// LatLon produces location query params from a pair of latitude and longitude
// coordinates.
type LatLon struct {
	// Latitude coordinate
	Lat float64 `json:"lat"`
	// Longitude coordinate
	Lon float64 `json:"lon"`
}

// LocationQueryParams implements the Location interface.
func (l LatLon) LocationQueryParams() url.Values {
	return url.Values{
		"lat": []string{strconv.FormatFloat(l.Lat, 'f', -1, 64)},
		"lon": []string{strconv.FormatFloat(l.Lon, 'f', -1, 64)},
	}
}

// LocationID produces location query params from a location ID.
type LocationID string

// LocationQueryParams implements the Location interface.
func (l LocationID) LocationQueryParams() url.Values {
	return url.Values{"location_id": []string{string(l)}}
}
