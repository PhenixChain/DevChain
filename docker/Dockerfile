FROM alpine
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
ADD DevServer /myApp/
WORKDIR /myApp
ENTRYPOINT ["./DevServer","-f", "data/main.toml"]