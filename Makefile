MAKEFLAGS += --silent

all: 
		install

install: 
		cd producer && go get -d -v
		cd consumer && go get -d -v

clean:
		cd producer && go clean 
		cd consumer && go clean

build:
		cd producer && go build 
		cd consumer && go build

.PHONY: install clean build