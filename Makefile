download:
	$(call download_database,GeoLite2-City)
	$(call download_database,GeoLite2-Country)
	$(call download_database,GeoLite2-ASN)

define download_database
	wget "https://download.maxmind.com/app/geoip_download?edition_id=${1}&license_key=${LICENSE_KEY}&suffix=tar.gz" -O ./var/${1}.tar.gz
	tar -xzf "./var/${1}.tar.gz" -C data --wildcards "*.mmdb" --strip-components 1
	rm "./var/${1}.tar.gz"
endef

install:
	go get golang.org/x/lint/golint
	go mod download

cs:
	gofmt -s -l .
	go vet ./...
	golint -set_exit_status $(go list ./...)
