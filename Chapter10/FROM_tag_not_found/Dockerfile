FROM docker.io/library/fedora:sometag

# Update image and install httpd
RUN dnf install -y httpd && dnf clean all -y

# Expose the default httpd port 80
EXPOSE 80

# Run the httpd
CMD ["/usr/sbin/httpd", "-DFOREGROUND"]

