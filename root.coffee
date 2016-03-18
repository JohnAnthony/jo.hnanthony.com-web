compress = require('compression')
express = require('express')
path = require('path')
vhost = require('vhost')

ja = require('./jo.hnanthony.com/app.coffee')

ks = express()
ks.use(express.static(path.join(__dirname, "kieransanderson.xyz")))
ks.use(compress())
ks.use((req, res) ->
	res.status(404).send('404: Page not Found')
)
ks.use((error, req, res, next) ->
	res.status(500).send('500: Internal Server Error')
)

app = express()
app.use(vhost("jo.hnanthony.com", ja))
app.use(vhost("www.jo.hnanthony.com", ja))
app.use(vhost("kieransanderson.xyz", ks))
app.use(vhost("www.kieransanderson.xyz", ks))
app.listen(8080)
