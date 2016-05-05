FROM node
MAINTAINER John Anthony
RUN mkdir /www
RUN npm install -g coffee-script
CMD ["/usr/local/bin/coffee", "/www/server.coffee"]
EXPOSE 29000

COPY package.json /www/
RUN cd /www && npm install --production

COPY server.coffee /www/
COPY static/ /www/static
