#!/bin/bash

sudo apt install curl gnupg2 ca-certificates lsb-release

curl -fsSL https://nginx.org/keys/nginx_signing.key | sudo apt-key add -

sudo apt-key fingerprint ABF5BD827BD9BF62

sudo apt update
sudo apt install nginx

# 查看 nginx 运行状态 监听80端口
sudo systemctl status nginx
sudo systemctl stop nginx
sudo systemctl start nginx
sudo systemctl restart nginx

# 打开防火墙
sudo ufw allow 'Nginx Full'

# 浏览器访问
# http://117.51.148.112:80/
