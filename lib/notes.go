package lib

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

type Note struct {
	Sum      string
	Path     string
	Fullpath string
	Created  time.Time
}

func (n *Note) Load() error {
	var (
		err     error
		content []byte
		fp      string
	)

	if fp = n.Fullpath; fp == "" {
		fp = n.Path
	}

	if content, err = os.ReadFile(fp); err != nil {
		Warn("failed to read \"%s\" loading the note: %s", fp, err.Error())
		return err
	}

	n.Sum = getSum(content)
	return nil
}

func (n *Note) Read() ([]byte, error) {
	var (
		content []byte
		err     error
		fp      string
	)

	if fp = n.Fullpath; fp == "" {
		fp = n.Path
	}

	if content, err = os.ReadFile(fp); err != nil {
		Fail("failed to read \"%s\" loading the note: %s", fp, err.Error())
		return nil, err
	}

	return content, nil
}

func (n *Note) SetFp(dir string) error {
	if n.Path == "" {
		return fmt.Errorf("note path is not set")
	}

	n.Fullpath = path.Join(dir, n.Path)
	return nil
}

func GetNotes(dir string) ([]Note, error) {
	var (
		err    error
		st     os.FileInfo
		all    []Note
		curdir string
	)

	if curdir, err = os.Getwd(); err != nil {
		return nil, fmt.Errorf("failed to get the working directory: %s", err.Error())
	}

	if st, err = os.Stat(dir); err != nil {
		return nil, fmt.Errorf("failed to access notes path: %s", dir)
	}

	if !st.IsDir() {
		return nil, fmt.Errorf("notes path is not a directory: %s", dir)
	}

	if err = os.Chdir(dir); err != nil {
		return nil, fmt.Errorf("failed to change directory to \"%s\": %s", dir, err.Error())
	}

	all = findNotes(".")

	sort.Slice(all, func(i, j int) bool {
		return all[i].Created.After(all[j].Created)
	})

	for i := range all {
		if err = all[i].SetFp(dir); err != nil {
			return nil, fmt.Errorf("failed to set the path for note: %s", err.Error())
		}
	}

	if err = os.Chdir(curdir); err != nil {
		return nil, fmt.Errorf("failed to change directory back to \"%s\": %s", curdir, err.Error())
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
			Path:     path.Clean(fp),
			Created:  st.ModTime(),
			Fullpath: "",
		}

		if note.Load() != nil {
			continue
		}

		result = append(result, note)
	}

	return result
}
