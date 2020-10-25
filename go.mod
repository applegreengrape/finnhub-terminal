module github.com/applegreengrape/finnhub-terminal

go 1.14

require (
	github.com/applegreengrape/finnhub-terminal/widgets v0.0.0-00010101000000-000000000000
	//github.com/applegreengrape/finnhub-terminal/finnhub v0.0.0-20201024162940-1043fb189c47 // indirect
	// github.com/applegreengrape/finnhub-terminal/widgets v0.0.0-20201023173905-bc79b24e54bd
	github.com/applegreengrape/finnhub-terminal/yahoo v0.0.0-20201022162204-567869bbc598
	github.com/mum4k/termdash v0.12.2
)

replace github.com/applegreengrape/finnhub-terminal/finnhub => ./finnhub

replace github.com/applegreengrape/finnhub-terminal/widgets => ./widgets
