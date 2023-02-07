[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/rubiojr/short/master/LICENSE)

## Shawty: URL Shortener Service

This service shortents URLs and stores them in a SQLite database.

It has 3 features: shorten, unshorten, and redirect.

## Deploy

Docker:

```
docker run -e SHORT_BASE_URL=https://change.me.domain -p 8080:8080 -v $PWD:/data ghcr.io/rubiojr/short:latest
```

Docker Compose:

```yaml
version: "3.4"
services:
  short:
    image: "ghcr.io/rubiojr/short:latest"
    restart: unless-stopped
    volumes:
      - "CHANGEME:/data" # sqlite database path
    environment:
      SHORT_BASE_URL: "https://change.me.domain" # the URL prepended to the redirection link
```

## Using

Shorten:

```
❯ curl -X POST -d "url=https://noyb.eu"  localhost:8080
https://short.guru/r/cfh0nilr2gts738en740
```

Your browser will redirect you to [noyb.eu](https://noyb.eu) if you go to that URL.

Decode:

```
✖  curl localhost:8080/d/cfh0nilr2gts738en740
https://noyb.eu
```
