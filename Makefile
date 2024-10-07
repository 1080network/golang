ifndef VERSION
VERSION=main
endif

.PHONY: all
all: clean prepare shared discount partner serviceprovider

.PHONY: shared
shared:
	cd shared && go mod tidy && go test

.PHONY: discount
discount: vendor
	rm -rf discount/proto/*
	cp -r proto/micashared discount/proto
	mkdir -p discount/proto/mica
	cp -r proto/mica/discount discount/proto/mica/.
	./generateGo.sh discount
	cd discount && go mod tidy

.PHONY: partner
partner: vendor
	rm -rf partner/proto/*
	cp -r proto/micashared partner/proto
	mkdir -p partner/proto/mica
	cp -r proto/mica/partner partner/proto/mica/.
	./generateGo.sh partner
	cd partner && go mod tidy

.PHONY: serviceprovider
serviceprovider: vendor
	rm -rf serviceprovider/proto/*
	cp -r proto/micashared serviceprovider/proto
	mkdir -p serviceprovider/proto/mica/discount
	cp -r proto/mica/serviceprovider serviceprovider/proto/mica/.
	cp -r proto/mica/discount/discount  serviceprovider/proto/mica/discount/.
	./generateGo.sh serviceprovider
	cd serviceprovider && go mod tidy

vendor:
	go mod vendor

.PHONY: generate
generate:
	go generate proto.go

prepare: generate vendor

.PHONY: protoupdate
protoupdate: 
	cd proto && git fetch --tags && git checkout $(VERSION)

clean: 
	rm -rf vendor
