package yahoo

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type TemperatureUnit uint8

const (
	TemperatureUnit_Fahrenheit TemperatureUnit = 'f'
	TemperatureUnit_Celcius    TemperatureUnit = 'c'
)

type WeatherForecast struct {
	Code string
	Date string
	Day  string
	High string
	Low  string
	Text string
}

func (s *Session) GetWeatherForecast(location string, limit uint, u TemperatureUnit) ([]WeatherForecast, error) {
	q := fmt.Sprintf(
		`select item.forecast from weather.forecast where woeid in (select woeid from geo.places(1) where text="%s") and u="%c"`,
		location, u)
	if limit > 1 {
		q += fmt.Sprintf(" limit %d", limit)
	}
	res, err := s.request("GET",
		fmt.Sprintf(
			`https://query.yahooapis.com/v1/public/yql?q=%s&format=json`,
			url.QueryEscape(q)), nil)
	if err != nil {
		return nil, err
	}
	var result struct {
		Query struct {
			Results struct {
				Channel []struct {
					Item struct {
						Forecast struct {
							Code string `json:"code"`
							Date string `json:"date"`
							Day  string `json:"day"`
							High string `json:"high"`
							Low  string `json:"low"`
							Text string `json:"text"`
						} `json:"forecast"`
					} `json:"item"`
				} `json:"channel"`
			} `json:"results"`
		} `json:"query"`
	}
	err = json.NewDecoder(res.Body).Decode(&result)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	wfs := make([]WeatherForecast, 0, len(result.Query.Results.Channel))
	for _, c := range result.Query.Results.Channel {
		f := c.Item.Forecast
		wfs = append(wfs, WeatherForecast{
			Code: f.Code,
			Date: f.Date,
			Day:  f.Day,
			High: f.High,
			Low:  f.Low,
			Text: f.Text,
		})
	}
	return wfs, nil
}
