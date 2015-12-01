SERVER=./src/server.go

all: clean build install

build:
	go build $(SERVER)

install:
	go install $(SERVER)

clean:
	rm server
	rm ./bin/server
