FROM ubuntu:14.04
MAINTAINER bizshuk <biz.shuk@gmail.com>


RUN     apt-get update && apt-get upgrade -y
RUN     apt-get install -y git curl vim automake libssl-dev libpq-dev build-essential


WORKDIR /root

# openresty setting
RUN     git clone http://192.168.2.52:3000/shuk/openresty_setting.git
RUN     cd openresty_setting; /bin/bash -c ./server_install.sh; /bin/bash -c ./lua_profiler_setup.sh;


EXPOSE 9000
CMD ["/root/openresty/nginx/sbin/nginx","-g","daemon off;"]




