# 构建基础镜像，可直接编译go程序

FROM 192.168.2.4/devel/centos7:v1.0

LABEL maintainer="mrl@qq.com"

ADD docker-18.06.0-ce.tgz /opt

WORKDIR /opt/docker
# RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN ln -s /opt/docker/* /usr/bin/
