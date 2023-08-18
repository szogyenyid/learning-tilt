# Backend
## Main app
docker_build('php-with-tilt-image', '.', live_update=[
  sync('src/', '/var/www/html/'),
])
k8s_yaml('k8s/php-with-tilt.yaml')
k8s_resource('php-with-tilt', port_forwards=8000, labels=["backend"])

# Infra
## MySQL instance
k8s_yaml('k8s/mysql.yaml')
k8s_resource('mysql-deployment', port_forwards=3306, labels=["infra"])

## PhpMyAdmin
k8s_yaml('k8s/phpmyadmin.yaml')
k8s_resource('phpmyadmin-deployment', port_forwards=8080, labels=["infra"])