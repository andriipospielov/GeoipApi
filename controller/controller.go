package controller

import (
	"encoding/json"
	"fmt"
	"github.com/andriipospielov/GeoipApi/repository"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

type Controller interface{}

type CrudController interface {
	Controller
	Index(w http.ResponseWriter, r *http.Request)
}

type IpController struct {
	repositoryRegistry repository.Registry
	cache              *repository.CachedResults
}

func NewCrudController() IpController {
	registry, _ := repository.NewRegistry()
	return IpController{repositoryRegistry: registry, cache: repository.NewCachedResults()}
}

func (c *IpController) Index(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	source := vars["source"]

	geoipRepository, ok := c.repositoryRegistry.Instances()[source]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "Unknown sourcee: %s\n", source)
		return
	}

	responseData, hit := c.cache.Value(source, net.ParseIP(vars["needle"]))
	var err error = nil
	if !hit {
		responseData, err = geoipRepository.FindByIp(net.ParseIP(vars["needle"]))
	}

	if !hit && err == nil {
		c.cache.Set(source, net.ParseIP(vars["needle"]), responseData)
	}

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintln(w, "IP data not found")
		return
	}

	var marshalledResponse []byte
	marshalledResponse, err = json.Marshal(responseData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Cannot encode to JSON ", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, _ = fmt.Fprint(w, string(marshalledResponse))
}
