/*
Part of exercise file for go lang course at
https://web.learncodeonline.in
*/

const express = require('express')
const app = express()
const port = 1111
let data =[]
let dataForm =[]

app.use(express.json()); 
app.use(express.urlencoded({extended: true}));

app.get('/', (req, res) => {
  res.status(200).send("Welcome to My server BITCH")
})

app.get('/getData', (req, res) => {
    res.status(200).send(data)



  })

  app.get('/get', (req, res) => {
    res.status(200).json({
      name: "Mahmoud",
      age: 23,
      uni : "IT",
      
    })
  })


app.post('/post', (req, res) => {
    let myJson = req.body;
    data.push(req.body)     // your JSON
	
	res.status(200).send(myJson);
})

app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
    // dataForm.push(req.body)
    data.push(req.body)
})
  

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})