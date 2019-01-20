FROM httpd:2.4.37

RUN mkdir -p /etc/letsencrypt/live/www.osaketen.com/

ADD ./docker/httpd/cert.pem /etc/letsencrypt/live/www.osaketen.com/cert.pem
# ファイルは各自で用意／所有者・パーミッションを確認
ADD ./docker/httpd/privkey.pem /etc/letsencrypt/live/www.osaketen.com/privkey.pem