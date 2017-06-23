# Introduction

Reads CSV from stdin, uses MaxMind GeoIP 2 Country DB to look up IP and write a new CSV to stdout with Country appended.

May add some formatting.

# Install

```
go get -u github.com/bradleyfalzon/addgeoip
```

# Usage

```
Usage of addgeoip:
  -countryDB string
        MaxMind GEOIP2 Country DB path (default "GeoIP2-Country.mmdb")
  -ipfield int
        Field number containing the IP address, starting from 0
```

Example:

```
$ cat blah.csv
timestamp,direction,remote,packets,bytes,duration
01:18:25,outbound,176.215.0.0,3,180,3.33
01:18:27,inbound,101.173.0.0,447,5827000,43.52
01:18:27,inbound,176.215.0.0,4,240,2.62
01:18:27,inbound,52.64.0.0,13,225000,0.13
```

```
$ addgeoip -ipfield 2 < blah.csv
timestamp,direction,remote,packets,bytes,duration
01:18:25,outbound,176.215.0.0,3,180,3.33,RU
01:18:27,inbound,101.173.0.0,447,5827000,43.52,AU
01:18:27,inbound,176.215.0.0,4,240,2.62,RU
01:18:27,inbound,52.64.0.0,13,225000,0.13,AU
```
