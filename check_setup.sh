#!/bin/bash

# Check if Go is installed
if ! command -v go &> /dev/null
then
    echo "go is not installed. Please install go first. https://go.dev/doc/install"
    exit 1
else
    echo "go is installed: $(go version)"
fi

# Check if make is installed
if ! command -v make &> /dev/null
then
    echo "make is not installed. Please install make first."
    exit 1
else
    echo "make is installed: $(make --version | head -n1)"
fi

# Check if swag is installed
if ! command -v swag &> /dev/null
then
    echo "swag is not installed. Please install swag first. github.com/swaggo/swag"
    exit 1
else
    echo "swag is installed: $(swag --version)"
fi
