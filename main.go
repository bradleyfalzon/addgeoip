package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	maxminddb "github.com/oschwald/maxminddb-golang"
)

func main() {
	ipfield := flag.Int("ipfield", 0, "Field number containing the IP address, starting from 0")
	countryDB := flag.String("countryDB", "GeoIP2-Country.mmdb", "MaxMind GEOIP2 Country DB path")
	flag.Parse()

	geoIPCountry, err := maxminddb.Open(*countryDB)
	if err != nil {
		log.Fatalf("Cannot open MaxMind GeoIP Country %q: %v", *countryDB, err)
	}

	err = addGeoIP(os.Stdin, os.Stdout, *ipfield, geoIPCountry)
	if err != nil {
		log.Fatal(err)
	}
}

func addGeoIP(reader io.Reader, writer io.Writer, ipfield int, geoIPCountry *maxminddb.Reader) error {
	r := csv.NewReader(reader)
	w := csv.NewWriter(writer)
	defer w.Flush()

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if ipfield < len(record) {
			ip := net.ParseIP(strings.TrimSpace(record[ipfield]))
			if ip != nil {
				country, err := countryFromIP(geoIPCountry, ip)
				if err != nil {
					return err
				}

				// Append country to record
				record = append(record, country)
			}
		}

		if err := w.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func countryFromIP(geoIPCountry *maxminddb.Reader, ip net.IP) (string, error) {
	var geoIP struct {
		Country struct {
			ISOCode string `maxminddb:"iso_code"`
			//Name    string `maxminddb:"name"`
		} `maxminddb:"country"`
	}

	err := geoIPCountry.Lookup(ip, &geoIP)
	if err != nil {
		return "", fmt.Errorf("could not lookup IP %s: %v", ip, err)
	}

	return geoIP.Country.ISOCode, nil
}
