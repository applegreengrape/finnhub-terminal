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

// NewsSentiment .. 
type NewsSentiment struct {
	Buzz struct {
		ArticlesInLastWeek int     `json:"articlesInLastWeek"`
		Buzz               float64 `json:"buzz"`
		WeeklyAverage      int     `json:"weeklyAverage"`
	} `json:"buzz"`
	CompanyNewsScore            float64 `json:"companyNewsScore"`
	SectorAverageBullishPercent float64 `json:"sectorAverageBullishPercent"`
	SectorAverageNewsScore      float64 `json:"sectorAverageNewsScore"`
	Sentiment                   struct {
		BearishPercent float64 `json:"bearishPercent"`
		BullishPercent float64 `json:"bullishPercent"`
	} `json:"sentiment"`
	Symbol string `json:"symbol"`
}

// Trend ..
type Trend []struct {
	Buy        int    `json:"buy"`
	Hold       int    `json:"hold"`
	Period     string `json:"period"`
	Sell       int    `json:"sell"`
	StrongBuy  int    `json:"strongBuy"`
	StrongSell int    `json:"strongSell"`
	Symbol     string `json:"symbol"`
}