FROM nginx

RUN rm /etc/nginx/modules/ngx_http_image_filter_module.so \
    /etc/nginx/modules/ngx_http_xslt_filter_module.so \
    /etc/nginx/modules/ngx_http_image_filter_module-debug.so \
    /etc/nginx/modules/ngx_http_xslt_filter_module-debug.so && \
    rm /etc/nginx/fastcgi_params \
    /etc/nginx/uwsgi_params && \
    mkdir -p /etc/nginx/ssl

COPY ./privatekey.key /etc/nginx/ssl/privatekey.key
COPY ./certificate.crt /etc/nginx/ssl/certificate.crt

COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./default.conf /etc/nginx/conf.d/default.conf
COPY ./app.conf /etc/nginx/conf.d/app.conf
EXPOSE 443
CMD ["nginx", "-g", "daemon off;"]
