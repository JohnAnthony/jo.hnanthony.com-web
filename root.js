var express = require('express');
var path = require('path');
var vhost = require('vhost');

var ja = express();
ja.use(express.static(path.join(__dirname, "jo.hnanthony.com")));
ja.use(function(req, res) {
	res.send('404: Page not Found', 404);
});
ja.use(function(error, req, res, next) {
	res.send('500: Internal Server Error', 500);
});

var ks = express();
ks.use(express.static(path.join(__dirname, "kieransanderson.xyz")));
ks.use(function(req, res) {
	res.send('404: Page not Found', 404);
});
ks.use(function(error, req, res, next) {
	res.send('500: Internal Server Error', 500);
});

var app = express();
app.use(vhost("jo.hnanthony.com", ja));
app.use(vhost("www.jo.hnanthony.com", ja));
app.use(vhost("kieransanderson.xyz", ks));
app.use(vhost("www.kieransanderson.xyz", ks));
app.listen(8080);
