var http = require('http');

var window = {"loc":{}}; 
var fs = require('fs');
var vm = require('vm');
var includeInThisContext = function(path) {
    var code = fs.readFileSync(path);
    vm.runInThisContext(code, path);
}.bind(this);
includeInThisContext('channel.js');


var options = {
  hostname: 'vidao-channel-api.appspot.com',
  port: 80,
  path:'/createChannel?user=test',
  method: 'GET'
};

var req = http.request(options, function(res) {
  res.setEncoding('utf8');
  res.on('data', function (chunk) {
      console.log("Channel Created:",chunk);
      channel = new goog.appengine.Channel(JSON.parse(chunk));
      console.log(channel);
      channel.open({
          onopen : function () {
              console.log("===============================================");
              console.log("Connected to [" + chunk + "]");
          },
          onmessage : function (msg) {

          },
          onerror : function (err) {
              console.log("error (" + err.code + ": " + err.description);
          },
          onclose : function () {
              console.log("===============================================");
              console.log("Disconnected");
          }
      });
  });
});

req.on('error', function(e) {
  console.log('problem with request: ' + e.message);
});

req.end();