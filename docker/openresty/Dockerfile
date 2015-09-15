FROM ubuntu:14.04
MAINTAINER bizshuk <biz.shuk@gmail.com>

USER root
WORKDIR /root

# basic setup
RUN     apt-get update && apt-get upgrade -y
RUN     apt-get install -y build-essential libssl-dev automake libpq-dev
RUN     apt-get install -y git curl vim 

# env setup



# openresty setting
RUN     git clone http://192.168.2.52:3000/shuk/openresty_setting.git
RUN     cd openresty_setting; /bin/bash -c ./server_install.sh; /bin/bash -c ./lua_conf_setup.sh;
RUN     chmod +x /root: 



EXPOSE 8000 8888
CMD ["/root/openresty/nginx/sbin/nginx","-g","daemon off;"]



