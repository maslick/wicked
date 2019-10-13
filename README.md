# =wicked=
simple HTTP wiki

## Features
* written in Go
* saves wiki entries in files
* Dockerfile

## Installation
```shell script
git clone https://github.com/maslick/wicked.git
cd wicked
go build -ldflags="-s -w" -o wiki
go build -ldflags="-s -w" -o wiki.zip && upx wiki.zip
./wiki
```

## Usage
```shell script
./wiki
./wiki.zip
```

## Docker
```shell script
GOOS=linux go build -ldflags="-s -w" -o build/gowiki && upx build/gowiki
docker build -t gowiki .
docker run -d -p 8081:8080 gowiki
open http://`docker-machine ip default`:8081
```

## Links
* [upx](https://upx.github.io/)