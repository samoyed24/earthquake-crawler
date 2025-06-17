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
	MaxIntensity    *string          `json:"maxIntensity"`
	Magnitude       *float64         `json:"magnitude"`
	Depth           string           `json:"depth"`
	Latitude        string           `json:"latitude"`
	Longitude       string           `json:"longitude"`
	Info            string           `json:"info"`
	LocationReports []LocationReport `json:"locationReports"`
}

type JapanEEWData struct {
	Result struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		IsAuth  bool   `json:"is_auth"`
	} `json:"result"`
	ReportTime    string `json:"report_time"`
	RegionCode    string `json:"region_code"`
	RequestTime   string `json:"request_time"`
	RegionName    string `json:"region_name"`
	Longitude     string `json:"longitude"`
	IsCancel      string `json:"is_cancel"`
	Depth         string `json:"depth"`
	Calcintensity string `json:"calcintensity"`
	IsFinal       string `json:"is_final"`
	IsTraining    string `json:"is_training"`
	Latitude      string `json:"latitude"`
	OriginTime    string `json:"origin_time"`
	Security      struct {
		Realm string `json:"realm"`
		Hash  string `json:"hash"`
	} `json:"security"`
	Magunitude      string `json:"magunitude"`
	ReportNum       string `json:"report_num"`
	RequestHypoType string `json:"request_hypo_type"`
	ReportID        string `json:"report_id"`
}
