package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

var (
	cityDatabase    *geoip2.Reader
	countryDatabase *geoip2.Reader
	asnDatabase     *geoip2.Reader
)

const (
	typeAsn     = "asn"
	typeCity    = "city"
	typeCountry = "country"
)

func main() {
	fmt.Println("Starting server")

	cityDatabase, _ = loadDatabase("City")
	countryDatabase, _ = loadDatabase("Country")
	asnDatabase, _ = loadDatabase("ASN")

	r := mux.NewRouter()
	r.HandleFunc("/{database}/{address}", geoIPHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func geoIPHandler(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("X-Api-Key")

	if authHeader == "" {
		_, _ = fmt.Fprintln(w, "Forbidden")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	apiKeys := strings.Split(os.Getenv("API_KEYS"), ",")
	if contains := contains(apiKeys, authHeader); !contains {
		_, _ = fmt.Fprintln(w, "Forbidden")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	address := net.ParseIP(vars["address"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var addressData interface{}
	var err error

	switch database := vars["database"]; database {
	case typeAsn:
		addressData, err = asnDatabase.ASN(address)
	case typeCity:
		addressData, err = cityDatabase.City(address)
	case typeCountry:
		addressData, err = countryDatabase.Country(address)
	default:
		_, _ = fmt.Fprintln(w, "Unknown database")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		_, _ = fmt.Fprintln(w, "IP data not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ipJSON, err := json.Marshal(addressData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Cannot encode to JSON ", err)
	}
	_, _ = fmt.Fprint(w, string(ipJSON))
}

func loadDatabase(name string) (database *geoip2.Reader, err error) {
	db, err := geoip2.Open("data/GeoLite2-" + name + ".mmdb")
	if err != nil {
		return nil, errors.New("can't work with it")
	}

	return db, nil
}
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
