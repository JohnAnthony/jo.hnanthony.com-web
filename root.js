var compress = require('compression');
var express = require('express');
var path = require('path');
var vhost = require('vhost');

var ja = express();
ja.use(express.static(path.join(__dirname, "jo.hnanthony.com")));
ja.use(compress());
ja.use(function(req, res) {
	res.status(404).send('404: Page not Found');
});
ja.use(function(error, req, res, next) {
	res.status(500).send('500: Internal Server Error');
});

var ks = express();
ks.use(express.static(path.join(__dirname, "kieransanderson.xyz")));
ks.use(compress());
ks.use(function(req, res) {
	res.status(404).send('404: Page not Found');
});
ks.use(function(error, req, res, next) {
	res.status(500).send('500: Internal Server Error');
});

var app = express();
app.use(vhost("jo.hnanthony.com", ja));
app.use(vhost("www.jo.hnanthony.com", ja));
app.use(vhost("kieransanderson.xyz", ks));
app.use(vhost("www.kieransanderson.xyz", ks));
app.listen(8080);
