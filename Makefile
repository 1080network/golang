ifndef VERSION
    $(error VERSION is not set)
endif

.PHONY: all
all: clean protoupdate prepare shared discount partner serviceprovider fullsdk

.PHONY: shared
shared:
	cd shared && go mod tidy && go test

.PHONY: discount
discount: vendor
	rm -rf discount/proto/*
	cp -r proto/micashared discount/proto
	mkdir -p discount/proto/mica
	cp -r proto/mica/discount discount/proto/mica/.
	cp -r proto/mica/common discount/proto/mica/.
	./generateGo.sh discount
	cd discount && go mod tidy

.PHONY: partner
partner: vendor
	rm -rf partner/proto/*
	cp -r proto/micashared partner/proto
	mkdir -p partner/proto/mica
	cp -r proto/mica/partner partner/proto/mica/.
	cp -r proto/mica/member partner/proto/mica/.
	cp -r proto/mica/common partner/proto/mica/.
	./generateGo.sh partner
	cd partner && go mod tidy

.PHONY: serviceprovider
serviceprovider: vendor
	rm -rf serviceprovider/proto/*
	cp -r proto/micashared serviceprovider/proto
	mkdir -p serviceprovider/proto/mica/discount
	cp -r proto/mica/serviceprovider serviceprovider/proto/mica/.
	cp -r proto/mica/discount/discount  serviceprovider/proto/mica/discount/.
	cp -r proto/mica/member serviceprovider/proto/mica/.
	cp -r proto/mica/common serviceprovider/proto/mica/.
	./generateGo.sh serviceprovider
	cd serviceprovider && go mod tidy

.PHONY: fullsdk
fullsdk: vendor
	rm -rf fullsdk/proto/*
	cp -r proto/micashared fullsdk/proto
	cp -r proto/mica fullsdk/proto
	rm -rf fullsdk/proto/mica/connect
	./generateGo.sh fullsdk
	cd fullsdk && go mod tidy

vendor:
	go mod vendor

.PHONY: generate
generate:
	go generate proto.go

prepare: generate vendor

.PHONY: protoupdate
protoupdate:
	cd proto && git fetch --tags && git checkout $(VERSION)

.PHONY: publish
publish:
	@if [[ -z "${VERSION}" ]]; then echo "Must set VERSION=vX.X.X"; exit 1; fi
	@git diff-index --quiet HEAD -- || (echo "Uncommitted changes found. Please stash or commit them before publishing" && exit 1)
	git tag discount/$(VERSION)
	git tag partner/$(VERSION)
	git tag serviceprovider/$(VERSION)
	git tag fullsdk/$(VERSION)
	git push origin discount/$(VERSION)
	git push origin partner/$(VERSION)
	git push origin serviceprovider/$(VERSION)
	git push origin fullsdk/$(VERSION)

clean: 
	rm -rf vendor
