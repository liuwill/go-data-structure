#!/bin/bash
#
# Code coverage generation
rm -rf dist
mkdir dist

go test -covermode=count -coverprofile=./dist/cover.out ./...

# Display the global code coverage
go tool cover -func=./dist/cover.out ;

# If needed, generate HTML report
if [ "$1" == "html" ]; then
  go tool cover -html=./dist/cover.out -o ./dist/coverage.html ;
  cp ./dist/coverage.html ./dist/index.html
fi
