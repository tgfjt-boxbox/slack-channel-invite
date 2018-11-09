GO ?= go
app := scit

build:
	@echo "Yes! now buiding ${app}"
	@$(GO) build -o bin/$(app)
.PHONY: build

install:
	@echo "Yes! Installing ${app} ${GOPATH}/bin/${app}"
	@$(GO) install
.PHONY: install
