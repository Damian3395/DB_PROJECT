#!/bin/bash
echo Starting NEO4J Service
gnome-terminal -x ./db/NEO4J/bin/neo4j console &

echo Setting Up Proxy Service Environment
gvm use go1.4.1
export GOPATH=/home/damian/Documents/DB_Project/platform
export GOBIN=$GOPATH/bin
echo Checking Environment
echo GOPATH: $GOPATH
echo GOBIN: $GOBIN
go version

echo Compiling Proxy Service
cd ./platform
make build
make install

echo Starting Proxy Service
cd ./bin
sleep 10s
gnome-terminal -x ./server &
