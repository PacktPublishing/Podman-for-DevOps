#!/bin/bash
#
# This script can be used to prepare a WordPress single pod environment that 
# can be used to test the podman generate kube command.
#

# Create the volumes
for vol in wpvol dbvol; do podman volume create $vol; done

# Create the main WordPress pod
podman pod create --name wordpress-pod -p 8080:80

# Create the MySQL container
podman create --pod wordpress-pod --name db -v dbvol:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=myrootpasswd -e MYSQL_DATABASE=wordpress -e MYSQL_USER=wordpress -e MYSQL_PASSWORD=wordpress docker.io/library/mysql

# Create the WordPress container
podman create --pod wordpress-pod --name wordpress -v wpvol:/var/www/html -e WORDPRESS_DB_HOST=127.0.0.1 -e WORDPRESS_DB_USER=wordpress -e WORDPRESS_DB_PASSWORD=wordpress -e WORDPRESS_DB_NAME=wordpress docker.io/library/wordpress

# Start the pod
podman pod start wordpress-pod

