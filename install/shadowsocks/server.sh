#!/bin/bash

# https://leihungjyu.com/post/ubuntu-install-shadowsocks.html
# ss://Y2hhY2hhMjAtaWV0Zi1wb2x5MTMwNTpkNThsaWMzSW9zb2ZDR3pmQDUyLjE5OS4yMTMuMjAwOjEwODAw

sudo apt install python3-pip
sudo apt-get -y install python-m2crypto
sudo pip3 install https://github.com/shadowsocks/shadowsocks/archive/master.zip
# sudo pip install shadowsocks
ssserver --version

sudo vim /etc/shadowsocks.json
{
    "server":"52.199.213.200",
    "server_port":10800,
    "local_address": "127.0.0.1",
    "local_port":1080,
    "password":"d58lic3IosofCGzf",
    "timeout":300,
    "method":"aes-256-cfb"
}
sudo chmod 755 /etc/shadowsocks.json

# 后台启动
ssserver -c /etc/shadowsocks.json -d start


vim /lib/systemd/system/shadowsocks.service

[Unit]
Description=Shadowsocks Server
After=network.target

[Service]
ExecStart=/usr/local/bin/ssserver -c /etc/shadowsocks.json
Restart=on-abort

[Install]
WantedBy=multi-user.target

systemctl start shadowsocks.service
systemctl enable shadowsocks.service
systemctl status shadowsocks.service