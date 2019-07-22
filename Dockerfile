FROM nginx:alpine
MAINTAINER John Anthony
COPY bg.png /usr/share/nginx/html
COPY index.html /usr/share/nginx/html
