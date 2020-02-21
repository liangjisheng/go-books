#!/bin/bash

wget https://nodejs.org/dist/v12.16.0/node-v12.16.0-linux-x64.tar.xz
xz -d node-v12.16.0-linux-x64.tar.xz
tar -xf node-v12.16.0-linux-x64.tar
cd node-v12.16.0-linux-x64

# 设置个软连方便全局使用node和npm命令
ln -s /usr/local/node-v12.16.0-linux-x64/bin/node /usr/local/bin/
ln -s /usr/local/node-v12.16.0-linux-x64/bin/npm /usr/local/bin/
