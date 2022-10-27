# 构建web程序镜像包

FROM centos7_go:v1.0

LABEL maintainer="mrl@qq.com"
RUN mkdir -p /opt/${AppName}
ARG AppName
ADD ${AppName} /opt/${AppName}
RUN chmod +x /opt/${AppName}/${AppName}

ENV ProjectName=${AppName}
ENTRYPOINT /opt/${ProjectName}/${ProjectName}

# RUN go env -w GOPROXY=https://goproxy.cn,direct

