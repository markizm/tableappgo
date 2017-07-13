#Using same OS as my ssvm for my image inside the container
FROM centos:centos7
MAINTAINER Mark Magaling mm343g@yp.com
COPY . /
COPY ./run /
RUN chmod -v +x /run
WORKDIR /

EXPOSE 9011

CMD ["./run"]
