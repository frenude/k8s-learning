FROM alpine:latest

MAINTAINER frenude frenude@gmail.com

WORKDIR /root

COPY k8s-test /root/

LABEL app=k8s-test version=v1

EXPOSE 80

CMD ["/bin/sh"]

ENTRYPOINT ["/root/k8s-test"]