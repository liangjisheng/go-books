#!/bin/bash

wget https://releases.hashicorp.com/consul/1.7.1/consul_1.7.1_linux_amd64.zip
mkdir consul
unzip consul_1.7.1_linux_amd64.zip -d consul
cd consul
mkdir data
ln -s /usr/local/consul/consul /usr/local/bin/consul

# 172.16.1.218
# consul agent -server -ui -bootstrap-expect=3 -data-dir=/data/consul -node=agent-1 -client=0.0.0.0 -bind=172.16.1.218 -datacenter=dc1
# 172.16.1.219
# consul agent -server -ui -bootstrap-expect=3 -data-dir=/data/consul -node=agent-2 -client=0.0.0.0 -bind=172.16.1.219 -datacenter=dc1 -join 172.16.1.218
# 172.16.1.220
# consul agent -server -ui -bootstrap-expect=3 -data-dir=/data/consul -node=agent-3 -client=0.0.0.0 -bind=172.16.1.220 -datacenter=dc1 -join 172.16.1.218

consul agent -dev -client=0.0.0.0
