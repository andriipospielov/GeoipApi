package repository

type Registry interface {
	Instances() map[string]GeoipRepository
}

type IpReposRegistry map[string]GeoipRepository

func (i IpReposRegistry) Instances() map[string]GeoipRepository {
	return i
}

func NewRegistry() (Registry, error) {

	return &IpReposRegistry{
		"asn":     NewAsnRepository(),
		"country": NewCountryRepository(),
		"city":    NewCityRepository(),
	}, nil
}
