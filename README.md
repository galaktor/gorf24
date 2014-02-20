# Radio-fy your Raspberry Pi
It's software that allows you to control the low-cost Nordic Semiconductor nRF24L01+ radio transceiver.

It works or can be modified to make it work on any system that satisfies the following conditions:
* Linux OS
* SPI
* Go programming language

I developed it on/for the Raspberry Pi running Arch Linux for ARM. The Pi has GPIO pins with SPI. If your device needs to control the pins and/or SPI differently, the relevant code is simple and not very hard to change. The majority of gorf24 code deals with the transceiver logic according to the official specification and will work if you can make the gpio and spi code work for you.

## HOW TO INSTALL
```
$> go get github.com/galaktor/gorf24
```

TODO: setup details with Arch on Pi

# COPYRIGHT AND LICENSE

Copyright 2013, 2014 Raphael Estrada

Licensed under the [GNU General Public License verison 3](http://www.gnu.org/licenses/gpl-3.0.txt "GNU GPL v3")

```
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
```


# THE GIANTS' SHOULDERS
I wrote the current version of gorf24 entirely from
scratch, mostly based on the following sources:

* spidev documentation:    <https://www.kernel.org/doc/Documentation/spi/spidev>
* RPi gpio tutorial:       <https://sites.google.com/site/semilleroadt/raspberry-pi-tutorials/gpio>
* nRF24L01+ specification: <http://www.nordicsemi.com/eng/Products/2.4GHz-RF/nRF24L01P>

Early versions of gorf24 dynamically liked to the RF24
library for Raspberry Pi by Stanley Seow. Some of the
comments in the code pointed me in the right direction,
most notably with regards to timing.
<https://github.com/stanleyseow/RF24>

Seow's work is stronly derived from maniacbug's original RF24 library.
Much kudos to maniacbug for the great work.
https://github.com/maniacbug/RF24
http://maniacbug.wordpress.com/


