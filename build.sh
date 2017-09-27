#!/bin/sh
export GOPATH=~/gocode:`pwd`
export PATH=$GOPATH/bin:$PATH
gopherjs build main.go
