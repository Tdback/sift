BINARY_NAME := sift
SOURCE := ./src/sift.go

all: build

build:
	go build -o ${BINARY_NAME} ${SOURCE}

run:
	go build -o ${BINARY_NAME} ${SOURCE}
	./${BINARY_NAME}

.PHONY: clean
clean:
	go clean
	rm ${BINARY_NAME}
