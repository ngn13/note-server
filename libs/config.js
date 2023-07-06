const fs = require("fs")

let config = {}
try {
  config = JSON.parse(fs.readFileSync("./config.json"))
} catch (error) {
  console.log("Cannot read config!")
  process.exit()
}

let required = ["port"]
for(let k in config){
  if(!required.includes(k))
    break;
  required.splice(required.indexOf(k), 1)
}

if(required.length!==0){
  console.log("Following keys are required and not defined in the config:")
  console.log(required)
  process.exit()
}

module.exports = config
