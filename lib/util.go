package lib

import (
	"os"
	"path"
)

func getSum(content []byte) string {
	if len(content) > 200 {
		content = append(content[:200], []byte("...")...)
	}
	return string(content)
}

func FindFile(file string, dir string) string {
	var (
		files []os.DirEntry
		st    os.FileInfo
		fp    string
		res   string
		err   error
	)

	if files, err = os.ReadDir(dir); err != nil {
		Warn("failed to access \"%s\" while searching for \"%s\": %s", dir, file, err)
		return ""
	}

	for _, f := range files {
		if f.Name() == "." || f.Name() == ".." {
			continue
		}

		fp = path.Join(dir, f.Name())

		if st, err = os.Stat(fp); err != nil {
			Warn("failed to check \"%s\" while searching for \"%s\": %s", fp, file, err.Error())
			continue
		}

		if st.IsDir() {
			if res = FindFile(file, fp); res != "" {
				return res
			}
		}

		if f.Name() == file {
			return fp
		}
	}

	return ""
}
