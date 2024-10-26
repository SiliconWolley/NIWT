package iface

import "time"

type WeatherCode int

const (
	CodeUnknown WeatherCode = iota
	CodeClear
	CodeFewClouds
	CodeCloudy
	CodeRainShower
	CodeRain
	CodeThunderStorm
	CodeSnow
	CodeMist
)

type Cond struct {
	// Time is the time, where this weather condition applies.
	Time time.Time

	// Code is the general weather condition and must be one the WeatherCode
	// constants.
	Code WeatherCode

	// Desc is a short string describing the condition. It should be just one
	// sentence.
	Desc string

	// TempC is the temperature in degrees celsius.
	TempC *float32

	// FeelsLikeC is the felt temperature (with windchill effect e.g.) in
	// degrees celsius.
	FeelsLikeC *float32

	// ChanceOfRainPercent is the probability of rain or snow. It must be in the
	// range [0, 100].
	ChanceOfRainPercent *int

	// PrecipM is the precipitation amount in meters(!) per hour. Must be >= 0.
	PrecipM *float32

	// VisibleDistM is the visibility range in meters(!). It must be >= 0.
	VisibleDistM *float32

	// WindspeedKmph is the average wind speed in kilometers per hour. The value
	// must be >= 0.
	WindspeedKmph *float32

	// WindGustKmph is the maximum temporary wind speed in kilometers per
	// second. It should be > WindspeedKmph.
	WindGustKmph *float32

	// WinddirDegree is the direction the wind is blowing from on a clock
	// oriented circle with 360 degrees. 0 means the wind is blowing from north,
	// 90 means the wind is blowing from east, 180 means the wind is blowing
	// from south and 270 means the wind is blowing from west. The value must be
	// in the range [0, 359].
	WinddirDegree *int

	// Humidity is the *relative* humidity and must be in [0, 100].
	Humidity *int
}

type WeatherResponse struct {
	Cod  string `json:"cod"`
	City struct {
		Name     string `json:"name"`
		Country  string `json:"country"`
		TimeZone int64  `json: "timezone"`
		// sunrise/sunset are once per call
		SunRise int64 `json: "sunrise"`
		SunSet  int64 `json: "sunset"`
	} `json:"city"`
	List []WeatherSpan `json:"list"`
}

type WeatherSpan struct {
	Dt         int64 `json:"dt"`
	Visibility int   `json:"visibility"`
	Main       struct {
		TempC      float32 `json:"temp"`
		FeelsLikeC float32 `json:"feels_like"`
		Humidity   int     `json:"humidity"`
	} `json:"main"`

	Weather []struct {
		Description string `json:"description"`
		ID          int    `json:"id"`
	} `json:"weather"`

	Wind struct {
		Speed float32 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float32 `json:"gust"`
	} `json:"wind"`

	Rain struct {
		MM3h float32 `json:"3h"`
	} `json:"rain"`
}
