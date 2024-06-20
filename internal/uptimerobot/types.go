package uptimerobot

type Monitor struct {
	ID              int            `json:"id"`
	FriendlyName    string         `json:"friendly_name"`
	Status          int            `json:"status"`
	AvgResponseTime string         `json:"average_response_time"`
	Type            int            `json:"type"`
	URL             string         `json:"url"`
	HttpUsername    string         `json:"http_username"`
	Interval        int            `json:"interval"`
	KeywordType     string         `json:"keyword_type"`
	KeywordValue    string         `json:"keyword_value"`
	Port            string         `json:"port"`
	SubType         string         `json:"sub_type"`
	ResponseTimes   []ResponseTime `json:"response_times"`
}

type ResponseTime struct {
	DateTime int `json:"datetime"`
	Value    int `json:"value"`
}

type AccountDetails struct {
	MonitorLimit   int `json:"monitor_limit"`
	UpMonitors     int `json:"up_monitors"`
	DownMonitors   int `json:"down_monitors"`
	PausedMonitors int `json:"paused_monitors"`
}

type monitorsResponse struct {
	Monitors []Monitor `json:"monitors"`
}

type accountDetailsResponse struct {
	Account AccountDetails `json:"account"`
}
