#!/bin/bash

echo "fetching C++ RF24 for Raspberry Pi"
git clone https://github.com/stanleyseow/RF24

echo "building C++ RF24 for Raspberry Pi"
cd RF24/librf24-rpi/librf24
make
make install
cd ../../..

echo "building C++ RF24 for Raspberry Pi EXAMPLES"
cd RF24/librf24-rpi/librf24/examples
make
cd ../../../..

echo "buliding ANSI C wrapper for C++ RF24 library"
cd RF24_c
bash build.sh
cd ..


