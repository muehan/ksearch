#!/bin/bash
go build -o bin/ksearch src/*.go

echo $OSTYPE
if [[ "$OSTYPE" == "msys" ]]; then
    cp bin/ksearch /c/bin
else
    sudo cp bin/ksearch /usr/local/bin
fi