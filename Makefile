
OS ?= linux
ARCH ?= amd64
DB_TYPE :=
DB_NAME :=
SCHEMA_NAME :=
CREDENTIALS :=


.PHONY: modtidy
modtidy:
	go mod tidy

.PHONY: test
test: 
test: modtidy 
	echo $(CREDENTIALS)
	env DB_TYPE=$(DB_TYPE) DB_NAME=$(DB_NAME) SCHEMA_NAME=$(SCHEMA_NAME) CREDENTIALS=$(CREDENTIALS)  GOOS=$(OS) GOARCH=$(ARCH) go test -v ./pkg/test/...
