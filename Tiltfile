# Frontend
## Frontend
docker_build('frontend-image', 'frontend',
  live_update=[
    # when package.json changes, we need to do a full build
    fall_back_on(['frontend/package.json', 'frontend/package-lock.json']),
    # Map the local source code into the container under /src
    sync('frontend', '/src'),
  ])
k8s_yaml('k8s/frontend.yaml')
k8s_resource('frontend', port_forwards=8888, labels=["frontend"])

# Backend
## Main app
docker_build('dummy-php-app-image', 'dummy-php-app', live_update=[
  sync('dummy-php-app/src/', '/var/www/html/'),
])
k8s_yaml('k8s/dummy-php-app.yaml')
k8s_resource('dummy-php-app', port_forwards=8000, labels=["backend"])

## Reverser service
docker_build('go-reverser-image', 'go-reverser')
k8s_yaml('k8s/go-reverser.yaml')
k8s_resource('go-reverser', port_forwards=8001, labels=["backend"])

# Infra
## MySQL instance
k8s_yaml(kustomize('k8s/mysql'))
k8s_resource('mysql', labels=["infra"])

## PhpMyAdmin
k8s_yaml('k8s/phpmyadmin.yaml')
k8s_resource('phpmyadmin', port_forwards=8080, labels=["infra"])