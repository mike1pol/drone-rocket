FROM alpine
ADD rocket /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/rocket
