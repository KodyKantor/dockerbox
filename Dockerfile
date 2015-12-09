FROM ubuntu:15.04

ADD dockerbox.sh /
ADD http://grafvm1.rmnus.sen.symantec.com/hackathon/dockerbox /

RUN chmod 755 /dockerbox
RUN chmod 755 /dockerbox.sh

RUN /dockerbox.sh
ENV PATH /kodybin:$PATH

