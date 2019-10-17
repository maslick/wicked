# =wicked=
simple HTTP wiki

[![image size](https://img.shields.io/badge/image%20size-3MB-blue.svg)](https://hub.docker.com/r/maslick/wicked)

## Features
* written in Go
* saves wiki entries in files
* Docker image size ~3MB

## Installation
```shell script
git clone https://github.com/maslick/wicked.git
cd wicked
go build -ldflags="-s -w" -o wiki
go build -ldflags="-s -w" -o wiki.zip && upx wiki.zip
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

 ```shell script
docker run -d -p 8080:8080 maslick/wicked
open http://`docker-machine ip default`:8080
```

## Links
* [upx](https://upx.github.io/)