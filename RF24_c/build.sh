#!/bin/bash

rfhdrdir=../RF24/librf24-rpi/librf24
rflibdir=$rfhdrdir

g++ -g -O2 -fPIC -I$rfhdrdir -I. -L$rflibdir  -lrf24 -shared -o librf24_c.so *.cpp -ansi -pedantic

cp librf24_c.so /usr/local/lib
