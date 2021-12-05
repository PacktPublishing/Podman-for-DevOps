var http = require("http");

http.createServer(function(request, response) {
  response.writeHead(200, {"Content-Type": "text/plain"});
  response.write("Hello Podman and Buildah friends. This page is provided to you through a container running Node.js version: ");
  response.write(process.version);
  response.end();
}).listen(8080);

