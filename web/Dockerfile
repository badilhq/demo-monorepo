FROM nginxinc/nginx-unprivileged:1.30.6-alpine3.24
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY site/ /usr/share/nginx/html/
EXPOSE 8080
