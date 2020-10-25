package data

// News ..
type News []struct {
	Category string `json:"category"`
	Datetime int64  `json:"datetime"`
	Headline string `json:"headline"`
	ID       int    `json:"id"`
	Image    string `json:"image"`
	Related  string `json:"related"`
	Source   string `json:"source"`
	Summary  string `json:"summary"`
	URL      string `json:"url"`
}

// CompanyNewsStruct ..
type CompanyNewsStruct []struct {
	Category string `json:"category"`
	Datetime int64  `json:"datetime"`
	Headline string `json:"headline"`
	ID       int    `json:"id"`
	Image    string `json:"image"`
	Related  string `json:"related"`
	Source   string `json:"source"`
	Summary  string `json:"summary"`
	URL      string `json:"url"`
}

// EarningsCalendarStruct ..
type EarningsCalendarStruct struct {
	EarningsCalendar []struct {
		Date            string  `json:"date"`
		EpsActual       float64 `json:"epsActual"`
		EpsEstimate     float64 `json:"epsEstimate"`
		Hour            string  `json:"hour"`
		Quarter         int     `json:"quarter"`
		RevenueActual   float64 `json:"revenueActual"`
		RevenueEstimate float64 `json:"revenueEstimate"`
		Symbol          string  `json:"symbol"`
		Year            int     `json:"year"`
	} `json:"earningsCalendar"`
}
