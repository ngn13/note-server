package routes

import (
	"os"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/microcosm-cc/bluemonday"
	"github.com/ngn13/note-server/lib"
	"github.com/russross/blackfriday/v2"
)

func Ret404(c *fiber.Ctx) error {
  return c.Status(404).Render("error", fiber.Map{
    "message": "404 Not Found",
  }) 
}

func GetNote(c *fiber.Ctx) error {
  pth := c.Path()
  
  // no you aint getting any lfi 
  bad := []string{"..", "\\"} 
  for _, b := range bad {
    if strings.Contains(pth, b) {
      return Ret404(c)
    }
  }

  if pth == "/" {
    return Ret404(c)
  }

  if !strings.HasSuffix(pth, ".md") {
    _, f := path.Split(pth)
    fpth := lib.FindFile(f) 
    if fpth == "" {
      return Ret404(c)
    }

    return c.SendFile(fpth)
  }

  s := bluemonday.UGCPolicy()
  pth = s.Sanitize(pth)

  fp := path.Join("notes", pth)
  content, err := os.ReadFile(fp)
  if err != nil && os.IsNotExist(err) {
    return Ret404(c)
  }else if err != nil {
    log.Warnf("Cannot read file: %s", fp)
    return c.Status(500).Render("error", fiber.Map{
      "message": "500 Server Error",
    }) 
  }

  ext := blackfriday.FencedCode
  ext |= blackfriday.SpaceHeadings
  ext |= blackfriday.HardLineBreak

  md := blackfriday.Run(
    content, 
    blackfriday.WithExtensions(ext),
  )

  return c.Render("note", fiber.Map{
    "path": pth,
    "markdown": string(md), 
  })
}
