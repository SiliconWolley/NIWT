package backends

import (
	"NIWT/iface"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
)

func GetWeatherData(apiKey, city string) (*iface.WeatherResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&appid=%s&units=metric&cnt=8", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weather iface.WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, err
	}
	return &weather, nil
}
