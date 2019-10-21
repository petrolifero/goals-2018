#!/bin/bash


export GOPATH=`pwd`
cd "$GOPATH/src/mapreduce"
go test -run Sequential mapreduce/ ...
