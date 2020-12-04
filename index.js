const http = require('http')
const express = require('express')
const multer = require('multer')
const fs = require('fs')
const { promisify } = require('util')
const pipeline = promisify(require('stream').pipeline)

const app = express()

const upload = multer()

const server = http.createServer(app)

app.post('/upload', upload.single('file'), async (req, res) => {
  console.log(req.file)
  await pipeline(
    req.file.stream,
    fs.createWriteStream('./upload.txt')
  )
  res.sendStatus(200)
})

server.listen(8080, () => {
  console.log('Listening on port 8080')
})
