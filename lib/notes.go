package lib

import (
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

type Note struct {
	Sum      string
	Name     string
	Fullpath string
	Created  time.Time
}

func (n *Note) Load(dir string) error {
	var (
		err     error
		content []byte
		file    *os.File
		size    int
	)

	if n.Name = strings.TrimPrefix(n.Fullpath, dir); n.Name[0] == '/' {
		n.Name = n.Name[1:]
	}

	content = make([]byte, 200)

	if file, err = os.Open(n.Fullpath); err != nil {
		return err
	}
	defer file.Close()

	if size, err = io.ReadFull(file, content); err != nil && err != io.ErrUnexpectedEOF {
		return err
	}

	if size == 200 {
		content = append(content, []byte("...")...)
	}

	n.Sum = string(content)
	return nil
}

func (n *Note) Read() ([]byte, error) {
	var (
		content []byte
		err     error
	)

	if content, err = os.ReadFile(n.Fullpath); err != nil {
		return nil, err
	}

	return content, nil
}

func GetNotes(dir string) ([]Note, error) {
	var (
		err error
		st  os.FileInfo
		all []Note
	)

	if st, err = os.Stat(dir); err != nil {
		return nil, fmt.Errorf("failed to access notes path: %s", dir)
	}

	if !st.IsDir() {
		return nil, fmt.Errorf("notes path is not a directory: %s", dir)
	}

	all = findNotes(dir)

	sort.Slice(all, func(i, j int) bool {
		return all[i].Created.After(all[j].Created)
	})

	for i := range all {
		if err = all[i].Load(dir); err != nil {
			return nil, fmt.Errorf("failed to load the note \"%s\": %s", all[i].Name, err.Error())
		}
	}

	return all, nil
}

func findNotes(dir string) []Note {
	var (
		files  []os.DirEntry
		result []Note
		st     os.FileInfo
		fp     string
		note   Note
		err    error
	)

	if files, err = os.ReadDir(dir); err != nil {
		Warn("failed to access \"%s\" while searching for notes: %s", dir, err.Error())
		return result
	}

	for _, f := range files {
		fp = path.Join(dir, f.Name())

		if st, err = os.Stat(fp); err != nil {
			Warn("failed to check the \"%s\" while searching for notes: %s", fp, err.Error())
			continue
		}

		if st.IsDir() {
			result = append(result, findNotes(path.Join(fp))...)
		}

		if !strings.HasSuffix(fp, ".md") {
			continue
		}

		note = Note{
			Fullpath: path.Clean(fp),
			Created:  st.ModTime(),
		}

		result = append(result, note)
	}

	return result
}
