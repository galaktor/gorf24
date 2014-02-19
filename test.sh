#!/bin/bash


root=$(pwd)
for pkg in ../gorf24 ./cmd ./reg ./reg/addr ./reg/addrw ./reg/autoack ./reg/config ./reg/dynpd ./reg/enrxaddr ./reg/feature ./reg/fifo ./reg/retrans ./reg/rfchan ./reg/rfsetup ./reg/rpd ./reg/rxaddr ./reg/rxpw ./reg/status ./reg/txobs ./reg/txaddr ./reg/xaddr  ./pipe ./util ./spi ./gpio
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
