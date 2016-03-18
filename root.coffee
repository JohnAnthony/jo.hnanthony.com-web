express = require 'express'
jade = require 'jade'
path = require 'path'

app = express()
app.use '/js', express.static(
	path.join(__dirname, 'bower_components/jquery/dist')
)
app.use '/js', express.static(
	path.join(__dirname, 'bower_components/jquery-ui')
)
app.use '/fullpage', express.static(
	path.join(__dirname, 'bower_components/fullpage.js/dist')
)

app.use '/', (req, res) ->
	res.send jade.renderFile('jade/index.jade')
app.listen 8080
