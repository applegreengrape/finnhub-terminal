module github.com/applegreengrape/finnhub-terminal

go 1.14

require (
	github.com/applegreengrape/finnhub-terminal/finnhub v0.0.0-00010101000000-000000000000
	github.com/applegreengrape/finnhub-terminal/widgets v0.0.0-00010101000000-000000000000
	//github.com/applegreengrape/finnhub-terminal/finnhub v0.0.0-20201025230228-c86e2992f37f
	//github.com/applegreengrape/finnhub-terminal/widgets v0.0.0-00010101000000-000000000000
	//github.com/applegreengrape/finnhub-terminal/widgets v0.0.0-20201025230228-c86e2992f37f
	github.com/applegreengrape/finnhub-terminal/yahoo v0.0.0-20201025232649-188e4ef22e4d
	github.com/mum4k/termdash v0.12.2
)

replace github.com/applegreengrape/finnhub-terminal/widgets => ./widgets

replace github.com/applegreengrape/finnhub-terminal/finnhub => ./finnhub
