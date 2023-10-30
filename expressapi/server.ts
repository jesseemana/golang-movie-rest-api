import express, { Request, Response } from 'express'

// npm i express nodemon
// npm i -D @types/express @types/nodemon

const PORT = 8080

const app = express()

app.use(express.json())
app.use(express.urlencoded({ extended: true }))

app.get('/ping', (req: Request, res: Response) => {
  res.status(200).send('pong')
})

app.get('/get', (req: Request, res: Response) => {
  res.status(200).json({ message: "Welcome to Jay's backend REST API" })
})

app.post('/post', (req: Request, res: Response) => {
  let myBody = req.body
  res.status(200).send(myBody)
})

app.post('/postform', (req: Request, res: Response) => {
  res.status(200).send(JSON.stringify(req.body))
})

app.listen(PORT, () => console.log(`Server running on http://localhost:${PORT}`))   