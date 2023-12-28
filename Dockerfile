FROM alpine

RUN apk update  && apk add --no-cache ca-certificates

CMD ["/bin/sh", "-c", "/local/bin/zadanie"]

EXPOSE 8088
EXPOSE 2112

WORKDIR /zadanie

COPY ./build/bin/zadanie /local/bin/

RUN chmod +x /local/bin/zadanie