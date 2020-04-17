package climacell

import (
	"time"
)

// ForecastDay contains the predicted value for a single day in a location's
// daily weather forecast (6AM-6AM timeframe), and is returned from the
// ClimaCell API's /weather/forecast/daily endpoint.
//
// Like on the Weather type, the pointer fields on a ForecastDay will be nil if
// for a daily forecast a particular field was not requested, or if data is not
// available for that day. For convenience, the pointer values on a ForecastDay
// struct have GetValue methods so that callers don't need to worry about
// checking if pointers are nil. For example, to check the value of the high
// temperature for today, the code would look like:
//
// highTemp, ok := f.Temp.Max().GetValue()
// if !ok {
// 	/* handle the case where the max temp value is absent */
// }
// /* work with the retrieved max temp value */
type ForecastDay struct {
	// The latitude coordinate for this weather sample.
	Lat float64 `json:"lat"`
	// The longitude coordinate for this weather sample.
	Lon float64 `json:"lon"`
	// The date when this weather sample is forecast for.
	ObservationTime TimeValue `json:"observation_time"`
	// The temperature for this weather sample.
	Temp *ForecastMinAndMax `json:"temp,omitempty"`
	// The temperature it feels like for this weather sample, based on wind
	// chill and heat window.
	FeelsLike *ForecastMinAndMax `json:"feels_like,omitempty"`
	// The percent relative humidity for this weather sample.
	Humidity *ForecastMinAndMax `json:"humidity,omitempty"`
	// The wind speed for this weather sample.
	WindSpeed *ForecastMinAndMax `json:"wind_speed,omitempty"`
	// The direction of the wind in degrees for this weather sample, where
	// 0 degrees means the wind is going exactly north.
	WindDirection *ForecastMinAndMax `json:"wind_direction,omitempty"`
	// The surface barometric pressure for this weather sample.
	BaroPressure *ForecastMinAndMax `json:"baro_pressure,omitempty"`
	// The amount of precipitation for this weather sample.
	Precipitation *ForecastMinAndMax `json:"precipitation,omitempty"`
	// The precipitation accumulation for this weather sample.
	PrecipitationAccumulation *FloatValue `json:"precipitation_accumulation,omitempty"`
	// The percent probability of precipitation for this day.
	PrecipitationProbability *FloatValue `json:"precipitation_probability,omitempty"`
	// The sunrise time for this location.
	Sunrise *TimeValue `json:"sunrise"`
	// The sunset time for this location.
	Sunset *TimeValue `json:"sunset"`
	// The visibility distance for this weather sample.
	Visibility *ForecastMinAndMax `json:"visibility,omitempty"`
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

// ForecastJSONMinMax is the miniumum or maximum value for a day in a daily
// forecast. It is primarily intended to be used for JSON deserialization, and
// its value can be received from the more convenient ForecastMinAndMax.Min or
// ForecastMinAndMax.Max methods.
type ForecastJSONMinMax struct {
	// The timestamp for this minimum or maximum.
	ObservationTime time.Time `json:"observation_time"`
	// If this is a minumum value, min contains its value and units of
	// measure. If this is present
	Min *FloatValue `json:"min"`
	// If this is a maxumum value, max contains its value and units of
	// measure.
	Max *FloatValue `json:"max"`
}

// FloatAtTimeValue is a field on a ForecastDay returned from the ClimaCell
// API, in which the data are composed of a timestamp and a FloatValue
// representing a minimum or maximum on a forecast for a kind of information
// on the weather, like temperature or precipitation intensity.
type FloatAtTimeValue struct {
	// The timestamp for this minimum or maximum.
	ObservationTime time.Time
	// The value at this observation time.
	Value *FloatValue
}

// [TODO] Figure out if it is possible at all for the API to return a nil
// FloatAtTimeValue.Value, or nil FloatAtTimeValue.Value.Value, for example
// {"observation_time": "2020-04-12T13:49:22.316Z", "value": null}, or
// {
//   "observation_time": "2020-04-12T13:49:22.316Z",
//   "value": {"units": "F", "value": null}
// }.
//
// I doubt that would be the case, though; since a high or low measurement for
// a piece of weather data from a daily forecast includes a timestamp, I don't
// think we would know when the predicted high or low is for the day without
// knowing that high or low value.

// GetValue returns this FloatAtTimeValue's float value and a true "ok" if
// present, or returns 0.0 and false "ok" if either this FloatAtTimeValue is
// nil, its Value field is nil, or its underlying Value is nil.
func (f *FloatAtTimeValue) GetValue() (val float64, ok bool) {
	if f == nil || f.Value == nil || f.Value.Value == nil {
		return 0.0, false
	}
	return *f.Value.Value, true
}

// GetUnits returns this struct's float value and a true "ok" if present, or
// returns an empty string and false "ok" if either this FloatAtTimeValue is
// nil, or its Value is nil.
func (f *FloatAtTimeValue) GetUnits() (units string, ok bool) {
	if f == nil || f.Value == nil {
		return "", false
	}
	return f.Value.Units, true
}

// ForecastMinAndMax contains the minimum and maximum values for a type of
// weather data on a daily forecast, which can be accessed with its Min or Max
// fields.
type ForecastMinAndMax []ForecastJSONMinMax

// Min returns the minimum value for this ForecastMinAndMax. If nil is
// returned, then that means the ForecastMinAndMax did not include a minimum
// value.
func (f ForecastMinAndMax) Min() *FloatAtTimeValue {
	for _, v := range f {
		if v.Min != nil {
			return &FloatAtTimeValue{
				ObservationTime: v.ObservationTime,
				Value:           v.Min,
			}
		}
	}
	return nil
}

// Max returns the maximum value for this ForecastMinAndMax. If nil is
// returned, then that means the ForecastMinAndMax did not include a maximum
// value.
func (f ForecastMinAndMax) Max() *FloatAtTimeValue {
	for _, v := range f {
		if v.Max != nil {
			return &FloatAtTimeValue{
				ObservationTime: v.ObservationTime,
				Value:           v.Max,
			}
		}
	}
	return nil
}
