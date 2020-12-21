package repository

import (
	"errors"
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

const (
	dataBasePath = "data/GeoLite2-"
)

type Connection struct {
	Reader *geoip2.Reader
}

func NewConnection(databaseName string) *Connection {
	file := dataBasePath + databaseName + ".mmdb"
	db, err := geoip2.Open(file)
	if err != nil {
		log.Fatal(errors.New(fmt.Sprintf("corrupt or missing database file: %s", file)))
	}

	return &Connection{Reader: db}
}

type Repository interface {
	CloseConnection()
}

type GeoipRepository interface {
	Repository
	FindByIp(ip net.IP) (interface{}, error)
}

type AsnRepository struct {
	Connection Connection
}

func NewAsnRepository() *AsnRepository {
	return &AsnRepository{Connection: *NewConnection("ASN")}
}

func (a *AsnRepository) CloseConnection() {
	_ = a.Connection.Reader.Close()
}

func (a *AsnRepository) FindByIp(ip net.IP) (interface{}, error) {
	return a.Connection.Reader.ASN(ip)
}

type CountryRepository struct {
	Connection Connection
}

func NewCountryRepository() *CountryRepository {
	return &CountryRepository{Connection: *NewConnection("Country")}
}

func (a *CountryRepository) CloseConnection() {
	_ = a.Connection.Reader.Close()
}

func (a *CountryRepository) FindByIp(ip net.IP) (interface{}, error) {
	return a.Connection.Reader.Country(ip)
}

type CityRepository struct {
	Connection Connection
}

func NewCityRepository() *CityRepository {
	return &CityRepository{Connection: *NewConnection("City")}
}

func (a *CityRepository) CloseConnection() {
	_ = a.Connection.Reader.Close()
}

func (a *CityRepository) FindByIp(ip net.IP) (interface{}, error) {
	return a.Connection.Reader.City(ip)
}
