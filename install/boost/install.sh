#!/bin/bash

# 如果没有的话 安装 boost 1.58.0
wget http://sourceforge.mirrorservice.org/b/bo/boost/boost/1.58.0/boost_1_58_0.tar.gz
tar xf boost_1_58_0.tar.gz
cd boost_1_58_0/
./bootstrap.sh
./b2 -a -sHAVE_ICU=1 # the parameter means that it support icu or unicode 编译时间较长 hpc01机器编译半个多小时
sudo ./b2 install
ldconfig

# 也可以下面的这个命令安装
sudo apt-get install libboost-all-dev
