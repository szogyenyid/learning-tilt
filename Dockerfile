FROM php:8.2-apache
COPY ./../../src/ /var/www/html/

RUN docker-php-ext-install pdo_mysql
RUN sed -i 's/80/8000/g' /etc/apache2/sites-available/000-default.conf /etc/apache2/ports.conf
EXPOSE 8000