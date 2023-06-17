FROM busybox

EXPOSE 3306

COPY ./ls-kh-resolvedata /
COPY ./conf.yaml /etc/

ENTRYPOINT ["/ls-kh-resolvedata"]