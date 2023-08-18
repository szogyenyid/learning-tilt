# Backend
## Main app
docker_build('dummy-php-app-image', 'dummy-php-app', live_update=[
  sync('dummy-php-app/src/', '/var/www/html/'),
])
k8s_yaml('k8s/dummy-php-app.yaml')
k8s_resource('dummy-php-app', port_forwards=8000, labels=["backend"])

# Infra
## MySQL instance
k8s_yaml('k8s/mysql.yaml')
k8s_resource('mysql', port_forwards=3306, labels=["infra"])

## PhpMyAdmin
k8s_yaml('k8s/phpmyadmin.yaml')
k8s_resource('phpmyadmin', port_forwards=8080, labels=["infra"])