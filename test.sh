#!/bin/bash

failed=0
root=$(pwd)

for pkg in ../gorf24 ./cmd ./reg ./reg/addr ./reg/addrw ./reg/autoack ./reg/config ./reg/dynpd ./reg/enrxaddr ./reg/feature ./reg/fifo ./reg/retrans ./reg/rfchan ./reg/rfsetup ./reg/rpd ./reg/rxaddr ./reg/rxpw ./reg/status ./reg/txobs ./reg/txaddr ./reg/xaddr  ./pipe ./util ./spi ./gpio

do
	echo "***** TESTING PACKAGE: '$pkg' *****"
	cd $pkg
	go test -v
	# get exit code from 'go test' IF FAILED
	exitcode=$?
	if [ "${exitcode}" -ne "0" ] ; then
	    failed=$(($failed+$exitcode))
	fi

	echo ""
	cd $root
done

echo "************************"
echo "  $failed  TESTS FAILED"
echo "************************"

exit $failed
