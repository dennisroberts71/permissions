FROM golang:1.6-alpine

ARG git_commit=unknown
LABEL org.cyverse.git-ref="$git_commit"

COPY . /go/src/github.com/cyverse-de/permissions

RUN apk update

RUN apk add git

RUN git clone https://github.com/swagger-api/swagger-ui.git /tmp/swagger-ui \
    && cd /tmp/swagger-ui \
    && git checkout v2.1.4 \
    && mkdir -p /docs \
    && cp -pr dist/* /docs/ \
    && cd / \
    && rm -rf /tmp/swagger-ui \
    && cp /go/src/github.com/cyverse-de/permissions/index.html /docs/index.html

RUN go get github.com/constabulary/gb/...

RUN cd /go/src/github.com/cyverse-de/permissions && \
    gb build && \
    cp bin/permissions-server /bin/permissions

EXPOSE 60000
ENTRYPOINT ["permissions"]
CMD ["--help"]
