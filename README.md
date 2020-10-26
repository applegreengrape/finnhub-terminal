# finnhub-terminal

![example view](./_img/pic1.png)

config.json:
```
{
    "apiKey": "",
    "stocks":[
        "FSLY",
        "NET", 
        "FFIV"
    ]
}
```

```
brew tap applegreengrape/finnhub-terminal https://github.com/applegreengrape/finnhub-terminal
brew install finnhub-terminal

config_path=./config.json fterm
```