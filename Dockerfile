FROM golang:1.14.6-stretch

ENV LD_LIBRARY_PATH="$LD_LIBRARY_PATH:/usr/lib"
COPY inc/* /usr/include/
COPY driver/libtaos.so.1.6.6.1 /usr/lib/libtaos.so.1

RUN sed -i "s/deb.debian.org/mirrors.aliyun.com/g" /etc/apt/sources.list && \
    sed -i "s/security.debian.org/mirrors.aliyun.com/g" /etc/apt/sources.list && \
    apt-get update -y && \
    apt-get -y install gcc build-essential && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    echo "y" |cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    mkdir -p /etc/taos && \
    ln -s /usr/lib/libtaos.so.1 /usr/lib/libtaos.so

COPY taos.cfg /etc/taos

WORKDIR /app


