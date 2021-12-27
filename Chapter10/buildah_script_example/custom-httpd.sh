#!/bin/bash

set -euo pipefail

# Trying to pull a non-existing tag of Fedora official image 
container=$(buildah from docker.io/library/fedora:non-existing-tag)
buildah run $container -- dnf install -y httpd; dnf clean all -y
buildah config --cmd "httpd -DFOREGROUND" $container
buildah config --port 80 $container
buildah commit $container myhttpd
buildah tag custom-httpd registry.example.com/custom-httpd:v0.0.1
