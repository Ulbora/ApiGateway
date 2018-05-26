FROM ubuntu:16.04

#RUN sudo apt-get update
RUN apt-get update  
RUN apt-get install -y ca-certificates
ADD main /main
ADD entrypoint.sh /entrypoint.sh
WORKDIR /

EXPOSE 3020
ENTRYPOINT ["/entrypoint.sh"]

