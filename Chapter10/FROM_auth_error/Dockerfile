FROM local-registry.example.com/ubi8

# Update image and install httpd
RUN yum install -y httpd && yum clean all -y

# Expose the default httpd port 80
EXPOSE 80

# Run the httpd
CMD ["/usr/sbin/httpd", "-DFOREGROUND"]

