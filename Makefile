download:
	$(call download_database,GeoLite2-City)
	$(call download_database,GeoLite2-Country)
	$(call download_database,GeoLite2-ASN)

define download_database


#	wget "https://download.maxmind.com/app/geoip_download_by_token?edition_id=${1}&date=20201215&suffix=tar.gz&token=v2.local.nMUfc1e_aoJl5IV9yoVeJLhXPTdCXuC9mRzFvQL05lkQIxsjCA8_x-Fs30Ru3vXikJBgn8ygvujKRcvydhbjOYKP65dW1-P7H85gJr8PYxEDl7j6jkJ4hM92tIYBirP9aeP_kFqVg4Kr4_FTgrCOAmop_t9mq_IOdVztsuntf1X0NI-AqYbFfu_LwBVE2T3GoLyhDw" -O ./var/${1}.tar.gz
	wget "https://download.maxmind.com/app/geoip_download?edition_id=${1}&license_key=1bmmVwFRYr5t8W6Z&suffix=tar.gz" -O ./var/${1}.tar.gz

	ls
	echo "\n"
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
