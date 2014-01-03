#!/bin/bash

#  Copyright 2013, Raphael Estrada
#  Author email:  <galaktor@gmx.de>
#  Project home:  <https://github.com/galaktor/gorf24>
#  Licensed under The MIT License (see README and LICENSE files)

rfhdrdir=../RF24/librf24-rpi/librf24
rflibdir=$rfhdrdir

g++ -g -O2 -fPIC -I$rfhdrdir -I. -L$rflibdir  -lrf24 -shared -o librf24_c.so *.cpp -ansi -pedantic

cp librf24_c.so /usr/local/lib
