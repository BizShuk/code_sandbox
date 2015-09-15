# base image
FROM ubuntu:14.04
MAINTAINER bizshuk <biz.shuk@gmail.com>


# basic setup
RUN     apt-get update && apt-get upgrade -y
RUN     apt-get install -y build-essential automake
RUN     apt-get install -y wget curl vim git

# env setup
RUN     apt-get install -y libssl-dev libpq-dev


# add user first
RUN     useradd -m -s /bin/bash git

USER git
WORKDIR /home/git



# common setting
RUN     mkdir lib
RUN     git clone http://192.168.2.52:3000/shuk/openresty_setting.git


# go setup
RUN     cp -r openresty_setting/external_lib/go lib/
RUN     echo 'export GOROOT=$HOME/lib/go' >>  .bashrc
RUN     echo 'export GOPATH=$HOME/go' >> .bashrc
RUN     echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> .bashrc
RUN     /bin/bash -c "source .bashrc"



# gogs
RUN     cp -r openresty_setting/external_lib/gogs gogs



CMD ["gogs/gogs","web"]

EXPOSE 22 3000
