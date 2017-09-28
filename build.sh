#!/bin/sh
export GOPATH=~/gocode:`pwd`
export PATH=~/gocode/bin:$PATH
gopherjs build main.go
