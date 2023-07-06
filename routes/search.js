const express = require("express")
const router = express.Router()
const path = require("path")
const fs = require("fs")

function grep(file, term, type){
  if(!file.endsWith(".md")){
    return {}
  }
  
  const contents = fs.readFileSync(file, "utf-8").toString()
  const lines = contents.split("\n")

  for(let l in lines){
    const line = lines[l]
    if((type==="header" ? !line.startsWith("#") : false) || !line.toLowerCase().includes(term))
      continue
   
    return {
      "name": file.replace(/^.+?[/]/, ""),
      "link":`${file}`,
      "line": line
    }
  } 

  return {}
}

function _find(dir, term, type){
  if(dir.includes(".git")){
    return []
  }

  const notes = fs.readdirSync(dir) 
  let results = []

  for(let n in notes){
    const current = path.join(dir, notes[n])
    if(fs.statSync(current).isDirectory()){
      results = results.concat(_find(current, term, type))
    }else{
      let res={}
      if(type==="dir"){
        if(current.includes(term)){
          res = grep(current, "", type)
        }
      }
      else{
        res = grep(current, term, type)
      }
      if(Object.keys(res).length==0)
        continue
      results.push(res)
    }
  } 

  return results
}

function find(term, type){
  const dir = "./notes"
  const notes = fs.readdirSync(dir)
  let results = []

  for(let n in notes){
    const note = notes[n]
    results = results.concat(
      _find(path.join(dir, note), term, type)
    )
  }
    
  return results
}

router.get("/", (req,res)=>{
  const term = req.query.term 
  const type = req.query.type

  const results = find(term, type)
  res.render("search", { results: results })
})

module.exports = router
