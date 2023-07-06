const express = require("express")
const router = express.Router()
const path = require("path")
const fs = require("fs")
const { marked } = require("marked")
const createDomPurify = require("dompurify")
const { JSDOM } = require("jsdom")
const dompurify = createDomPurify(new JSDOM().window)

function handler(req, res){
  const file = path.join("./notes", req.path)
  if(!file.endsWith(".md"))
    return res.redirect("/")

  const content = fs.readFileSync(file).toString()
  
  res.render("note.ejs", { content: dompurify.sanitize(marked(content, {mangle: false, headerIds: false})) })
}

function _create_route(file){
  if(!file.endsWith(".md"))
    return

  router.get("/"+file.replace(/^.+?[/]/, ""), handler)
}

function create_route(dirname){

  const dir = fs.readdirSync(dirname)
  for(let d in dir){
    const current = path.join(dirname, dir[d])
    if(current.includes(".git")){
      continue
    }
    if(fs.statSync(current).isDirectory()){
      create_route(current)
      continue
    }
    _create_route(current) 
  }
}

create_route("./notes")

module.exports = router
