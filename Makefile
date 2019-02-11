.PHONY: ci_dep install run test dep_install bin

ROOT_PATH=$(shell pwd)

PROJECT_PATH=$(ROOT_PATH)/src/github.com/globocom/cerebro

export GOPATH := $(ROOT_PATH)
export PATH := $(GOPATH)/bin/:$(PATH)

test_dep:
	@go get github.com/jstemmer/go-junit-report

ci_dep: test_dep
	@mkdir -p $(GOPATH)/bin/
	@curl -L -s https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 -o $(GOPATH)/bin/dep
	@chmod +x $(GOPATH)/bin/dep


dep_install:
	@cd $(PROJECT_PATH) && dep ensure -add $(DEP)

install: test_dep
	@cd $(PROJECT_PATH) && dep ensure && go install

es:
	@docker-compose -f "docker-compose.yml" up -d --build --scale elasticsearch_data=2

es_logs:
	@docker-compose logs -f

es_down:
	@curl -X DELETE http://localhost:9200/_all -H 'cache-control: no-cache'
	@docker-compose down

run: install es_down es
	./bin/cerebro

test: install
	@cd $(PROJECT_PATH) && go test ./... -coverprofile=$(ROOT_PATH)/coverage.out | go-junit-report > $(ROOT_PATH)/test.xml

install_linux:
	@cd $(PROJECT_PATH) && GOOS=linux GOARCH=amd64 go install

bin: install
	@cp -f $(ROOT_PATH)/bin/linux_amd64/cerebro $(ROOT_PATH)/deploy/
	@chmod a+x $(ROOT_PATH)/deploy/cerebro