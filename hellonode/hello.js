var http = require('http');
var os = require("os");

process.on('SIGINT', function() {
    console.log("Caught interrupt signal");
    process.exit();
});

http.createServer(function (request, response) {
    response.writeHead(200, {'Content-Type': 'text/plain'});
    response.end(os.hostname() + ' > hello NODE\n');
}).listen(8080);

console.log('Server started');

