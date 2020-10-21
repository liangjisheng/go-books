var http = require('http');

http.createServer(function(req, res){
    res.writeHead(200, {'Content-Type': 'text/plain'});
    res.end('Hello World\n');
}).listen(8808, '127.0.0.1');

console.log('Server running at http://127.0.0.1:8808');

// 终端 curl 输出 Hello World
// curl http://127.0.0.1:8808/
