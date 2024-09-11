FROM alpine
ADD ./bin/pod /pod
ENTRYPOINT [ "/pod" ]