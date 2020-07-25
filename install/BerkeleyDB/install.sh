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


vim ~/.bashrc
C_INCLUDE_PATH=$C_INCLUDE_PATH:/usr/local/BerkeleyDB.5.1/include
export C_INCLUDE_PATH
CPLUS_INCLUDE_PATH=$CPLUS_INCLUDE_PATH:/usr/local/BerkeleyDB.5.1/include
export CPLUS_INCLUDE_PATH
source .bashrc

# apt 命令安装
sudo apt-get install libdb++-dev
sudo apt-get remove libdb++-dev
sudo apt-get install libdb5.1-dev
sudo apt-get install libdb5.1++-dev
sudo apt-get install libdb5.3++ libdb5.3++-dev
