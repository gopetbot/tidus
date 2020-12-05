VERSION         :=      $(shell cat ./VERSION)

help:
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo '    clean              Remove binaries, artifacts and releases.'
	@echo '    test               Run unit tests.'
	@echo '    lint               Run golint.'
	@echo '    imports            Run goimports.'
	@echo '    fmt                Run go fmt.'
	@echo '    env                Display Go environment.'
	@echo '    build              Build project for current platform.'
	@echo '    version            Display Go version.'

tools:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/golang/lint/golint

test:
	go test ./... -v

lint:
	@LINTOUT=`$(GOLINT) $(ALL_PKGS) | grep -v $(TRACE_ID_LINT_EXCEPTION) | grep -v $(TRACE_OPTION_LINT_EXCEPTION) 2>&1`; \
	if [ "$$LINTOUT" ]; then \
		echo "$(GOLINT) FAILED => clean the following lint errors:\n"; \
		echo "$$LINTOUT\n"; \
		exit 1; \
	else \
	    echo "Lint finished successfully"; \
	fi


fmt:
   @FMTOUT=`$(GOFMT) -s -l $(ALL_SRC) 2>&1`; \
   	if [ "$$FMTOUT" ]; then \
   		echo "$(GOFMT) FAILED => gofmt the following files:\n"; \
   		echo "$$FMTOUT\n"; \
   		exit 1; \
   	else \
   	    echo "Fmt finished successfully"; \
   	fi

imports:
	goimports -l -w .

env:
	@go env

build:
    GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
run:
	go run main.go

version:
	@go version