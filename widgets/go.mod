module github.com/applegreengrape/finnhub-terminal/widgets

go 1.14

require (
	github.com/applegreengrape/finnhub-terminal/finnhub v0.0.0-00010101000000-000000000000
	//github.com/applegreengrape/finnhub-terminal/finnhub v0.0.0-20201022224320-74d3c93a3a14
	github.com/applegreengrape/finnhub-terminal/yahoo v0.0.0-20201022162204-567869bbc598
	github.com/mum4k/termdash v0.12.2
)

replace github.com/applegreengrape/finnhub-terminal/finnhub => ../finnhub

replace github.com/applegreengrape/finnhub-terminal/yahoo => ../yahoo
