#!/bin/bash

cd ..
#GOPATH=..
./util/swagger generate model -f swagger.json -t src
