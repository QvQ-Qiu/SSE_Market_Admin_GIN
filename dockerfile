FROM golang

WORKDIR "/app"

COPY . .
# Debian换源
RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list.d/debian.sources
# 安装MySQL客户端   实际上只是为了使用mysqldump备份
RUN apt-get update && apt-get install -y default-mysql-client
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN go build -o serve

CMD [ "./serve" ]