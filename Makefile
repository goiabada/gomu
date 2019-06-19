RM=rm
MV=mv
GO=go
FLAGS=GOOS=linux GOARCH=amd64 CGO_ENABLED=0
NAME=gomu

test:
	$(GO) test --cover -v ./pkg/*

build:
	$(FLAGS) $(GO) build -o ./bin/$(NAME) ./cmd/$(NAME)

run:
	$(FLAGS) $(GO) run ./cmd/$(NAME)

install:
	sudo $(MV) ./bin/$(NAME) /usr/local/bin/

clean:
	$(RM) -rf build/$(NAME)

