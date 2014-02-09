#!/bin/bash


root=$(pwd)
for pkg in ../gorf24 ./cmd ./reg ./pipe ./util ./spi ./gpio
do
	echo "***** TESTING PACKAGE: '$pkg' *****"
	cd $pkg
	go test #-v
	echo ""
	cd $root
done

# TODO: any way to join these in one single loop?
#for file in *_test.go; do
#    echo "testing $file"
#    go test $file -v
#done

#for file in **/*_test.go; do
#    echo "testing $file"
#    go test $file -v
#done
