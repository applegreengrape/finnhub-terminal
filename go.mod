module github.com/applegreengrape/finnhub-terminal

go 1.14

require (
	github.com/applegreengrape/finnhub-terminal/finnhub v0.0.0-20201026011118-28077744f0d2 // indirect
	github.com/applegreengrape/finnhub-terminal/widgets v0.0.0-20201026011450-98c96d4e8f02 // indirect
	github.com/applegreengrape/finnhub-terminal/yahoo v0.0.0-20201026011450-98c96d4e8f02
	github.com/mum4k/termdash v0.12.2
)

//replace github.com/applegreengrape/finnhub-terminal/widgets => ./widgets
//replace github.com/applegreengrape/finnhub-terminal/finnhub => ./finnhub
