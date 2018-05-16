all: | sanity install run

sanity:
	@dep ensure
	@go list ./... | grep -v vendor/ | xargs go vet 
	@go list ./... | grep -v vendor/ | xargs golint
	@go list ./... | grep -v vendor/ | xargs go fmt
	@go list ./... | grep -v vendor/ | xargs go test --cover 
.PHONY: sanity

install:
	@go install -v github.com/KyleBanks/hows-the-market
.PHONY: install

run: | install
	@hows-the-market
.PHONY: run

