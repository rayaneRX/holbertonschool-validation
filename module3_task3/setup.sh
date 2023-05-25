#!/bin/bash
apt-get update && apt-get install -y make curl npm
curl -LO https://github.com/gohugoio/hugo/releases/download/v0.84.0/hugo_extended_0.84.0_Linux-64bit.tar.gz
tar -xzf hugo_extended_0.84.0_Linux-64bit.tar.gz -C /usr/local/bin/
rm hugo_extended_0.84.0_Linux-64bit.tar.gz
export PATH=/usr/local/bin:$PATH 
make build