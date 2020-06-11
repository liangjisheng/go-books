#!/bin/bash

# https://www.linuxidc.com/Linux/2019-06/159059.htm

# 该命令将安装一堆新包，包括gcc，g ++和make
# Ubuntu 16.04存储库中可用的默认GCC版本是5.4.0
# Ubuntu 18.04存储库中可用的默认GCC版本是7.4.0
sudo apt install build-essential

# 安装多个GCC版本
# 添加ubuntu-toolchain-r/test PPA
sudo apt install software-properties-common
sudo add-apt-repository ppa:ubuntu-toolchain-r/test
sudo apt-get update
# sudo apt install gcc-7 g++-7 gcc-8 g++-8 gcc-9 g++-9
sudo apt install gcc-8 g++-8 gcc-9 g++-9

# 将为每个版本配置替代方案并将优先级与其关联。 默认版本是具有最高优先级的版本，在我们的例子中是gcc-9
sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-9 90 --slave /usr/bin/g++ g++ /usr/bin/g++-9
sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-8 80 --slave /usr/bin/g++ g++ /usr/bin/g++-8

# 如果要更改默认版本，请使用update-alternatives命令
sudo update-alternatives --config gcc
