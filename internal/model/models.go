package model

type Location struct {
	Prefecture string   `json:"prefecture"`
	Subareas   []string `json:"subareas"`
}

type LocationReport struct {
	Intensity string     `json:"intensity"`
	Locations []Location `json:"locations"`
}

type JapanEarthquakeDetail struct {
	EarthquakeTime  string           `json:"earthquakeTime"` // 这个是存储原始时间，用于在查询时去重的
	OccurTime       string           `json:"occurTime"`      // 这个存的是格式化后的ISO标准时间
	Center          string           `json:"center"`
	MaxIntensity    string           `json:"maxIntensity"`
	Magnitude       float64          `json:"magnitude"`
	Depth           string           `json:"depth"`
	Latitude        string           `json:"latitude"`
	Longitude       string           `json:"longitude"`
	Info            string           `json:"info"`
	LocationReports []LocationReport `json:"locationReports"`
}
