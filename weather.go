package climacell

import (
	"time"
)

// Weather contains the data for a single weather sample in a location, and is
// returned from the ClimaCell API's /weather/* endpoints.
type Weather struct {
	// The latitude coordinate for this weather sample.
	Lat float64 `json:"lat"`
	// The longitude coordinate for this weather sample.
	Lon float64 `json:"lon"`
	// The time when this weather sample is from.
	ObservationTime TimeValue `json:"observation_time"`
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
	// The level of risk of fires for this weather sample, from a scale of
	// 1-100, based on conditions that play a major role in fires.
	FireIndex *FloatValue `json:"fire_index"`
	// The road condition for this weather sample, only available for
	// weather samples in US locations. Possible values include
	// "low_risk", "moderate_risk", "mod_hi_risk", "high_risk", and
	// "extreme_risk".
	RoadRisk *StringValue `json:"road_risk"`
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

// StringValue is a field on a Weather returned from the ClimaCell API that is
// of type string.
type StringValue struct {
	// Value indicates the string value for this field on a Weather.
	Value string `json:"value"`
}

// FloatValue is a field on a Weather returned from the ClimaCell API that is a
// floating-point number.
type FloatValue struct {
	// Value indicates the float value for this field on a Weather.
	Value float64 `json:"value"`
	// Units, if present, indicates the unit of measure for this value.
	Units string `json:"units,omitempty"`
}

// IntValue is a field on a Weather returned from the ClimaCell API that is an
// integer.
type IntValue struct {
	// Value indicates the integer value for this field on a Weather.
	Value int `json:"value"`
	// Units, if present, indicates the unit of measure for this value.
	Units string `json:"units,omitempty"`
}

// TimeValue is a field on a Weather returned from the ClimaCell API that is a
// timestamp.
type TimeValue struct {
	// Value indicates the timestamp value for this field on a Weather.
	Value time.Time `json:"value"`
	// Units, if present, indicates the unit of measure for this value.
	Units string `json:"units,omitempty"`
}

// [TODO] If it can be determined that enum values like moon phase and
// precipitaiton type don't change their deserialization without the version
// number also being bumped up, it would be nice to have enums for these values
// instead of using StringValues.
