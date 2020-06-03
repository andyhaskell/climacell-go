package climacell

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// Client is the client for sending HTTP requests to ClimaCell's HTTP
// endpoints.
type Client struct {
	// the base URL to send requests to. In regular usage, this is
	// https://api.climacell.co/v3, but in test coverage this can be set to
	// the URL for a mock API server.
	baseURL string

	// the API key we are sending requests with
	apiKey string

	// net/http Client for contacting the ClimaCell API.
	c *http.Client
}

func newDefaultHTTPClient() *http.Client { return &http.Client{Timeout: time.Minute} }

// New takes in a ClimaCell API key and returns a client for the ClimaCell API.
// Note that the default client uses an underlying net/http Client where
// requests time out after a minute without a response. If you want to use a
// different net/http Client, you can instead create your API client using
// NewWithClient.
// WARNING: DO NOT share your API key with anyone; if someone else gains access
// to it, they can make requests to the API under your identity. Because of
// this, it is ill-advised to have the key directly in your source code.
func New(apiKey string) *Client { return NewWithClient(apiKey, newDefaultHTTPClient()) }

// NewWithClient takes in a ClimaCell API key and a net/http Client and returns
// a client for the ClimaCell API.
// WARNING: DO NOT share your API key with anyone; if someone else gains access
// to it, they can make requests to the API under your identity. Because of
// this, it is ill-advised to have the key directly in your source code.
func NewWithClient(apiKey string, c *http.Client) *Client {
	return &Client{
		baseURL: "https://api.climacell.co/v3/",
		apiKey:  apiKey,
		c:       c,
	}
}

//
// Weather endpoints
//

// Nowcast returns minute-by-minute weather predictions on successful requests
// to the /weather/nowcast endpoint, returning a slice of Weather samples on a
// 200 response, or an ErrorResponse on a 400, 401, 403, or 500 error. You are
// able to request nowcast data for up to 6 hours out.
//
// Note that if the error is not due to an eror response, then the error is
// wrapped in a pkg/errors withMessage to indicate its cause, so to work with
// the original error, you need to call pkg/errors.Cause(). These errors are
// things such as errors sending the request to the API, or unexpected errors
// deserializing responses.
func (c *Client) Nowcast(args ForecastArgs) ([]NowCastForecast, error) {
	var w []NowCastForecast
	if err := c.getWeatherSamples("weather/nowcast", args, &w); err != nil {
		return nil, err
	}
	return w, nil
}

// HourlyForecast returns an hourly forecast on successful requests to the
// /weather/forecast/hourly endpoint, returning a slice of Weather samples on a
// 200 response, or an ErrorResponse on a 400, 401, 403, or 500 error. You are
// able to request hourly forecast data up to 96 hours out.
//
// Note that if the error is not due to an eror response, then the error is
// wrapped in a pkg/errors withMessage to indicate its cause, so to work with
// the original error, you need to call pkg/errors.Cause(). These errors are
// things such as errors sending the request to the API, or unexpected errors
// deserializing responses.
func (c *Client) HourlyForecast(args ForecastArgs) ([]HourlyForecast, error) {
	var w []HourlyForecast
	if err := c.getWeatherSamples("weather/forecast/hourly", args, &w); err != nil {
		return nil, err
	}
	return w, nil
}

// DailyForecast returns an hourly forecast on successful requests to the
// /weather/forecast/daily endpoint, returning a slice of Weather samples on a
// 200 response, or an ErrorResponse on a 400, 401, 403, or 500 error. You are
// able to request hourly forecast data up to 15 days out.
//
// Note that if the error is not due to an eror response, then the error is
// wrapped in a pkg/errors withMessage to indicate its cause, so to work with
// the original error, you need to call pkg/errors.Cause(). These errors are
// things such as errors sending the request to the API, or unexpected errors
// deserializing responses.
func (c *Client) DailyForecast(args ForecastArgs) ([]ForecastDay, error) {
	var f []ForecastDay
	if err := c.getWeatherSamples("weather/forecast/daily", args, &f); err != nil {
		return nil, err
	}
	return f, nil
}

// HistoricalStation returns an hourly forecast on successful requests to the
// /weather/historical/station endpoint, returning a slice of Weather samples on a
// 200 response, or an ErrorResponse on a 400, 401, 403, or 500 error. You are
// able to request hourly forecast data up to 15 days out.
//
// Note that if the error is not due to an eror response, then the error is
// wrapped in a pkg/errors withMessage to indicate its cause, so to work with
// the original error, you need to call pkg/errors.Cause(). These errors are
// things such as errors sending the request to the API, or unexpected errors
// deserializing responses.
func (c *Client) HistoricalStation(args ForecastArgs) ([]HistoricalStation, error) {
	var f []HistoricalStation
	if err := c.getWeatherSamples("weather/historical/station", args, &f); err != nil {
		return nil, err
	}
	return f, nil
}

// HistoricalClimaCell returns past ClimaCell weather information on successful
// requests to the /weather/historical/climacell endpoint, or an ErrorResponse
// on a 400, 401, 403, or 500 error. You are able to request data for up to 6
// hours into the past.
//
// Note that if the error is not due to an eror response, then the error is
// wrapped in a pkg/errors withMessage to indicate its cause, so to work with
// the original error, you need to call pkg/errors.Cause(). These errors are
// things such as errors sending the request to the API, or unexpected errors
// deserializing responses.
func (c *Client) HistoricalClimaCell(args ForecastArgs) ([]HistoricalClimaCell, error) {
	var f []HistoricalClimaCell
	if err := c.getWeatherSamples("weather/historical/climacell", args, &f); err != nil {
		return nil, err
	}
	return f, nil
}

// RealTime returns observational data at the present time, down to the minute,
// on successful requests to the /weather/realtime endpoint, or an
// ErrorResponse on a 400, 401, 403, or 500 error.
//
// Note that if the error is not due to an eror response, then the error is
// wrapped in a pkg/errors withMessage to indicate its cause, so to work with
// the original error, you need to call pkg/errors.Cause(). These errors are
// things such as errors sending the request to the API, or unexpected errors
// deserializing responses.
func (c *Client) RealTime(args ForecastArgs) (RealTime, error) {
	var f RealTime
	if err := c.getWeatherSamples("weather/realtime", args, &f); err != nil {
		return RealTime{}, err
	}
	return f, nil
}

func (c *Client) getWeatherSamples(
	endpt string,
	args ForecastArgs,
	expectedResponse interface{},
) error {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return errors.WithMessage(err, "parsing base URL")
	}
	u = u.ResolveReference(&url.URL{Path: endpt})

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return errors.WithMessage(err, "making HTTP request")
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("apikey", c.apiKey)
	req.URL.RawQuery = args.QueryParams().Encode()

	res, err := c.c.Do(req)
	if err != nil {
		return errors.WithMessagef(err, "sending weather data request to %s", endpt)
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		if err := json.NewDecoder(res.Body).Decode(expectedResponse); err != nil {
			return errors.WithMessage(err, "deserializing weather response data")
		}
		return nil
	case 400, 401, 403, 404, 500:
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return errors.WithMessage(err, "deserializing weather error response")
		}

		if res.StatusCode == 401 || res.StatusCode == 403 {
			errRes.StatusCode = res.StatusCode
		}
		return &errRes
	default:
		return fmt.Errorf("unexpected HTTP response status code: %d", res.StatusCode)
	}
}

// ErrorResponse returns errors for 400, 401, 403, and 500 errors.
type ErrorResponse struct {
	// StatusCode indicates the HTTP status for this errored API request.
	// For 401 and 403 errors, this is not present in the actual API
	// response's JSON, so this is filled in for us.
	StatusCode int `json:"statusCode"`
	// ErrorCode is the error code for this request. Not present on 401 and
	// 403 errors.
	ErrorCode string `json:"errorCode"`
	// Message is a description of the error that took place.
	Message string `json:"message"`
}

func (err *ErrorResponse) Error() string {
	if err.ErrorCode == "" {
		return fmt.Sprintf("%d API error: %s", err.StatusCode, err.Message)
	}
	return fmt.Sprintf("%d (%s) API error: %s", err.StatusCode, err.ErrorCode, err.Message)
}
