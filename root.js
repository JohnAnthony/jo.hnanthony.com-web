var express = require('express');
var path = require('path');
var vhost = require('vhost');

var ja = express();
ja.use(express.static(path.join(__dirname, "jo.hnanthony.com")));

var ks = express();
ks.use(express.static(path.join(__dirname, "kieransanderson.xyz")));

var app = express();
app.use(vhost("jo.hnanthony.com", ja));
app.use(vhost("kieransanderson.xyz", ks));
app.use(vhost("www.kieransanderson.xyz", ks));
app.listen(8080);
