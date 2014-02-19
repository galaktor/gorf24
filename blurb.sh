#!/bin/bash

crsrc="/home/galaktor/mygo/src/github.com/galaktor/gorf24/COPYRIGHT"
for i in **/**/*.go
do
    echo "trying $i"
    if ! grep -q Copyright $i
    then
	echo "blurbing $i"
	cat $crsrc $i > $i.new && mv $i.new $i
    fi
done
