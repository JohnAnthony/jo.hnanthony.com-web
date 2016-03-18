compress = require('compression')
express = require('express')
path = require('path')
vhost = require('vhost')

ja = require('./jo.hnanthony.com/app.coffee')
ks = require('./kieransanderson.xyz/app.coffee')

app = express()
app.use(vhost("jo.hnanthony.com", ja))
app.use(vhost("www.jo.hnanthony.com", ja))
app.use(vhost("kieransanderson.xyz", ks))
app.use(vhost("www.kieransanderson.xyz", ks))
app.listen(8080)
