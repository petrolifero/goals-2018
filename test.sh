#!/bin/bash

function test1()
{
    export GOPATH=`pwd`
    cd "$GOPATH/src/mapreduce"
    go test -run Sequential mapreduce/ ...
    cd -
}

function test2()
{
    echo try
}

function test3()
{
    echo try
}

function test4()
{
    echo try
}

function test5()
{
    echo try
}




for i in `seq 5`
do
	test$i
done
