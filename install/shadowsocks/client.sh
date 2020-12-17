#!/bin/bash

# https://my.oschina.net/seeseven/blog/2207195
# https://blog.phpgao.com/privoxy-shadowsocks.html

sudo apt install python3-pip
sudo pip install shadowsocks

sudo vim /etc/shadowsocks.json
{
    "server":"52.199.213.200",
    "server_port":10800,
    "local_port":1080,
    "password":"d58lic3IosofCGzf",
    "timeout":600,
    "method":"aes-256-gcm"
}
sudo chmod 755 /etc/shadowsocks.json

# 后台启动
sslocal -c /etc/shadowsocks.json -d start


sudo apt-get install privoxy

vim /etc/privoxy/config

listen-address 127.0.0.1:8118   # 8118 是privoxy端口号
forward-socks5t / 127.0.0.1:1080 . # 1080 是ss的端口号，注意最后的.不要省略了

# 重启 privoxy
/etc/init.d/privoxy restart

vim ~/.bashrc
export http_proxy=http://127.0.0.1:8118   # 和上面privoxy的端口保持一致
export https_proxy=http://127.0.0.1:8118
source ~/.bashrc
