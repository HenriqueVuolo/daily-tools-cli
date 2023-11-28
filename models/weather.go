package models

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type Current struct {
	LastUpdatedEpoch int64     `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TempC            float64   `json:"temp_c"`
	TempF            float64   `json:"temp_f"`
	IsDay            int       `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindMph          float64   `json:"wind_mph"`
	WindKph          float64   `json:"wind_kph"`
	WindDegree       int       `json:"wind_degree"`
	WindDir          string    `json:"wind_dir"`
	PressureMb       float64   `json:"pressure_mb"`
	PressureIn       float64   `json:"pressure_in"`
	PrecipMm         float64   `json:"precip_mm"`
	PrecipIn         float64   `json:"precip_in"`
	Humidity         int       `json:"humidity"`
	Cloud            int       `json:"cloud"`
	FeelslikeC       float64   `json:"feelslike_c"`
	FeelslikeF       float64   `json:"feelslike_f"`
	VisKm            float64   `json:"vis_km"`
	VisMiles         float64   `json:"vis_miles"`
	Uv               float64   `json:"uv"`
	GustMph          float64   `json:"gust_mph"`
	GustKph          float64   `json:"gust_kph"`
}

type Day struct {
	MaxTempC          float64   `json:"maxtemp_c"`
	MaxTempF          float64   `json:"maxtemp_f"`
	MinTempC          float64   `json:"mintemp_c"`
	MinTempF          float64   `json:"mintemp_f"`
	AvgTempC          float64   `json:"avgtemp_c"`
	AvgTempF          float64   `json:"avgtemp_f"`
	MaxWindMph        float64   `json:"maxwind_mph"`
	MaxWindKph        float64   `json:"maxwind_kph"`
	TotalPrecipMm     float64   `json:"totalprecip_mm"`
	TotalPrecipIn     float64   `json:"totalprecip_in"`
	TotalSnowCm       float64   `json:"totalsnow_cm"`
	AvgVisKm          float64   `json:"avgvis_km"`
	AvgVisMiles       float64   `json:"avgvis_miles"`
	AvgHumidity       float64   `json:"avghumidity"`
	DailyWillItRain   int       `json:"daily_will_it_rain"`
	DailyChanceOfRain int       `json:"daily_chance_of_rain"`
	DailyWillItSnow   int       `json:"daily_will_it_snow"`
	DailyChanceOfSnow int       `json:"daily_chance_of_snow"`
	Condition         Condition `json:"condition"`
	Uv                float64   `json:"uv"`
}

type Astro struct {
	Sunrise          string `json:"sunrise"`
	Sunset           string `json:"sunset"`
	Moonrise         string `json:"moonrise"`
	Moonset          string `json:"moonset"`
	MoonPhase        string `json:"moon_phase"`
	MoonIllumination int    `json:"moon_illumination"`
	IsMoonUp         int    `json:"is_moon_up"`
	IsSunUp          int    `json:"is_sun_up"`
}

type Hour struct {
	TimeEpoch int64     `json:"time_epoch"`
	Time      string    `json:"time"`
	TempC     float64   `json:"temp_c"`
	TempF     float64   `json:"temp_f"`
	IsDay     int       `json:"is_day"`
	Condition Condition `json:"condition"`
}

type Forecastday struct {
	Date      string `json:"date"`
	DateEpoch int64  `json:"date_epoch"`
	Day       Day    `json:"day"`
	Astro     Astro  `json:"astro"`
	Hour      []Hour `json:"hour"`
}

type Forecast struct {
	Forecastday []Forecastday `json:"forecastday"`
}

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}
