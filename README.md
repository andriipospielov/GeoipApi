# GeoipApi

# Routes

|     Database    |      Url      |     Method    |
|:---------------:|:-------------:|:---------------:|
|   GeoLite-City  |   /city/{IP}  |   GET  |
| GeoLite-Country | /country/{IP} | GET |
|   GeoLite-ASN   |   /asn/{IP}   |   GET   | 

```bash
user@machine $ curl http://geoipserver/asn/8.8.8.8
{"AutonomousSystemNumber":15169,"AutonomousSystemOrganization":"Google LLC"}
```

# Installation

Download databases

```bash
make download
```

## Brining up

```bash
docker-compose up -d
```
