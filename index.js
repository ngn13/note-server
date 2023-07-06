const express = require("express")
const config = require("./libs/config.js")
const app = express()

app.use("/search", require("./routes/search.js"));
app.use("/notes", require("./routes/notes.js"));
app.use("/", express.static("./static"))
app.set("view engine", "ejs")

app.get("/", (req,res)=>{
  res.render("index", {version: "1.0"})
})

app.get("*", (req,res)=>{
  res.redirect("/")
})

app.listen(config["port"], ()=>{
  console.log("Note Server listening on port "+config["port"])
})
