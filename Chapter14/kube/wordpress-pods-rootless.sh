#!/bin/bash
#
# This script can be used to prepare a WordPress multi pod environment that 
# can be used to test the podman generate kube command.
#

# Create the rootless network
podman network create kubenet

# Create the volumes
for vol in wpvol dbvol; do podman volume create $vol; done

# Create the MySQL pod and its related container
podman pod create -p 3306:3306 --name mysql-pod --network kubenet
podman create --name db -v dbvol:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=myrootpasswd -e MYSQL_DATABASE=wordpress -e MYSQL_USER=wordpress -e MYSQL_PASSWORD=wordpress --pod mysql-pod docker.io/library/mysql

# Create the WordPress pod and its related container
podman pod create -p 8080:80 --name wordpress-pod --network kubenet
podman create --name wordpress --pod wordpress-pod -v wpvol:/var/www/html -e WORDPRESS_DB_HOST=mysql-pod -e WORDPRESS_DB_USER=wordpress -e WORDPRESS_DB_PASSWORD=wordpress -e WORDPRESS_DB_NAME=wordpress docker.io/library/wordpress

# Start the pods
podman pod start mysql-pod
podman pod start wordpress-pod

