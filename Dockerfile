# 2020/10/14最新versionを取得
FROM golang:1.20.5-alpine
# アップデートとgitのインストール！！
RUN apk update && apk add git
WORKDIR /usr/src/app
COPY ./src /usr/src/app/
COPY init.sh /usr/src/app/
RUN go mod download
EXPOSE 9000
CMD ["sh", "init.sh"]