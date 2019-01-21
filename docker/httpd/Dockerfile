FROM httpd:2.4.37

RUN mkdir -p /etc/letsencrypt/live/www.osaketen.com/

ADD ./docker/httpd/server.crt /etc/letsencrypt/live/www.osaketen.com/cert.pem
ADD ./docker/httpd/server.key /etc/letsencrypt/live/www.osaketen.com/privkey.pem