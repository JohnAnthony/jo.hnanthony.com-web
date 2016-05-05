compress = require 'compression'
express = require 'express'
path = require 'path'

app = express()
app.use(express.static(path.join(__dirname, '/static')))
app.use compress()
app.use (req, res) ->
	res.status(404).send '404: Page not found'
app.use (err, req, res, next) ->
	res.status(500).send '500: Internal server error'
app.listen(29000)
