const express = require("express");
const app = express();
const port = 3000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/get", (req, res) => {
  res.status(200).send("Hello World!");
});

app.post("/post", (req, res) => {
  let myJson = req.body;
  res.status(200).send(myJson);
});
app.post("/postform", (req, res) => {
  res.status(200).send(JSON.stringify(req.body));
});
app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
