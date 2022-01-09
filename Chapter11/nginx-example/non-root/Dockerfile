FROM docker.io/library/nginx:mainline-alpine
RUN rm /etc/nginx/conf.d/*
ADD hello-podman.conf /etc/nginx/conf.d/

RUN chmod -R a+w /var/cache/nginx/ \
        && touch /var/run/nginx.pid \
        && chmod a+w /var/run/nginx.pid 

EXPOSE 8080
USER nginx
