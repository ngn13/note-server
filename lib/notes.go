package lib

import (
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type Note struct {
  Sum string
  Path string
  Content string
  Created time.Time
}

func CheckNotes() {
  st, err := os.Stat("notes")
  if err != nil && os.IsNotExist(err) {
    log.Fatal("Notes directory does not exists")
  }else if err != nil {
    log.Fatalf("Cannot access the notes directory: %s", err)
  }

  if !st.IsDir() {
    log.Fatal("Notes should be directory, not a file")
  }
}

func getSum(content []byte) string {
  sum := string(content)
  if(len(sum) > 200){
    return sum[:200]+"..."
  }
  return sum
}

func getNotesDir(dirname string) []Note{
  var result []Note = []Note{}
  var fulldir string

  if(dirname != "") {
    fulldir = path.Join("notes", dirname)
  }else {
    fulldir = "notes" 
  }

  files, err := os.ReadDir(fulldir)
  if err != nil {
    log.Warnf("Cannot read %s: %s", fulldir, err)
    return result 
  }

  for _, f := range files {
    pth := path.Join(fulldir, f.Name())

    st, err := os.Stat(pth)
    if err != nil {
      log.Warnf("Error checking path: %s", pth)
      continue
    }

    if st.IsDir() {
      result = append(result, getNotesDir(path.Join(dirname, f.Name()))...)
    }

    if(!strings.HasSuffix(pth, ".md")){
      continue
    }

    content, err := os.ReadFile(pth)
    if err != nil {
      log.Warnf("Cannot read file: %s", pth)
      continue
    }

    result = append(result, Note{
      Sum: getSum(content),
      Content: string(content),
      Path: path.Join(dirname, f.Name()),
      Created: st.ModTime(),
    })
  }

  return result
}

func searchDir(file string, dir string) string {
  files, err := os.ReadDir(dir)
  if err != nil {
    log.Warnf("Cannot read %s: %s", dir, err)
    return "" 
  }

  for _, f := range files {
    pth := path.Join(dir, f.Name())

    st, err := os.Stat(pth)
    if err != nil {
      log.Warnf("Error checking path: %s", pth)
      continue
    }

    if st.IsDir() {
      res := searchDir(file, path.Join(dir, f.Name()))
      if res != "" {
        return res
      }
    }

    if f.Name() == file {
      return pth 
    }
  }

  return "" 
}

func FindFile(f string) string {
  res := searchDir(f, "notes")
  return res
}

func GetNotes() []Note{
  all := getNotesDir("")
  sort.Slice(all, func(i, j int) bool {
    return all[i].Created.After(all[j].Created)
  })
  return all
}
