#!/bin/bash

wget https://download.oracle.com/berkeley-db/db-4.8.30.NC.tar.gz
tar zxf db-4.8.30.NC.tar.gz
cd db-4.8.30.NC
cd build_unix
../dist/configure --enable-cxx
make
make install

cd /etc/ld.so.conf.d/
echo "/usr/local/BerkeleyDB.4.8/lib" > BerkeleyDB.conf
ldconfig
