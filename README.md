# finnhub-terminal

![example view](./_img/pic1.png)

config.json:
```
{
    "apiKey": "",
    "stocks":[
        "BABA",
        "AMZN",
        "GOOGL",
        "MSFT"
    ]
}
```

```
brew tap applegreengrape/finnhub-terminal https://github.com/applegreengrape/finnhub-terminal
brew install finnhub-terminal

config_path=./config.json fterm
```