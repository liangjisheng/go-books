#!/bin/bash

# node 官网
# https://nodejs.org/zh-cn/

cd /usr/local
wget https://nodejs.org/dist/v18.12.1/node-v18.12.1-darwin-x64.tar.gz
tar zxvf node-v18.12.1-darwin-x64.tar.gz

# 设置个软连方便全局使用node和npm命令
ln -s /usr/local/node-v12.16.0-linux-x64/bin/node /usr/local/bin/
ln -s /usr/local/node-v12.16.0-linux-x64/bin/npm /usr/local/bin/
ln -s /usr/local/node-v12.16.0-linux-x64/bin/npx /usr/local/bin/

# 查看是否安装成功
node -v
npm -v

# install linux
wget https://nodejs.org/dist/v18.16.0/node-v18.16.0-linux-x64.tar.xz
tar xvJf node-v18.16.0-linux-x64.tar.xz
