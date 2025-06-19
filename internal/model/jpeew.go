package model

type RawJapanEEWData struct {
	Result struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		IsAuth  bool   `json:"is_auth"`
	} `json:"result"`
	ReportTime    string      `json:"report_time"`
	RegionCode    string      `json:"region_code"`
	RequestTime   string      `json:"request_time"`
	RegionName    string      `json:"region_name"`
	Longitude     string      `json:"longitude"`
	IsCancel      interface{} `json:"is_cancel"`
	Depth         string      `json:"depth"`
	Calcintensity string      `json:"calcintensity"`
	IsFinal       interface{} `json:"is_final"`
	IsTraining    interface{} `json:"is_training"`
	Latitude      string      `json:"latitude"`
	OriginTime    string      `json:"origin_time"`
	Security      struct {
		Realm string `json:"realm"`
		Hash  string `json:"hash"`
	} `json:"security"`
	Magunitude      string  `json:"magunitude"`
	ReportNum       string  `json:"report_num"`
	RequestHypoType string  `json:"request_hypo_type"`
	ReportID        string  `json:"report_id"`
	AlertFlg        *string `json:"alertflg"`
}

// 从Raw中提取有用的部分并进行一些处理得到的
type JapanEEWData struct {
	ReportTime    string   `json:"reportTime"`
	RegionName    string   `json:"regionName"`
	OriginTime    string   `json:"originTime"`
	ReportID      string   `json:"reportId"`
	IsCancel      bool     `json:"isCancel"`
	IsFinal       bool     `json:"isFinal"`
	IsTraining    bool     `json:"isTraining"`
	Depth         string   `json:"depth"`
	Longitude     float64  `json:"longitude"`
	Latitude      float64  `json:"latitude"`
	CalcIntensity string   `json:"calcIntensity"`
	Magnitude     *float64 `json:"magnitude"`
	ReportNum     int      `json:"reportNum"`
	AlertFlg      string   `json:"alertFlg"`
}
