#! /bin/bash
wget -O go.tar.gz https://dl.google.com/go/go1.13.1.linux-amd64.tar.gz
mkdir ~/bin/go
tar -C ~/bin/go -xzf go.tar.gz
sudo rm -rf /var/lib/dpkg/lock
sudo dpkg --configure -a
sudo apt install python3-pip
/usr/bin/python3 -m pip --no-cache-dir install --upgrade --user cairocffi pip setuptools
/usr/bin/python3 -m pip install --user AwesomeBuild
cd ..
git clone git@github.com:github/hub.git
cd hub
export PATH=~/bin/go/go/bin/:$PATH
make install prefix=~/.local
mkdir ~/assets